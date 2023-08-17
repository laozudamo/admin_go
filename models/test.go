package models

import "gorm.io/gorm"

type Student struct {
	ID       uint
	Name     string
	Teachers []Teacher `gorm:"many2many:student_teachers;"` // 多对多
	gorm.Model
}

type Teacher struct {
	ID       uint
	Name     string
	Students []Student `gorm:"many2many:student_teachers;"` // 多对多
	gorm.Model
}
