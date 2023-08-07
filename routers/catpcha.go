package routers

import (
	"admin_go/controller"

	"github.com/gin-gonic/gin"
)

func InitCatpcha(router *gin.RouterGroup) {
	catpchaRouter := router.Group("/")
	{
		catpchaRouter.GET("catpcha", controller.GetCaptcha)
	}
}
