package test

import (
	"path"
	"runtime"

	"WhaMan/config"
	"WhaMan/pkg/database"
	"WhaMan/pkg/log"
	"WhaMan/pkg/validate"
)

func Init() {
	_, curFile, _, _ := runtime.Caller(0)
	config.Init(path.Join(path.Dir(path.Dir(path.Dir(curFile))), "conf.yml"))
	log.Init("debug") // 初始化日志
	database.Init()   // 初始化数据库
	validate.Init()   // 初始化验证器
}
