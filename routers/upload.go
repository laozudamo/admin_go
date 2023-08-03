package routers

import (
	"admin_go/controller"
	"admin_go/middlewares"

	"github.com/gin-gonic/gin"
)

func InitUpload(Router *gin.RouterGroup) {
	Router.POST("upload", middlewares.JWTAuth(), controller.Upload)
}
