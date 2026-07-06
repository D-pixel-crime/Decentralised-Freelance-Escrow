package posthandlers

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

func UpdateProfile(c *gin.Context) {
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

	var profileUpdate models.FreelancerProfile
	if err := c.ShouldBindJSON(&profileUpdate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if len(profileUpdate.DocumentCIDs) > 3 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Maximum of 3 documents allowed"})
		return
	}

	dbName := os.Getenv("DATABASE_NAME")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{"ethAccount": ethAccountStr}
	update := bson.M{
		"$set": bson.M{
			"profile.bio":            profileUpdate.Bio,
			"profile.resumeLink":     profileUpdate.ResumeLink,
			"profile.experience":     profileUpdate.Experience,
			"profile.education":      profileUpdate.Education,
			"profile.techStack":      profileUpdate.TechStack,
			"profile.githubLink":     profileUpdate.GithubLink,
			"profile.leetcodeLink":   profileUpdate.LeetCodeLink,
			"profile.codeforcesLink": profileUpdate.CodeforcesLink,
			"profile.documentCids":   profileUpdate.DocumentCIDs,
		},
	}

	result, err := utils.DBClient.Database(dbName).Collection("freelancer").UpdateOne(ctx, filter, update)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update profile"})
		return
	}

	if result.MatchedCount == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Freelancer not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Profile updated successfully"})
}
