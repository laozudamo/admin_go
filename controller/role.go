package controller

import (
	"admin_go/dao"
	"admin_go/models"
	"admin_go/response"
	"admin_go/utils"

	"github.com/gin-gonic/gin"
)

func CreatRole(c *gin.Context) {
	role := models.SysRole{}

	if err := c.ShouldBind(&role); err != nil {
		utils.HandleValidatorError(c, err)
		return
	}

	role.SysMenu = []*models.SysMenu{}

	for _, id := range role.SysMenuIds {
		menu, err := dao.MenuFindById(id)
		if err != nil {
			response.Err(c, 500, 500, "菜单不存在", nil)
			return
		}
		role.SysMenu = append(role.SysMenu, menu)
	}

	temRole, _ := dao.RoleFindByName(role.RoleKey)

	if temRole != nil {
		response.Err(c, 400, 400, "角色已经存在", nil)
		return
	}

	if err := dao.RoleCreat(&role); err != nil {
		response.Err(c, 500, 500, "创建失败", nil)
		return
	}

	response.Success(c, 200, "创建成功", nil)
}

func GetRole(c *gin.Context) {
	id := c.Query("id")
	role, err := dao.RoleFindById(id)
	if err != nil {
		response.Err(c, 500, 500, "查询失败", nil)
		return
	}
	response.Success(c, 200, "查询成功", role)
}

func UpdateRole(c *gin.Context) {

}

func DeleteRole(c *gin.Context) {

}

func GetRoleList(c *gin.Context) {

}
