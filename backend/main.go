package main

import (
	"os"

	getRouter "github.com/D-pixel-crime/Freelance_Escrow/backend/routers/getRouters"
	"github.com/D-pixel-crime/Freelance_Escrow/backend/utils"
	"github.com/gin-contrib/cors"

	"github.com/charmbracelet/log"
	"github.com/gin-gonic/gin"
)

func main() {
	utils.LoadENV()
	_, err := utils.ConnectToDb()
	if err != nil {
		log.Fatal(err)
	}
	_, err = utils.ConnectToRedis()
	if err != nil {
		log.Fatal(err)
	}

	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000", "http://localhost:8080"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"},
		AllowCredentials: true,
	}))

	getRouter.GET_Routes(router)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Infof("Server starting on port: %s...", port)
	router.Run(":" + port)
}
