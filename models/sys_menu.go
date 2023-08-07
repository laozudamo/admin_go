package models

type SysMenu struct {
	MenuId   int    `json:"menuId" gorm:"primaryKey;autoIncrement"`
	MenuName string `json:"menuName" gorm:"size:128;"`
	Title    string `json:"title" gorm:"size:128;"`
	Path     string `json:"path" gorm:"size:128;"`
}
