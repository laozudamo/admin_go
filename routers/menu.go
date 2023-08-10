package routers

import (
	"admin_go/controller"

	"github.com/gin-gonic/gin"
)

func InitMenu(Router *gin.RouterGroup) {
	router := Router.Group("/menu")
	{
		router.POST("", controller.CreatMenu)
		// router.PUT("", controller.UpdateRole)
		// router.DELETE("", controller.DeleteRole)
		// router.GET("", controller.GetRole)
		// router.GET("/list", controller.GetRoleList)
	}
}
