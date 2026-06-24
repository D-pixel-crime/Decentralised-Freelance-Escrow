package main

import (
	"os"

	"github.com/D-pixel-crime/Freelance_Escrow/backend/indexer"
	authRouter "github.com/D-pixel-crime/Freelance_Escrow/backend/routers/authRouters"
	getRouter "github.com/D-pixel-crime/Freelance_Escrow/backend/routers/getRouters"
	postRouter "github.com/D-pixel-crime/Freelance_Escrow/backend/routers/postRouters"
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

	_, err = utils.ConnectToWeb3()
	if err != nil {
		log.Errorf("Failed to connect to Web3: %v", err)
	} else {
		jobsColl := utils.DBClient.Database(os.Getenv("DATABASE_NAME")).Collection("jobs")
		go indexer.StartIndexer(jobsColl)
	}

	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000", "http://localhost:8080"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"},
		AllowCredentials: true,
	}))

	authRouter.AUTH_Routes(router)
	getRouter.GET_Routes(router)
	postRouter.POST_Routes(router)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Infof("Server starting on port: %s...", port)
	router.Run(":" + port)
}
