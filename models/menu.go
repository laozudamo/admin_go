package models

type Menu struct {
	Name     string `gorm:"name" json:"name"`                                              // 菜单名称
	Path     string `gorm:"path" json:"path"`                                              // 菜单路径
	ParentId uint   `gorm:"parent_id" json:"parent_id"`                                    // 父级菜单ID
	Icon     string `gorm:"icon" json:"icon"`                                              // 菜单图标
	Children []Menu `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"children"` // 子菜单
}
