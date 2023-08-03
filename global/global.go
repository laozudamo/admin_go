package global

import (
	"admin_go/config"

	"gorm.io/gorm"
)

var (
	DB       *gorm.DB
	Settings config.ServerConfig
)
