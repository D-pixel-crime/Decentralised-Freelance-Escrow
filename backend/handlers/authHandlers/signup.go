package authhandlers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/D-pixel-crime/Freelance_Escrow/backend/models"
	"github.com/D-pixel-crime/Freelance_Escrow/backend/utils"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type SignupRequest struct {
	Username   string `json:"username" binding:"required"`
	Email      string `json:"email" binding:"required"`
	Role       string `json:"role" binding:"required"`
	EthAccount string `json:"ethAccount" binding:"required"`
}

func userSignup(username, email, ethAccount, role string) error {
	var err error

	switch role {
	case "client":
		coll := utils.DBClient.Database(os.Getenv("DATABASE_NAME")).Collection("client")
		doc := models.Client{BaseUser: models.BaseUser{Username: username, Email: email, EthAccount: ethAccount}, RequestedJobs: []bson.ObjectID{}}
		_, err = coll.InsertOne(context.Background(), doc)
	case "freelancer":
		coll := utils.DBClient.Database(os.Getenv("DATABASE_NAME")).Collection("freelancer")
		doc := models.Freelancer{BaseUser: models.BaseUser{Username: username, Email: email, EthAccount: ethAccount}, ActiveJobs: []bson.ObjectID{}}
		_, err = coll.InsertOne(context.Background(), doc)
	case "arbitrator":
		coll := utils.DBClient.Database(os.Getenv("DATABASE_NAME")).Collection("arbitrator")
		doc := models.Arbitrator{BaseUser: models.BaseUser{Username: username, Email: email, EthAccount: ethAccount}}
		_, err = coll.InsertOne(context.Background(), doc)
	default:
		return fmt.Errorf("Invalid User Type!")
	}

	if err != nil {
		if mongo.IsDuplicateKeyError(err) {
			return fmt.Errorf("Username or Email already exists!")

		}
		return fmt.Errorf("Failed To Insert Client!")
	}

	return nil
}

func Signup(c *gin.Context) {
	var reqBody SignupRequest

	if err := c.ShouldBindJSON(&reqBody); err != nil {
		log.Println("Request Format Error:", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "Incorrect Request Format!"})
		return
	}

	accessToken, refreshToken, err := utils.GenerateTokens(reqBody.Username, reqBody.Email, reqBody.Role, reqBody.EthAccount)
	if err != nil {
		log.Fatalln("Error Generating Tokens!", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error Generating Token: " + err.Error()})
		return
	}

	err = userSignup(reqBody.Username, reqBody.Email, reqBody.EthAccount, reqBody.Role)
	if err != nil {
		log.Println("Error Creating User!", err.Error())
		if err.Error() == "Invalid User Type!" {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("accessToken", accessToken, int(time.Hour*24), "/", "localhost", false, true)
	c.SetCookie("refreshToken", refreshToken, int(time.Hour*24*7), "/", "localhost", false, true)

	c.JSON(http.StatusOK, gin.H{
		"message": "Signup Successful.",
	})
}
