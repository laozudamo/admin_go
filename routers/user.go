package routers

import (
	"admin_go/controller"

	"github.com/gin-gonic/gin"
)

func InitUserRouter(router *gin.RouterGroup) {
	userRouter := router.Group("users")
	{
		userRouter.POST("/login", controller.Login)
	}
}
