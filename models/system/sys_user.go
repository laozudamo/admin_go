package system

import "gorm.io/gorm"

type SysUser struct {
	UserName string
	Password string
	SysRole  SysRole
	gorm.Model
}
