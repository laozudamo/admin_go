package system

import "gorm.io/gorm"

type JwtBlacklist struct {
	Jwt string `gorm:"type:text;comment:jwt"`
	gorm.Model
}
