package posthandlers

import (
	"context"
	"net/http"
	"os"
	"time"

	"github.com/D-pixel-crime/Freelance_Escrow/backend/utils"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type DeliverJobRequest struct {
	JobID          string `json:"jobId" binding:"required"`
	DeliverableCID string `json:"deliverableCid" binding:"required"`
}

func DeliverJob(c *gin.Context) {
	// Verify freelancer role
	role, exists := c.Get("role")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Role not found in context"})
		return
	}
	roleStr, ok := role.(string)
	if !ok || roleStr != "freelancer" {
		c.JSON(http.StatusForbidden, gin.H{"error": "Only freelancers can deliver jobs"})
		return
	}

	ethAccount, exists := c.Get("ethAccount")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "EthAccount not found in context"})
		return
	}
	ethAccountStr, ok := ethAccount.(string)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid ethAccount type"})
		return
	}

	var req DeliverJobRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	jobID, err := bson.ObjectIDFromHex(req.JobID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid job ID"})
		return
	}

	dbName := os.Getenv("DATABASE_NAME")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// 1. Get freelancer ID using ethAccount
	var freelancer struct {
		ID bson.ObjectID `bson:"_id"`
	}
	err = utils.DBClient.Database(dbName).Collection("freelancer").FindOne(ctx, bson.M{"ethAccount": ethAccountStr}).Decode(&freelancer)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Freelancer not found"})
		return
	}

	// 2. Ensure the freelancer is assigned to this job
	filter := bson.M{
		"_id":          jobID,
		"freelancerId": freelancer.ID,
	}

	update := bson.M{
		"$set": bson.M{
			"deliverableCid": req.DeliverableCID,
		},
	}

	result, err := utils.DBClient.Database(dbName).Collection("jobs").UpdateOne(ctx, filter, update)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update job"})
		return
	}

	if result.MatchedCount == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Job not found or not assigned to this freelancer"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Job deliverable submitted successfully"})
}
