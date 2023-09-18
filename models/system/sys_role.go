package system

import "gorm.io/gorm"

type SysRole struct {
	RoleName  string `json:"role_name"`
	RoleKey   string `json:"role_key"`
	RoleDesc  string `json:"role_desc"`
	SysUserID int    `json:"sys_user_id"`
	gorm.Model
}
