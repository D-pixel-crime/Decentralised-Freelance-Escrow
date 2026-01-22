package main

import (
	"fmt"

	getRouter "github.com/D-pixel-crime/Freelance_Escrow/backend/routers/getRouters"
	"github.com/D-pixel-crime/Freelance_Escrow/backend/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	utils.LoadENV()

	router := gin.Default()

	getRouter.GET_Routes(router)

	fmt.Println("Server starting on :8080...")
	router.Run("localhost:3000")
}
