package getRouters

import (
	"net/http"

	"github.com/gin-gonic/gin"

	gethandlers "github.com/D-pixel-crime/Freelance_Escrow/backend/handlers/getHandlers"
)

func GET_Routes(getRouter *gin.Engine) {
	getRouter.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	getRouter.GET("/balance/:address", gethandlers.GetWalletBalance)
}
