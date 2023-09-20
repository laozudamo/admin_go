package routers

import (
	"admin_go/controller"
	"admin_go/middlewares"

	"github.com/gin-gonic/gin"
)

func InitUpload(Router *gin.RouterGroup) {
	Router.Use(middlewares.JWTAuth())
	Router.POST("upload", controller.Upload)
}
