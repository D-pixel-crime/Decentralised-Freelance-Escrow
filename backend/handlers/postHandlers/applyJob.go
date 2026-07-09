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

type applyJobRequest struct {
	JobID string `json:"jobId" binding:"required"`
}

func ApplyJob(c *gin.Context) {
	roleStr := c.GetString("role")
	if roleStr != "freelancer" {
		c.JSON(http.StatusForbidden, gin.H{"error": "Only freelancers are authorized to apply for jobs"})
		return
	}

	ethAccountStr := c.GetString("ethAccount")
	if ethAccountStr == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing or invalid ethAccount in context"})
		return
	}

	var req applyJobRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	jobIdObj, err := bson.ObjectIDFromHex(req.JobID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Job ID format"})
		return
	}

	dbName := os.Getenv("DATABASE_NAME")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{"_id": jobIdObj}
	update := bson.M{
		"$push": bson.M{
			"applicants": ethAccountStr,
		},
	}

	result, err := utils.DBClient.Database(dbName).Collection("jobs").UpdateOne(ctx, filter, update)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to apply for job"})
		return
	}

	if result.MatchedCount == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Job not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Applied successfully"})
}
