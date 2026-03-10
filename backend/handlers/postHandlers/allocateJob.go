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

type jobAllocationRequest struct {
	FreelancerEthAccount string        `json:"FreelancerEthAccount" binding:"required"`
	JobId                bson.ObjectID `json:"jobId" binding:"required"`
	ChainId              int           `json:"chainId" binding:"required"`
}

func checkFreelancer(freelancerEthAccount string) (bson.ObjectID, error) {
	var res models.Freelancer
	coll := utils.DBClient.Database(os.Getenv("DATABASE_NAME")).Collection("freelancer")
	filter := bson.M{"ethAccount": freelancerEthAccount}

	err := coll.FindOne(context.TODO(), filter).Decode(&res)
	if err != nil {
		return bson.ObjectID{}, err
	}

	return res.ID, nil
}

func jobAllocation(freelancerEthAccount string, jobId bson.ObjectID) error {
	var res models.Job
	coll := utils.DBClient.Database(os.Getenv("DATABASE_NAME")).Collection("jobs")
	filter := bson.M{"_id": jobId}

	freelancerId, err := checkFreelancer(freelancerEthAccount)
	if err != nil {
		return err
	}

	update := bson.M{"$set": bson.M{"freelancerId": freelancerId, "status": models.AGREED}}
	err = coll.FindOneAndUpdate(context.TODO(), filter, update).Decode(&res)
	if err != nil {
		return err
	}

	return nil
}

func AllocateJob(c *gin.Context) {
	var reqBody jobAllocationRequest
	var err error

	if err = c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Incorrect Request Format!"})
		return
	}

	err = jobAllocation(reqBody.FreelancerEthAccount, reqBody.JobId)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			c.JSON(http.StatusNotFound, gin.H{"error": "User Not Found!"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error!"})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Job Allocated!"})
}
