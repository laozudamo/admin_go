package routers

import (
	"admin_go/controller"

	"github.com/gin-gonic/gin"
)

func InitVideoRouter(router *gin.RouterGroup) {
	Router := router.Group("/video")
	{
		Router.GET("/test", controller.HandleVideo)
	}
}
