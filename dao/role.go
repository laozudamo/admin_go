package dao

import (
	"admin_go/global"
	"admin_go/models"
	"fmt"
)

// RoleCreat 创建角色
func RoleCreat(role *models.SysRole) error {
	if err := global.DB.Create(&role).Error; err != nil {
		return err
	}
	return nil
}

func RoleFindById(id string) (*models.SysRole, error) {
	role := models.SysRole{}
	if err := global.DB.Where("id = ?", id).Preload("SysMenu").First(&role).Error; err != nil {
		return nil, err
	}
	return &role, nil
}

func RoleFindByName(key string) (*models.SysRole, error) {
	role := models.SysRole{}
	if err := global.DB.Where("role_key = ?", key).First(&role).Error; err != nil {
		return nil, err
	}
	return &role, nil
}

func RoleUpdate(role *models.SysRole) error {
	if err := global.DB.Model(&models.SysRole{}).Where("id = ?", role.ID).Updates(role).Error; err != nil {
		return err
	}
	return nil
}

func RoleDelete(ids []int) error {
	if err := global.DB.Delete(&models.SysRole{}, ids).Error; err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func RoleList(page int, pageSize int, name string, roleKey string, status int) ([]models.SysRole, error) {
	var roles []models.SysRole
	if err := global.DB.Limit(pageSize).Offset((page - 1) * pageSize).Preload("SysMenu").Find(&roles).Error; err != nil {
		return nil, err
	}
	return roles, nil
}
