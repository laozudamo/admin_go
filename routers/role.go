package routers

import (
	"admin_go/controller"

	"github.com/gin-gonic/gin"
)

func InitRole(Router *gin.RouterGroup) {
	router := Router.Group("/role")
	{
		router.POST("", controller.CeeatRole)
		// router.POST("/update", controller.UpdateRole)
		// router.POST("/delete", controller.DeleteRole)
		// router.POST("/get", controller.GetRole)
		// router.POST("/list", controller.GetRoleList)
	}
}
