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
	ClientUsername     string `json:"clientUsername" binding:"required"`
	FreelancerUsername string `json:"freelancerUsername" binding:"required"`
}

func checkClientAndFreelancer(clientUsername, freelancerUsername string) (bson.ObjectID, bson.ObjectID, error) {
	var res1 models.Client
	var res2 models.Freelancer

	coll := utils.DBClient.Database(os.Getenv("DATABASE_NAME")).Collection("clients")
	filter := bson.D{{Key: "username", Value: clientUsername}}
	err := coll.FindOne(context.TODO(), filter).Decode(&res1)
	if err != nil {
		return bson.ObjectID{}, bson.ObjectID{}, err
	}

	coll = utils.DBClient.Database(os.Getenv("DATABASE_NAME")).Collection("freelancers")
	filter = bson.D{{Key: "username", Value: freelancerUsername}}
	err = coll.FindOne(context.TODO(), filter).Decode(&res2)
	if err != nil {
		return bson.ObjectID{}, bson.ObjectID{}, err
	}

	return res1.ID, res2.ID, nil
}

func CreateJob(c *gin.Context) {
	var reqBody ReqStruct

	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	clientID, freelancerID, err := checkClientAndFreelancer(reqBody.ClientUsername, reqBody.FreelancerUsername)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	coll := utils.DBClient.Database(os.Getenv("DATABASE_NAME")).Collection("jobs")
	doc := models.Job{ClientID: clientID, FreelancerID: freelancerID, Status: "Active"}

	res, err := coll.InsertOne(context.Background(), doc)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to Create Job! " + err.Error()})
		return
	}

	jobId := res.InsertedID.(bson.ObjectID).Hex()

	c.JSON(http.StatusOK, gin.H{
		"jobId":   jobId,
		"message": "Job Created",
	})
}
