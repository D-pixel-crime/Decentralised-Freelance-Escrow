package gethandlers

import (
	"context"
	"net/http"
	"os"
	"time"

	"github.com/D-pixel-crime/Freelance_Escrow/backend/utils"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func GetContact(c *gin.Context) {
	address := c.Param("address")
	role := c.Query("role") // "freelancer", "arbitrator", or "client"

	if address == "" || role == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Address and role are required"})
		return
	}

	validRoles := map[string]bool{"freelancer": true, "client": true, "arbitrator": true}
	if !validRoles[role] {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid role"})
		return
	}

	dbName := os.Getenv("DATABASE_NAME")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var result struct {
		Email string `bson:"email" json:"email"`
	}

	filter := bson.M{"ethAccount": address}
	if objId, err := bson.ObjectIDFromHex(address); err == nil {
		filter = bson.M{"_id": objId}
	}

	err := utils.DBClient.Database(dbName).Collection(role).FindOne(ctx, filter).Decode(&result)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"email": result.Email})
}
