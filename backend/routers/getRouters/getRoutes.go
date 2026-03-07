package getRouters

import (
	"net/http"

	"github.com/gin-gonic/gin"

	gethandlers "github.com/D-pixel-crime/Freelance_Escrow/backend/handlers/getHandlers"
	"github.com/D-pixel-crime/Freelance_Escrow/backend/utils"
)

func GET_Routes(getRouter *gin.Engine) {
	getRouter.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	protected := getRouter.Group("/api/get")

	protected.Use(utils.AuthMiddleware())
	{
		protected.GET("/balance/:address", gethandlers.GetWalletBalance)
	}

}
