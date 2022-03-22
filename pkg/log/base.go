package log

import (
	"fmt"
	"os"
	"strings"
	"time"

	"WhaMan/config"

	"github.com/sirupsen/logrus"
)

// Logger 全局日志对象
var Logger *logrus.Logger

// myFormatter 实现 logrus.Formatter 接口，自定义输出格式
type myFormatter struct{}

func (f *myFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	timestamp := time.Now().Local().Format("2006/01/02-15:04:05")
	msg := fmt.Sprintf("[%s] [%s] %s %s:%d %s\n",
		config.Conf.ProjectName,
		strings.ToUpper(entry.Level.String()), timestamp,
		entry.Caller.File, entry.Caller.Line, entry.Message)
	return []byte(msg), nil
}

// Init 初始化日志配置
func Init(level string) {
	Logger = logrus.New()

	// 配置日志输出：则输出到控制台
	Logger.SetOutput(os.Stdout)

	// 设置日志级别
	l, err := logrus.ParseLevel(level)
	if err != nil {
		panic(err)
	}
	Logger.SetLevel(l)

	// 设置在输出日志中添加文件名和方法信息
	Logger.SetReportCaller(true)

	// 设置自定义输出格式
	Logger.SetFormatter(new(myFormatter))
}
