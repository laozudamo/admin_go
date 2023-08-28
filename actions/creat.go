package actions

import "admin_go/global"

func CreatAction(m interface{}) error {
	if err := global.DB.Create(m).Error; err != nil {
		return err
	}
	return nil
}
