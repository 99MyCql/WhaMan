package test

import (
	"path"
	"runtime"

	"WhaMan/pkg/config"
	"WhaMan/pkg/database"
	"WhaMan/pkg/log"

	"github.com/sirupsen/logrus"
)

func Init() {
	_, curFile, _, _ := runtime.Caller(0)
	config.Init(path.Join(path.Dir(path.Dir(path.Dir(curFile))), "conf.yml"))
	log.Init(logrus.DebugLevel)
	database.Init()
}
