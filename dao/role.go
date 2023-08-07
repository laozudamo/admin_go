package dao

import (
	"admin_go/global"
	"admin_go/models"
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
	if err := global.DB.Save(role).Error; err != nil {
		return err
	}
	return nil
}

func RoleDelete(ids []uint) error {
	if err := global.DB.Delete(&models.SysRole{}, ids).Error; err != nil {
		return err
	}
	return nil
}

// func RoleList(params map[string]interface{}) ([]models.SysRole, error) {
// 	// var lists []models.SysRole

// }
