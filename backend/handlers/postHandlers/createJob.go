package posthandlers

import (
	"context"
	"net/http"
	"os"

	"github.com/D-pixel-crime/Freelance_Escrow/backend/models"
	"github.com/D-pixel-crime/Freelance_Escrow/backend/utils"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type ReqStruct struct {
	ClientID     bson.ObjectID `json:"clientID"`
	FreelancerID bson.ObjectID `json:"freelancerID"`
}

func checkClientAndFreelancer(clientID, freelancerID bson.ObjectID) (bool, error) {
	coll := utils.DBClient.Database(os.Getenv("DATABASE_NAME")).Collection("clients")
	filer := bson.D{{Key: "rating", Value: clientID}}
	cnt, err := coll.CountDocuments(context.TODO(), filer)
	if err != nil {
		return false, err
	}
	if cnt <= 0 {
		return false, nil
	}

	coll = utils.DBClient.Database(os.Getenv("DATABASE_NAME")).Collection("freelancers")
	filer = bson.D{{Key: "rating", Value: freelancerID}}
	cnt, err = coll.CountDocuments(context.TODO(), filer)
	if err != nil {
		return false, err
	}
	if cnt <= 0 {
		return false, nil
	}

	return true, nil
}

func CreateJob(c *gin.Context) {
	var reqBody ReqStruct

	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	isPresent, err := checkClientAndFreelancer(reqBody.ClientID, reqBody.FreelancerID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if !isPresent {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Client or Freelancer Not Found!"})
		return
	}

	coll := utils.DBClient.Database(os.Getenv("DATABASE_NAME")).Collection("jobs")
	doc := models.Job{ClientID: reqBody.ClientID, FreelancerID: reqBody.FreelancerID, Status: "Active"}

	_, err = coll.InsertOne(context.Background(), doc)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to Create Job! " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Job Created",
	})
}
