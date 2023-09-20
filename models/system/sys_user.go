package system

import "gorm.io/gorm"

type SysUser struct {
	UserName string
	Password string
	RoleKey  uint
	gorm.Model
}
