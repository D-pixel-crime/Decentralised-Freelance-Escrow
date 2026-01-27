package authhandlers

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/D-pixel-crime/Freelance_Escrow/backend/models"
	"github.com/D-pixel-crime/Freelance_Escrow/backend/utils"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type SignupRequest struct {
	Username   string `json:"username" binding:"required"`
	Email      string `json:"email" binding:"required"`
	Password   string `json:"password" binding:"required"`
	Role       string `json:"role" binding:"required"`
	EthAccount string `json:"ethAccount" binding:"required"`
}

func clientSignup(username, email, passwordHash, ethAccount string) error {
	coll := utils.DBClient.Database(os.Getenv("DATATBASE_NAME")).Collection("client")
	doc := models.Client{BaseUser: models.BaseUser{Username: username, Email: email, Password: passwordHash, EthAccount: ethAccount}, RequestedJobs: []bson.ObjectID{}}

	_, err := coll.InsertOne(context.Background(), doc)
	if err != nil {
		return fmt.Errorf("Failed To Insert Client!")
	}

	return nil
}

func freelancerSignup(username, email, passwordHash, ethAccount string) error {
	coll := utils.DBClient.Database(os.Getenv("DATATBASE_NAME")).Collection("freelancers")
	doc := models.Freelancer{BaseUser: models.BaseUser{Username: username, Email: email, Password: passwordHash, EthAccount: ethAccount}, ActiveJobs: []bson.ObjectID{}}

	_, err := coll.InsertOne(context.Background(), doc)
	if err != nil {
		return fmt.Errorf("Failed To Insert Freelancer!")
	}

	return nil
}

func Signup(c *gin.Context) {
	var reqBody SignupRequest

	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	passwordHash, err := utils.HashPassword(reqBody.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error Hashing Password!"})
		return
	}

	accessToken, refreshToken, err := utils.GenerateTokens(reqBody.Username, reqBody.Email, reqBody.Role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error Generating Token: " + err.Error()})
		return
	}

	switch reqBody.Role {
	case "client":
		err = clientSignup(reqBody.Username, reqBody.Email, passwordHash, reqBody.EthAccount)
	case "freelancers":
		err = freelancerSignup(reqBody.Username, reqBody.Email, passwordHash, reqBody.EthAccount)
	default:
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Type!"})
		return
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":      "Signup successful",
		"username":     reqBody.Username,
		"email":        reqBody.Email,
		"rol":          reqBody.Role,
		"accessToken":  accessToken,
		"refreshToken": refreshToken,
	})
}
