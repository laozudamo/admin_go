package models

import "gorm.io/gorm"

type Role struct {
	ID   uint   `json:"id" gorm:"primary_key"`
	Name string `json:"name" gorm:"type:varchar(20);not null;unique"`
	gorm.Model
}
