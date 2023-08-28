package actions

import "admin_go/global"

func DeleteAction(m interface{}) error {
	if err := global.DB.Delete(m).Error; err != nil {
		return err
	}
	return nil
}
