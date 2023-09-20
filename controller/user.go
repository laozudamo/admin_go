package controller

import (
	"admin_go/dao"
	"admin_go/models/system"
	"admin_go/models/system/request"
	"admin_go/response"
	"admin_go/utils"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	l := request.Login{}
	if err := c.ShouldBindJSON(&l); err != nil {
		utils.HandleValidatorError(c, err)
		return
	}

	user, err := dao.FindUser(l.UserName)

	if err != nil {
		response.Err(c, 200, 400, "用户或用户名错误", nil)
		return
	}
	if ok := utils.BcryptCheck(l.Password, user.Password); !ok {
		response.Err(c, 200, 400, "用户或用户名错误", nil)
		return
	}

	if !store.Verify(l.CaptchaId, l.Captcha, true) {
		response.Err(c, 200, 400, "验证码错误", nil)
		return
	}

	token := utils.CreateToken(c, user.ID, user.RoleKey)

	response.Success(c, 200, "成功", token)
}

func Logout(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")
	if token == "" {
		response.Err(c, 200, 400, "参数错误", nil)
		return
	}

	//TOKEN加入黑名单

	//if err := utils.DelToken(token); err != nil {
	//	response.Err(c, 200, 400, "退出失败", nil)
	//	return
	//}
	response.Success(c, 200, "成功", nil)
}

func AddUser(c *gin.Context) {
	l := request.UserCrete{}

	if err := c.ShouldBindJSON(&l); err != nil {
		utils.HandleValidatorError(c, err)
		return
	}

	u, err := dao.FindUser(l.UserName)

	if u != nil && err == nil {
		response.Err(c, 200, 400, "用户已存在", nil)
		return
	}

	pwd := utils.BcryptHash(l.Password)

	user := system.SysUser{
		UserName: l.UserName,
		Password: pwd,
		RoleKey:  l.RoleKey,
	}

	if err := dao.AddUser(&user); err != nil {
		response.Err(c, 200, 400, "添加失败", nil)
		return
	}

	response.Success(c, 200, "添加成功", nil)
}

func DelUser(c *gin.Context) {
	id := c.Query("id")
	if id == "" {
		response.Err(c, 200, 400, "参数错误", nil)
		return
	}
	if err := dao.DelUser(id); err != nil {
		response.Err(c, 200, 400, "删除失败", nil)
		return
	}
	response.Success(c, 200, "删除成功", nil)
}

func GetUserList(c *gin.Context) {
	pageInfo := request.SysPageInfo{
		Page:     1,
		PageSize: 10,
	}
	userList, err := dao.GetUserList(&pageInfo)

	if err != nil {
		response.Err(c, 200, 400, "获取失败", nil)
		return
	}
	response.Success(c, 200, "获取成功", userList)
}

func UpdateUser(c *gin.Context) {
	l := request.UserUpdate{}
	if err := c.ShouldBindJSON(&l); err != nil {
		utils.HandleValidatorError(c, err)
		return
	}

	pwd := utils.BcryptHash(l.Password)

	l.Password = pwd

	if err := dao.UpdateUser(&l); err != nil {
		response.Err(c, 200, 400, "更新失败", nil)
		return
	}

	response.Success(c, 200, "更新成功", nil)
}
