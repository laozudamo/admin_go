package models

import "gorm.io/gorm"

type SysRole struct {
	RoleName   string `form:"name" json:"name" gorm:"size:128;comment:角色名称" binding:"required"` // 角色名称
	RoleKey    string `form:"roleKey" json:"roleKey" gorm:"size:128;" binding:"required"`       //角色代码
	Sort       int    `form:"sort" json:"sort" gorm:"size:4;comment:排序" binding:"required"`     // 排序
	Status     int    `form:"int" json:"status" binding:"required"`                             // 状态 1启用 0禁用
	Remark     string `form:"remark" json:"remark" gorm:"size:255;comment:备注"`                  // 描述
	SysMenuIds []uint `form:"menuIds" gorm:"-" json:"menuIds"  binding:"required"`
	SysDeptIds []uint `form:"deptIds" json:"deptIds" gorm:"-"`

	SysMenu []*SysMenu `gorm:"many2many:sys_role_menu;" json:"menu"` // 角色菜单关联表
	SysDept []*SysDept `gorm:"many2many:sys_role_dept"`              // 角色部门关联表
	gorm.Model
}
