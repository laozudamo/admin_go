package controller

import (
	"admin_go/models/request"
	"admin_go/utils"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	loginForm := request.Login{}
	if err := c.ShouldBind(&loginForm); err != nil {
		utils.HandleValidatorError(c, err)
		return
	}
	//user := models.SysUser{
	//	Username: loginForm.UserName,
	//}

	//isPass := utils.CheckPwd(loginForm.Password, userInfo.Salt, userInfo.Password)
	//fmt.Println(isPass)

}

func Logout(c *gin.Context) {

}

func AddUser(c *gin.Context) {

}

func DelUser(c *gin.Context) {

}

func UpdateUser(c *gin.Context) {

}

func EditUser(c *gin.Context) {

}

func GetUserList(c *gin.Context) {

}

func GetUserInfo(c *gin.Context) {

}
