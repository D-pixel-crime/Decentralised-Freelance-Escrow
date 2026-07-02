package postRouters

import (
	posthandlers "github.com/D-pixel-crime/Freelance_Escrow/backend/handlers/postHandlers"
	"github.com/D-pixel-crime/Freelance_Escrow/backend/utils"
	"github.com/gin-gonic/gin"
)

func POST_Routes(postRouter *gin.Engine) {
	protectedRouter := postRouter.Group("/api/post")

	protectedRouter.Use(utils.AuthMiddleware())
	{
		protectedRouter.POST("/job/create", posthandlers.CreateJob)
		protectedRouter.POST("/job/allocate", posthandlers.AllocateJob)
		protectedRouter.POST("/job/apply", posthandlers.ApplyJob)
		protectedRouter.POST("/profile/update", posthandlers.UpdateProfile)
	}
}
