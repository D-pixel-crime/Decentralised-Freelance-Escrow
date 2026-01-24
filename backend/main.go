package main

import (
	getRouter "github.com/D-pixel-crime/Freelance_Escrow/backend/routers/getRouters"
	"github.com/D-pixel-crime/Freelance_Escrow/backend/utils"

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

	getRouter.GET_Routes(router)

	log.Info("Server starting on port: 3000...")
	router.Run("localhost:3000")
}
