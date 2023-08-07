package actions

import (
	"admin_go/global"
)

func GetTotal(model map[string]interface{}) (int, error) {
	var count int64
	if err := global.DB.Model(&model).Count(&count).Error; err != nil {
		return 0, err
	}
	return int(count), nil
}
