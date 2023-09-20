package routers

import (
	"admin_go/controller"
	"admin_go/middlewares"

	"github.com/gin-gonic/gin"
)

func InitUserRouter(router *gin.RouterGroup) {
	userRouter := router.Group("users")
	{
		userRouter.POST("/login", controller.Login)
		userRouter.GET("/logout", controller.Logout)

		userRouter.POST("", middlewares.JWTAuth(), controller.AddUser)
		userRouter.DELETE("", middlewares.JWTAuth(), controller.DelUser)
		userRouter.PUT("", middlewares.JWTAuth(), controller.UpdateUser)
		userRouter.GET("/list", middlewares.JWTAuth(), controller.GetUserList)
	}
}
