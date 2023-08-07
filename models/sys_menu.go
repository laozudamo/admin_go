package models

import "gorm.io/gorm"

type SysMenu struct {
	MenuName  string    `json:"menuName" gorm:"size:128;"`
	Title     string    `json:"title" gorm:"size:128;"`
	Icon      string    `json:"icon" gorm:"size:128;"`
	Path      string    `json:"path" gorm:"size:128;"`
	MenuType  string    `json:"menuType" gorm:"size:1;"`
	Action    string    `json:"action" gorm:"size:16;"`
	ParentId  int       `json:"parentId" gorm:"size:11;"` // 父菜单ID
	Component string    `json:"component" gorm:"size:255;"`
	Sort      int       `json:"sort" gorm:"size:4;"`
	Visible   string    `json:"visible" gorm:"size:1;"`
	DataScope string    `json:"dataScope" gorm:"-"`
	SysRoleId int       `gorm:"-"`
	Children  []SysMenu `json:"children,omitempty" gorm:"-"`
	gorm.Model
}
