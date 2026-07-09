package authhandlers

import (
	"context"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/D-pixel-crime/Freelance_Escrow/backend/utils"
	"github.com/gin-gonic/gin"
	"github.com/spruceid/siwe-go"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type LoginVerificationRequest struct {
	EthAccount string `json:"ethAccount" binding:"required"`
	Role       string `json:"role" binding:"required,oneof=client freelancer arbitrator"`
	Message    string `json:"message" binding:"required"`
	Signature  string `json:"signature" binding:"required"`
}

func VerifyLogin(c *gin.Context) {
	var reqBody LoginVerificationRequest

	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	filter := bson.M{"ethAccount": reqBody.EthAccount}
	coll := utils.DBClient.Database(os.Getenv("DATABASE_NAME")).Collection(reqBody.Role)

	var user struct {
		Username string `bson:"username"`
		Email    string `bson:"email"`
	}

	if err := coll.FindOne(context.TODO(), filter).Decode(&user); err != nil {
		if err == mongo.ErrNoDocuments {
			c.JSON(http.StatusNotFound, gin.H{"error": "User Not Found!"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error Fetching from Database!"})
		}
		return
	}

	nonce, err := utils.RedisClient.Get(context.TODO(), "nonce:"+reqBody.EthAccount).Result()
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Nonce expired or invalid!"})
		return
	}
	utils.RedisClient.Del(context.TODO(), "nonce:"+reqBody.EthAccount)

	siweMessage, err := siwe.ParseMessage(reqBody.Message)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error Parsing SIWE Message!"})
		return
	}

	if _, err := siweMessage.Verify(reqBody.Signature, nil, &nonce, nil); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Signature Verification Failed!"})
		return
	}

	if !strings.EqualFold(siweMessage.GetAddress().String(), reqBody.EthAccount) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Signing Address Doesn't Match the Account Address!"})
		return
	}

	accessToken, refreshToken, err := utils.GenerateTokens(user.Username, user.Email, reqBody.Role, reqBody.EthAccount)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error Generating Token: " + err.Error()})
		return
	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("accessToken", accessToken, int(time.Hour*24), "/", "localhost", false, true)
	c.SetCookie("refreshToken", refreshToken, int(time.Hour*24*7), "/", "localhost", false, true)

	c.JSON(http.StatusOK, gin.H{
		"message":  "Login Successful",
		"username": user.Username,
		"email":    user.Email,
	})
}
