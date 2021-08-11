package global

import (
	"WhaMan/pkg/config"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

// Conf 全局配置数据
var Conf *config.Conf

// DB 全局数据库操作对象
var DB *gorm.DB

// Log 全局日志对象
var Log *logrus.Logger
