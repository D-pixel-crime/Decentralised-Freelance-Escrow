package posthandlers

import (
	"context"
	"net/http"
	"os"

	"github.com/D-pixel-crime/Freelance_Escrow/backend/models"
	"github.com/D-pixel-crime/Freelance_Escrow/backend/utils"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type jobCreationRequest struct {
	ClientEthAccount string        `json:"clientEthAccount" binding:"required"`
	JobId            bson.ObjectID `json:"jobId" binding:"required"`
	Title            string        `json:"title" binding:"required"`
	Description      string        `json:"description"`
	Deadline         string        `json:"deadline"`
	ContactEmail     string        `json:"contactEmail"`
	PayMin           float64       `json:"payMin"`
	PayMax           float64       `json:"payMax"`
}

func CreateJob(c *gin.Context) {
	var reqBody jobCreationRequest

	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Incorrect Request Format!"})
		return
	}

	clientColl := utils.DBClient.Database(os.Getenv("DATABASE_NAME")).Collection("client")
	var client models.Client

	if err := clientColl.FindOne(context.TODO(), bson.M{"ethAccount": reqBody.ClientEthAccount}).Decode(&client); err != nil {
		if err == mongo.ErrNoDocuments {
			c.JSON(http.StatusNotFound, gin.H{"error": "User Not Found!"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error!"})
		}
		return
	}

	jobsColl := utils.DBClient.Database(os.Getenv("DATABASE_NAME")).Collection("jobs")
	doc := models.Job{
		ClientID:     client.ID,
		Status:       models.UNALLOCATED,
		Title:        reqBody.Title,
		Description:  reqBody.Description,
		Deadline:     reqBody.Deadline,
		ContactEmail: reqBody.ContactEmail,
		PayMin:       reqBody.PayMin,
		PayMax:       reqBody.PayMax,
	}

	if _, err := jobsColl.InsertOne(context.Background(), doc); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create job"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Job Created!"})
}
