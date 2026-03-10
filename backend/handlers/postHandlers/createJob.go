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
}

func checkClient(clientEthAccount string) (bson.ObjectID, error) {
	var res models.Client
	coll := utils.DBClient.Database(os.Getenv("DATABASE_NAME")).Collection("client")
	filter := bson.M{"ethAccount": clientEthAccount}

	err := coll.FindOne(context.TODO(), filter).Decode(&res)
	if err != nil {
		return bson.ObjectID{}, err
	}

	return res.ID, nil
}

func jobCreation(clientEthAccount string) error {
	clientId, err := checkClient(clientEthAccount)
	if err != nil {
		return err
	}

	coll := utils.DBClient.Database(os.Getenv("DATABASE_NAME")).Collection("jobs")
	doc := models.Job{ClientID: clientId, Status: models.UNALLOCATED}

	_, err = coll.InsertOne(context.Background(), doc)
	if err != nil {
		return err
	}

	return nil
}

func CreateJob(c *gin.Context) {
	var reqBody jobCreationRequest
	var err error

	if err = c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Incorrect Request Format!"})
		return
	}

	err = jobCreation(reqBody.ClientEthAccount)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			c.JSON(http.StatusNotFound, gin.H{"error": "User Not Found!"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error!"})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Job Created!"})
}
