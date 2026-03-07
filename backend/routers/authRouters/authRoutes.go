package authrouters

import (
	"github.com/gin-gonic/gin"

	authHandlers "github.com/D-pixel-crime/Freelance_Escrow/backend/handlers/authHandlers"
)

func AUTH_Routes(authRouter *gin.Engine) {
	groupedAuth := authRouter.Group("/auth")

	groupedAuth.POST("/signup", authHandlers.Signup)
	groupedAuth.POST("/loginInitiate", authHandlers.LoginInitiate)
	groupedAuth.POST("/loginVerify", authHandlers.LoginVerify)
}
