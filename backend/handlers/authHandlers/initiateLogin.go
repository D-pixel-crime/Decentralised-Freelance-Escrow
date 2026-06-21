package authhandlers

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/D-pixel-crime/Freelance_Escrow/backend/models"
	"github.com/D-pixel-crime/Freelance_Escrow/backend/utils"
	"github.com/gin-gonic/gin"
	"github.com/spruceid/siwe-go"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type LoginRequest struct {
	EthAccount string `json:"ethAccount" binding:"required"`
	Role       string `json:"role" binding:"required"`
}

func checkUserAndProduceNonce(ethAccount, role string) (string, error) {
	filter := bson.M{"ethAccount": ethAccount}
	// Creating nonce
	nonce := siwe.GenerateNonce()

	var err error

	switch role {
	case "client":
		var res models.Client
		coll := utils.DBClient.Database(os.Getenv("DATABASE_NAME")).Collection("client")
		err = coll.FindOne(context.TODO(), filter).Decode(&res)
	case "freelancer":
		var res models.Freelancer
		coll := utils.DBClient.Database(os.Getenv("DATABASE_NAME")).Collection("freelancer")
		err = coll.FindOne(context.TODO(), filter).Decode(&res)
	default:
		return "", fmt.Errorf("Invalid User Type!")
	}

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return "", err
		}
		return "", fmt.Errorf("Error Fetching from Database! Error:%s", err)
	}

	err = utils.RedisClient.Set(context.TODO(), "nonce:"+ethAccount, nonce, 5*time.Minute).Err()
	if err != nil {
		return "", fmt.Errorf("Error Setting Nonce in Redis! Error:%s", err)
	}

	return nonce, nil
}

func InitiateLogin(c *gin.Context) {
	var reqBody LoginRequest

	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	nonce, err := checkUserAndProduceNonce(reqBody.EthAccount, reqBody.Role)
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

	c.JSON(http.StatusOK, gin.H{
		"nonce": nonce,
	})
}
