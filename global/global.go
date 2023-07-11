package global

import (
	"go.uber.org/zap"
	"golang_web/conf"
	"gorm.io/gorm"
)

var (
	Logger      *zap.SugaredLogger
	DB          *gorm.DB
	RedisClient *conf.RedisClient
)
