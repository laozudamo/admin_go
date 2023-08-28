package controller

import (
	"admin_go/actions"
	"admin_go/models"
	"admin_go/response"
	"admin_go/utils"

	"github.com/gin-gonic/gin"
)

func CreatMenu(c *gin.Context) {
	menu := models.SysMenu{}
	if err := c.ShouldBindJSON(&menu); err != nil {
		utils.HandleValidatorError(c, err)
		return
	}
	if err := actions.CreatAction(&menu); err != nil {
		response.Err(c, 500, 500, "创建失败", nil)
		return
	}
	// response.Success(c, 200, "创建成功", nil)
}
