package routers

import (
	"admin_go/controller"

	"github.com/gin-gonic/gin"
)

func InitRole(Router *gin.RouterGroup) {
	router := Router.Group("/role")
	{
		router.POST("", controller.CreatRole)
		router.PUT("", controller.UpdateRole)
		router.DELETE("", controller.DeleteRole)
		router.GET("", controller.GetRole)
		router.GET("/list", controller.GetRole)
	}
}
