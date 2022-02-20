package middleware

import (
	myErr "WhaMan/pkg/error"
	"WhaMan/pkg/log"
	"WhaMan/pkg/rsp"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// AuthSession 通过session验证是否登录
func AuthSession() func(c *gin.Context) {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		// TODO: 设置超时时间
		if session.Get("isLogin") == nil {
			log.Logger.Error("未登录")
			c.JSON(rsp.Err(myErr.NotLogin))
			c.Abort()
			return
		}
		c.Next()
	}
}
