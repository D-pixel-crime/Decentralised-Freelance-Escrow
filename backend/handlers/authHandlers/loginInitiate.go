package authhandlers

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/D-pixel-crime/Freelance_Escrow/backend/models"
	"github.com/D-pixel-crime/Freelance_Escrow/backend/utils"
	"github.com/gin-gonic/gin"
	"github.com/spruceid/siwe-go"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type LoginRequest struct {
	UsernameOrEmail string `json:"username" binding:"required"`
	// Password        string `json:"password" binding:"required"`
	Role       string `json:"role" binding:"required"`
	EthAccount string `json:"ethAccount" binding:"required"`
}

func checkClientAndProduceNonce(usernameOrEmail string) (string, error) {
	filter := bson.M{
		"$or": []bson.M{
			{"username": usernameOrEmail},
			{"email": usernameOrEmail},
		},
	}
	// Creating nonce
	nonce := siwe.GenerateNonce()
	update := bson.M{"$set": bson.M{"nonce": nonce}}

	var res models.Client
	coll := utils.DBClient.Database(os.Getenv("DATABASE_NAME")).Collection("clients")
	err := coll.FindOneAndUpdate(context.TODO(), filter, update).Decode(&res)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return "", err
		}

		return "", fmt.Errorf("Error Fetching from Database! Error:%s", err)
	}

	return nonce, nil
}

func checkFreelancerAndProduceNonce(usernameOrEmail string) (string, error) {
	filter := bson.M{
		"$or": []bson.M{
			{"username": usernameOrEmail},
			{"email": usernameOrEmail},
		},
	}

	// Creating nonce
	nonce := siwe.GenerateNonce()
	update := bson.M{"$set": bson.M{"nonce": nonce}}

	var res models.Freelancer
	coll := utils.DBClient.Database(os.Getenv("DATABASE_NAME")).Collection("freelancers")
	err := coll.FindOneAndUpdate(context.TODO(), filter, update).Decode(&res)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return "", err
		}

		return "", fmt.Errorf("Error Fetching from Database! Error:%s", err)
	}

	return nonce, nil
}

func LoginInitiate(c *gin.Context) {
	var reqBody LoginRequest
	var nonce string
	var err error

	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	switch reqBody.Role {
	case "client":
		nonce, err = checkClientAndProduceNonce(reqBody.UsernameOrEmail)
	case "freelancer":
		nonce, err = checkFreelancerAndProduceNonce(reqBody.UsernameOrEmail)
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

	c.JSON(http.StatusOK, gin.H{
		"nonce": nonce,
	})
}
