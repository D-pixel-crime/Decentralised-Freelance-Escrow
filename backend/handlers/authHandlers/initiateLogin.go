package authhandlers

import (
	"context"
	"net/http"
	"os"
	"time"

	"github.com/D-pixel-crime/Freelance_Escrow/backend/utils"
	"github.com/gin-gonic/gin"
	"github.com/spruceid/siwe-go"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type LoginRequest struct {
	EthAccount string `json:"ethAccount" binding:"required"`
	Role       string `json:"role" binding:"required,oneof=client freelancer arbitrator"`
}

func InitiateLogin(c *gin.Context) {
	var reqBody LoginRequest

	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	filter := bson.M{"ethAccount": reqBody.EthAccount}
	coll := utils.DBClient.Database(os.Getenv("DATABASE_NAME")).Collection(reqBody.Role)

	var result bson.M
	if err := coll.FindOne(context.TODO(), filter).Decode(&result); err != nil {
		if err == mongo.ErrNoDocuments {
			c.JSON(http.StatusNotFound, gin.H{"error": "User Not Found!"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error Fetching from Database!"})
		}
		return
	}

	nonce := siwe.GenerateNonce()
	if err := utils.RedisClient.Set(context.TODO(), "nonce:"+reqBody.EthAccount, nonce, 5*time.Minute).Err(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error Setting Nonce in Redis!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"nonce": nonce})
}
