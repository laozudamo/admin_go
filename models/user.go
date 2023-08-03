package models

import (
	"gorm.io/gorm"
)

type User struct {
	Name     string `grom:"name" json:"name"`                                           // 用户名称
	Password string `grom:"password" json:"password"`                                   // 用户密码
	Des      string `gorm:"des" json:"des"`                                             // 用户描述
	Avatar   string `gorm:"avatar" json:"avatar"`                                       // 用户头像
	Roles    []Role `grom:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"roles"` // 用户角色
	gorm.Model
}
