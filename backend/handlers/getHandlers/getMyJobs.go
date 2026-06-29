package gethandlers

import (
	"context"
	"net/http"
	"os"
	"time"

	"github.com/D-pixel-crime/Freelance_Escrow/backend/models"
	"github.com/D-pixel-crime/Freelance_Escrow/backend/utils"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func GetMyJobs(c *gin.Context) {
	// Extract role and ethAccount from Gin context
	role, exists := c.Get("role")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Role not found in context"})
		return
	}
	roleStr, ok := role.(string)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid role type"})
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

	// Correcting type mismatches: The Gin context provides 'ethAccount' as a string, 
	// but the jobs collection uses 'bson.ObjectID' for ClientID/FreelancerID.
	// We convert the string ID (ethAccount) into an ObjectID by querying the respective collection first.
	var userID bson.ObjectID
	dbName := os.Getenv("DATABASE_NAME")
	
	if roleStr == "client" {
		var client models.Client
		err := utils.DBClient.Database(dbName).Collection("client").FindOne(context.TODO(), bson.M{"ethAccount": ethAccountStr}).Decode(&client)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Client not found"})
			return
		}
		userID = client.ID
	} else if roleStr == "freelancer" {
		var freelancer models.Freelancer
		err := utils.DBClient.Database(dbName).Collection("freelancer").FindOne(context.TODO(), bson.M{"ethAccount": ethAccountStr}).Decode(&freelancer)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Freelancer not found"})
			return
		}
		userID = freelancer.ID
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid role"})
		return
	}

	// Wrap in a timeout context for resiliency
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Query jobs collection
	coll := utils.DBClient.Database(dbName).Collection("jobs")
	
	var filter bson.M
	if roleStr == "client" {
		filter = bson.M{"clientId": userID}
	} else {
		filter = bson.M{"freelancerId": userID}
	}

	cursor, err := coll.Find(ctx, filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to query jobs"})
		return
	}
	defer cursor.Close(ctx)

	var jobs []models.Job
	if err := cursor.All(ctx, &jobs); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decode jobs"})
		return
	}

	if jobs == nil {
		jobs = []models.Job{}
	}

	c.JSON(http.StatusOK, jobs)
}
