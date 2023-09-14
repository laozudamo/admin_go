package models

import "gorm.io/gorm"

func (SysUser) TableName() string {
	return "sys_user"
}

type SysUser struct {
	Username string     `json:"username" gorm:"size:64;comment:用户名"`
	Password string     `json:"-" gorm:"size:128;comment:密码"`
	Status   StatusType `json:"status" gorm:"size:4;comment:状态"`
	RoleIds  []int      `json:"roleIds" gorm:"-"`
	Salt     string     `json:"-" gorm:"size:255;comment:盐"`
	gorm.Model
}

type StatusType int

const (
	Enable  StatusType = 0 // 禁用
	Disable StatusType = 1 // 启用
)
