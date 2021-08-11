package main

import (
	"WhaMan/pkg/config"
	"WhaMan/pkg/database"
	"WhaMan/pkg/global"
	"WhaMan/pkg/log"

	"github.com/gin-gonic/gin"
)

func init() {
	global.Log = log.New(log.DebugLevel)
	var err error
	if global.Conf, err = config.New("conf.yml"); err != nil {
		global.Log.Fatalf("%+v", err)
	}
	if global.DB, err = database.New(); err != nil {
		global.Log.Fatalf("%+v", err)
	}
}

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run(global.Conf.Host + ":" + global.Conf.Port)
}
