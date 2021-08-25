package middleware

import (
	"WhaMan/pkg/global"

	"github.com/gin-gonic/gin"
	"github.com/unrolled/secure"
)

// TlsHandler 用于TLS
func TlsHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		secureMiddleware := secure.New(secure.Options{
			SSLRedirect: true,
			SSLHost:     global.Conf.Host + ":" + global.Conf.Port,
		})
		err := secureMiddleware.Process(c.Writer, c.Request)

		// If there was an error, do not continue.
		if err != nil {
			return
		}

		c.Next()
	}
}
