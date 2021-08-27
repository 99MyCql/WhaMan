package middleware

import (
	"net/http"

	"WhaMan/pkg/global"
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
			global.Log.Error("未登录")
			c.JSON(http.StatusOK, rsp.Err(rsp.NotLoginErr))
			c.Abort()
			return
		}
		c.Next()
	}
}
