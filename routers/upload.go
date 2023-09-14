package routers

import (
	"admin_go/controller"

	"github.com/gin-gonic/gin"
)

func InitUpload(Router *gin.RouterGroup) {
	Router.POST("upload", controller.Upload)
	//Router.GET("findFile", controller.FindFile)
}
