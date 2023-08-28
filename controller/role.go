package controller

import (
	"admin_go/actions"
	"admin_go/dao"
	"admin_go/models"
	"admin_go/response"
	"admin_go/utils"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreatRole(c *gin.Context) {
	role := models.SysRole{}

	if err := c.ShouldBind(&role); err != nil {
		utils.HandleValidatorError(c, err)
		return
	}

	err := validatMenu(c, &role)
	if err != nil {
		response.Err(c, 500, 500, "菜单不存在", nil)
		return
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
	role := models.SysRole{}
	if err := c.ShouldBind(&role); err != nil {
		utils.HandleValidatorError(c, err)
		return
	}

	err := validatMenu(c, &role)
	if err != nil {
		response.Err(c, 500, 500, "菜单不存在", nil)
		return
	}

	temRole, _ := dao.RoleFindById(strconv.Itoa(int(role.ID)))

	if temRole == nil {
		response.Err(c, 500, 500, "没有找到该用户", nil)
		return
	}

	if err := dao.RoleUpdate(&role); err != nil {
		response.Err(c, 500, 500, "更新失败", nil)
		return
	}
	response.Success(c, 200, "更新成功", nil)
}

type Params struct {
	IDs []int `json:"ids" binding:"required"`
}

func DeleteRole(c *gin.Context) {

	params := Params{}
	if err := c.ShouldBind(&params); err != nil {
		utils.HandleValidatorError(c, err)
		return
	}

	if err := dao.RoleDelete(params.IDs); err != nil {
		response.Err(c, 500, 500, "删除失败", nil)
		return
	}
	response.Success(c, 200, "删除成功", nil)

}

func GetRoleList(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))
	name := c.Query("name")
	//fmt.Printf("name: %v\n", name)
	//return
	roleKey := c.Query("roleKey")
	status, _ := strconv.Atoi(c.Query("status"))
	fmt.Println(page, pageSize, name, roleKey, status)

	roles, err := dao.RoleList(page, pageSize, name, roleKey, status)
	if err != nil {
		response.Err(c, 500, 500, "查询失败", nil)
		return
	}
	total, err := actions.GetTotal(models.SysRole{})

	if err != nil {
		response.Err(c, 500, 500, "查询失败", nil)
		return
	}
	data := map[string]interface{}{
		"list":     roles,
		"total":    total,
		"page":     page,
		"pageSize": pageSize,
	}
	response.Success(c, 200, "查询成功", data)

}

func validatMenu(c *gin.Context, role *models.SysRole) error {

	role.SysMenu = []*models.SysMenu{}

	for _, id := range role.SysMenuIds {
		menu, err := dao.MenuFindById(id)
		if err != nil {

			return err
		}
		role.SysMenu = append(role.SysMenu, menu)
	}
	return nil
}
