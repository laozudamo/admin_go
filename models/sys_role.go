package models

import "gorm.io/gorm"

type SysRole struct {
	RoleName  string    `json:"roleName" gorm:"size:128;comment:角色名称" binding:"required"` // 角色名称
	SysUserID uint      `json:"userId" gorm:"comment:用户ID"`
	Status    string    `json:"status" gorm:"size:4;defult:1"`        // 状态 1禁用 2正常
	MenuIDs   []uint    `json:"menuIds" gorm:"-"`                     // 角色菜单关联表
	DeptIDs   []uint    `json:"deptIds" gorm:"-"`                     // 角色部门关联表
	SysMenu   []SysMenu `gorm:"many2many:sys_role_menu;" json:"menu"` // 角色菜单关联表
	SysDept   []SysDept `gorm:"many2many:sys_role_dept"`              // 角色部门关联表
	Remark    string    `json:"remark" gorm:"size:255;comment:备注"`    // 描述
	gorm.Model
}
