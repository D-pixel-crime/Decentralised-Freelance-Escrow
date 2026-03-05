package authhandlers

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/D-pixel-crime/Freelance_Escrow/backend/models"
	"github.com/D-pixel-crime/Freelance_Escrow/backend/utils"
	"github.com/gin-gonic/gin"
	"github.com/spruceid/siwe-go"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type LoginVerification struct {
	EthAccount string `json:"ethAccount" binding:"required"`
	Role       string `json:"role" binding:"required"`
	Message    string `json:"message" binding:"required"`
	Signature  string `json:"signature" binding:"required"`
}

func signatureVerification(ethAccount, message, signature, nonce string, update, filter bson.M, coll *mongo.Collection) error {
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

	_, _ = coll.UpdateOne(context.TODO(), filter, update)

	return nil
}

// returns username, email, role and error
func verifyClient(ethAccount, message, signature string) (string, string, error) {
	filter := bson.M{"ethAccount": ethAccount}
	update := bson.M{"$set": bson.M{"nonce": ""}}

	var res models.Client
	coll := utils.DBClient.Database(os.Getenv("DATABASE_NAME")).Collection("clients")
	err := coll.FindOne(context.TODO(), filter).Decode(&res)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return "", "", err
		}
		return "", "", fmt.Errorf("Error Fetching from Database! Error:%s", err)
	}

	if signatureVerification(ethAccount, message, signature, res.Nonce, update, filter, coll) != nil {
		return "", "", err
	}

	return res.Username, res.Email, nil
}

func verifyFreelancer(ethAccount, message, signature string) (string, string, error) {
	filter := bson.M{"ethAccount": ethAccount}
	update := bson.M{"$set": bson.M{"nonce": ""}}

	var res models.Freelancer
	coll := utils.DBClient.Database(os.Getenv("DATABASE_NAME")).Collection("freelancers")
	err := coll.FindOne(context.TODO(), filter).Decode(&res)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return "", "", err
		}
		return "", "", fmt.Errorf("Error Fetching from Database! Error:%s", err)
	}

	if err = signatureVerification(ethAccount, message, signature, res.Nonce, update, filter, coll); err != nil {
		return "", "", err
	}

	return res.Username, res.Email, nil
}

func LoginPhase2(c *gin.Context) {
	var reqBody LoginVerification
	var username, email, role string
	var err error

	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	switch reqBody.Role {
	case "client":
		username, email, err = verifyClient(reqBody.EthAccount, reqBody.Message, reqBody.Signature)
	case "freelancer":
		username, email, err = verifyFreelancer(reqBody.EthAccount, reqBody.Message, reqBody.Signature)
	default:
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid User Type!"})
		return
	}
	if err != nil {
		if err == mongo.ErrNoDocuments {
			c.JSON(http.StatusNotFound, gin.H{"error": "User Not Found!"})
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

	c.JSON(http.StatusOK, gin.H{
		"username":     username,
		"email":        email,
		"role":         role,
		"ethAccount":   reqBody.EthAccount,
		"accessToken":  accessToken,
		"refreshToken": refreshToken,
	})
}
