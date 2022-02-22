package middleware

import (
	"bytes"
	"fmt"
	"io/ioutil"

	"WhaMan/pkg/log"

	"github.com/gin-gonic/gin"
)

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}
func (w bodyLogWriter) WriteString(s string) (int, error) {
	w.body.WriteString(s)
	return w.ResponseWriter.WriteString(s)
}

// Log 打印请求响应信息
func Log() func(c *gin.Context) {
	return func(c *gin.Context) {
		bodyLogWriter := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = bodyLogWriter

		info := fmt.Sprintf("===> %s:%s", c.Request.Method, c.Request.RequestURI)
		if c.Request.Method == "POST" {
			data, _ := ioutil.ReadAll(c.Request.Body) // c.Request.Body 一次性数据，读完就没有了
			info += fmt.Sprintf(", body:%s", string(data))
			c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(data)) // 将数据放回
		}
		log.Logger.Info(info)

		c.Next()

		log.Logger.Infof("<=== %s", bodyLogWriter.body.String())
	}
}
