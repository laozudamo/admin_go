package actions

import "admin_go/global"

func IndexAction(m interface{}, fileds string, value string) (interface{}, error) {
	if err := global.DB.Where(fileds+"= ?", value).First(&m).Error; err != nil {
		return nil, err
	}
	return m, nil
}
