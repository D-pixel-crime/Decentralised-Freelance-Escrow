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
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type LoginRequest struct {
	UsernameOrEmail string `json:"username" binding:"required"`
	Password        string `json:"password" binding:"required"`
	Role            string `json:"role" binding:"required"`
}

func getClientHash(usernameOrEmail string) (string, string, error) {
	filter := bson.M{
		"$or": []bson.M{
			{"username": usernameOrEmail},
			{"email": usernameOrEmail},
		},
	}

	var res models.Client
	coll := utils.DBClient.Database(os.Getenv("DATABASE_NAME")).Collection("clients")
	err := coll.FindOne(context.TODO(), filter).Decode(&res)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return "", "", fmt.Errorf("No documents found!")
		}

		return "", "", fmt.Errorf("Error Fetching from Database! Error:%s", err)
	}

	return res.Password, res.Email, nil
}

func getFreelancerHash(usernameOrEmail string) (string, string, error) {
	filter := bson.M{
		"$or": []bson.M{
			{"username": usernameOrEmail},
			{"email": usernameOrEmail},
		},
	}

	var res models.Freelancer
	coll := utils.DBClient.Database(os.Getenv("DATABASE_NAME")).Collection("freelancers")
	err := coll.FindOne(context.TODO(), filter).Decode(&res)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return "", "", err
		}

		return "", "", fmt.Errorf("Error Fetching from Database! Error:%s", err)
	}

	return res.Password, res.Email, nil
}

func Login(c *gin.Context) {
	var reqBody LoginRequest
	var passwordHash string
	var email string
	var err error

	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	switch reqBody.Role {
	case "clients":
		passwordHash, email, err = getClientHash(reqBody.UsernameOrEmail)
	case "freelancer":
		passwordHash, email, err = getFreelancerHash(reqBody.UsernameOrEmail)
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

	isMatch := utils.ComparePasswordAndHash(passwordHash, reqBody.Password)
	if !isMatch {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Incorrect Password!"})
		return
	}

	accessToken, refreshToken, err := utils.GenerateTokens(reqBody.UsernameOrEmail, email, reqBody.Role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error Generating Token: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":      "Login successful",
		"user":         reqBody.UsernameOrEmail,
		"role":         reqBody.Role,
		"accessToken":  accessToken,
		"refreshToken": refreshToken,
	})
}
