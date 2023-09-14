package models

import "gorm.io/gorm"

type SyeRole struct {
	RoleName  string `json:"role_name"`
	RoleDesc  string `json:"role_desc"`
	SysUserID int    `json:"sys_user_id"`
	gorm.Model
}
