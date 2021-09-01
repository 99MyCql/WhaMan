package global

import (
	"WhaMan/pkg/config"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var (
	// Conf 全局配置数据
	Conf *config.Conf

	// DB 全局数据库操作对象
	DB *gorm.DB

	// Log 全局日志对象
	Log *logrus.Logger
)
