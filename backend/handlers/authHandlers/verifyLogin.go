package authhandlers

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/D-pixel-crime/Freelance_Escrow/backend/models"
	"github.com/D-pixel-crime/Freelance_Escrow/backend/utils"
	"github.com/gin-gonic/gin"
	"github.com/spruceid/siwe-go"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type LoginVerificationRequest struct {
	EthAccount string `json:"ethAccount" binding:"required"`
	Role       string `json:"role" binding:"required"`
	Message    string `json:"message" binding:"required"`
	Signature  string `json:"signature" binding:"required"`
}

func signatureVerification(ethAccount, message, signature, nonce string) error {
	siweMesssage, err := siwe.ParseMessage(message)
	if err != nil {
		return fmt.Errorf("Error Parsing SIWE Message! Error:%s", err)
	}

	_, err = siweMesssage.Verify(signature, nil, &nonce, nil)
	if err != nil {
		return fmt.Errorf("Signature Verification Failed! Error:%s", err)
	}

	if !strings.EqualFold(siweMesssage.GetAddress().String(), ethAccount) {
		return fmt.Errorf("Signing Address Doesn't Match the Account Address! Error:%s", err)
	}

	return nil
}

func verifyUser(ethAccount, role, message, signature string) (string, string, error) {
	filter := bson.M{"ethAccount": ethAccount}

	var err error
	var username, email string
	var coll *mongo.Collection

	switch role {
	case "client":
		var res models.Client
		coll = utils.DBClient.Database(os.Getenv("DATABASE_NAME")).Collection("client")
		err = coll.FindOne(context.TODO(), filter).Decode(&res)
		if err != nil {
			if err == mongo.ErrNoDocuments {
				return "", "", err
			}
			return "", "", fmt.Errorf("Error Fetching from Database! Error:%s", err)
		}

		username = res.Username
		email = res.Email
	case "freelancer":
		var res models.Freelancer
		coll = utils.DBClient.Database(os.Getenv("DATABASE_NAME")).Collection("freelancer")
		err = coll.FindOne(context.TODO(), filter).Decode(&res)
		if err != nil {
			if err == mongo.ErrNoDocuments {
				return "", "", err
			}
			return "", "", fmt.Errorf("Error Fetching from Database! Error:%s", err)
		}

		username = res.Username
		email = res.Email
	default:
		return "", "", fmt.Errorf("Invalid User Type!")
	}

	nonce, err := utils.RedisClient.Get(context.TODO(), "nonce:"+ethAccount).Result()
	if err != nil {
		return "", "", fmt.Errorf("Error getting nonce from Redis (may be expired)! Error:%s", err)
	}
	utils.RedisClient.Del(context.TODO(), "nonce:"+ethAccount)

	if err = signatureVerification(ethAccount, message, signature, nonce); err != nil {
		return "", "", err
	}
	return username, email, nil
}

func VerifyLogin(c *gin.Context) {
	var reqBody LoginVerificationRequest

	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	username, email, err := verifyUser(reqBody.EthAccount, reqBody.Role, reqBody.Message, reqBody.Signature)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			c.JSON(http.StatusNotFound, gin.H{"error": "User Not Found!"})
		} else if err.Error() == "Invalid User Type!" {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error!"})
		}
		return
	}

	accessToken, refreshToken, err := utils.GenerateTokens(username, email, reqBody.Role, reqBody.EthAccount)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error Generating Token: " + err.Error()})
		return
	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("accessToken", accessToken, int(time.Hour*24), "/", "localhost", false, true)
	c.SetCookie("refreshToken", refreshToken, int(time.Hour*24*7), "/", "localhost", false, true)

	c.JSON(http.StatusOK, gin.H{
		"message":  "Login Successful",
		"username": username,
		"email":    email,
	})
}
