package dao

import (
	"admin_go/global"
	"admin_go/models/system"
	"admin_go/models/system/request"
	Response "admin_go/models/system/response"
)

func FindUser(userName string) (*system.SysUser, error) {
	user := system.SysUser{}
	if err := global.DB.First(&user, "user_name = ?", userName).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func AddUser(user *system.SysUser) error {
	if err := global.DB.Create(&user).Error; err != nil {
		return err
	}
	return nil
}

func DelUser(id string) error {
	if err := global.DB.Delete(&system.SysUser{}, id).Error; err != nil {
		return err
	}
	return nil
}

func UpdateUser(m *request.UserUpdate) error {
	user := system.SysUser{
		Password: m.Password,
		RoleKey:  m.RoleKey,
	}

	if err := global.DB.Model(&system.SysUser{}).Where("id = ?", m.ID).Updates(&user).Error; err != nil {
		return err
	}
	return nil
}

func GetUserList(pageInfo *request.SysPageInfo) ([]*Response.User, error) {
	var userList []*Response.User
	if err := global.DB.Model(&system.SysUser{}).Offset((pageInfo.Page-1)*pageInfo.PageSize).Limit(pageInfo.PageSize).Select("user_name", "id", "role_key", "creat_at").Find(&userList).Error; err != nil {
		return nil, err
	}
	return userList, nil

}
