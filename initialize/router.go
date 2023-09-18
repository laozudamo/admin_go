package initialize

import (
	"admin_go/middlewares"
	"admin_go/routers"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	Router := gin.Default()

	Router.Use(middlewares.Cors(), middlewares.GinLogger(), middlewares.GinRecovery(true))
	// 图片路径 大小限制
	Router.MaxMultipartMemory = 8 << 20 // 8 MiB
	// 静态文件
	Router.Static("/uploads", "./uploads")
	//Router.Static("/video", "./video")
	Router.StaticFile("/video", "./video")

	ApiGroup := Router.Group("/")
	routers.InitUserRouter(ApiGroup)
	routers.InitCatpcha(ApiGroup)
	routers.InitUpload(ApiGroup)
	routers.InitVideoRouter(ApiGroup)
	return Router
}
