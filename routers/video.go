package routers

import (
	"admin_go/controller"
	"admin_go/middlewares"

	"github.com/gin-gonic/gin"
)

func InitVideoRouter(router *gin.RouterGroup) {

	Router := router.Group("/video")
	Router.Use(middlewares.JWTAuth())
	{
		Router.GET("/test", controller.HandleVideo)
	}
}
