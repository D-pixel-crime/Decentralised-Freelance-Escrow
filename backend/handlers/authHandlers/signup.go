package authhandlers

import (
	"context"
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

func Signup(c *gin.Context) {
	var reqBody SignupRequest

	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Incorrect Request Format!"})
		return
	}

	accessToken, refreshToken, err := utils.GenerateTokens(reqBody.Username, reqBody.Email, reqBody.Role, reqBody.EthAccount)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error Generating Token: " + err.Error()})
		return
	}

	var doc interface{}
	baseUser := models.BaseUser{
		Username:   reqBody.Username,
		Email:      reqBody.Email,
		EthAccount: reqBody.EthAccount,
	}

	switch reqBody.Role {
	case "client":
		doc = models.Client{BaseUser: baseUser, RequestedJobs: []bson.ObjectID{}}
	case "freelancer":
		doc = models.Freelancer{BaseUser: baseUser, ActiveJobs: []bson.ObjectID{}}
	case "arbitrator":
		doc = models.Arbitrator{BaseUser: baseUser}
	default:
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid User Type!"})
		return
	}

	coll := utils.DBClient.Database(os.Getenv("DATABASE_NAME")).Collection(reqBody.Role)
	if _, err := coll.InsertOne(context.Background(), doc); err != nil {
		if mongo.IsDuplicateKeyError(err) {
			c.JSON(http.StatusConflict, gin.H{"error": "Username or Email already exists!"})
			return
		}
		log.Println("Error Creating User!", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to Insert User!"})
		return
	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("accessToken", accessToken, int(time.Hour*24), "/", "localhost", false, true)
	c.SetCookie("refreshToken", refreshToken, int(time.Hour*24*7), "/", "localhost", false, true)

	c.JSON(http.StatusOK, gin.H{"message": "Signup Successful."})
}
