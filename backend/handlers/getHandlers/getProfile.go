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

func GetProfile(c *gin.Context) {
	role, exists := c.Get("role")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Role not found in context"})
		return
	}
	roleStr, ok := role.(string)
	if !ok || roleStr != "freelancer" {
		c.JSON(http.StatusForbidden, gin.H{"error": "Only freelancers have a profile"})
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

	dbName := os.Getenv("DATABASE_NAME")
	var freelancer models.Freelancer

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := utils.DBClient.Database(dbName).Collection("freelancer").FindOne(ctx, bson.M{"ethAccount": ethAccountStr}).Decode(&freelancer)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Freelancer not found"})
		return
	}

	c.JSON(http.StatusOK, freelancer.Profile)
}

func GetProfileByAddress(c *gin.Context) {
	address := c.Param("address")
	if address == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Address is required"})
		return
	}

	dbName := os.Getenv("DATABASE_NAME")
	var freelancer models.Freelancer

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := utils.DBClient.Database(dbName).Collection("freelancer").FindOne(ctx, bson.M{"ethAccount": address}).Decode(&freelancer)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Freelancer not found"})
		return
	}

	c.JSON(http.StatusOK, freelancer.Profile)
}
