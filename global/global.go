package global

import (
	"admin_go/config"

	ut "github.com/go-playground/universal-translator"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	DB       *gorm.DB
	Settings config.ServerConfig
	Lg       *zap.Logger
	Trans    ut.Translator
)
