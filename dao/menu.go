package dao

import (
	"admin_go/global"
	"admin_go/models"
)

// MenuCreat 创建菜单
func MenuCreat(Menu *models.SysMenu) error {
	if err := global.DB.Create(&Menu).Error; err != nil {
		return err
	}
	return nil
}

func MenuFindById(id uint) (*models.SysMenu, error) {
	Menu := models.SysMenu{}
	if err := global.DB.Where("id = ?", id).First(&Menu).Error; err != nil {
		return nil, err
	}
	return &Menu, nil
}

func MenuFindByName(path string) (*models.SysMenu, error) {
	Menu := models.SysMenu{}
	if err := global.DB.Where("component_path = ?", path).First(&Menu).Error; err != nil {
		return nil, err
	}
	return &Menu, nil
}

func MenuUpdate(Menu *models.SysMenu) error {
	if err := global.DB.Save(Menu).Error; err != nil {
		return err
	}
	return nil
}

func MenuDelete(ids []uint) error {
	if err := global.DB.Delete(&models.SysMenu{}, ids).Error; err != nil {
		return err
	}
	return nil
}
