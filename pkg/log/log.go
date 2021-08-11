package log

import (
	"fmt"
	"os"
	"strings"
	"time"

	"WhaMan/pkg/global"

	"github.com/sirupsen/logrus"
)

// myFormatter 实现 logrus.Formatter 接口，自定义输出格式
type myFormatter struct{}

func (f *myFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	timestamp := time.Now().Local().Format("2006/01/02-15:04:05")
	msg := fmt.Sprintf("[%s] [%s] %s %s:%d %s\n",
		global.ProjectName,
		strings.ToUpper(entry.Level.String()), timestamp,
		entry.Caller.File, entry.Caller.Line, entry.Message)
	return []byte(msg), nil
}

const (
	DebugLevel = "Debug"
	InfoLevel  = "Info"
	WarnLevel  = "Warn"
	ErrorLevel = "Error"
	FatalLevel = "Fatal"
)

// New 初始化日志配置
func New(level string) *logrus.Logger {
	log := logrus.New()

	// 配置日志输出：则输出到控制台
	log.SetOutput(os.Stdout)

	// 设置日志级别
	switch level {
	case DebugLevel:
		log.SetLevel(logrus.DebugLevel)
	case InfoLevel:
		log.SetLevel(logrus.InfoLevel)
	case WarnLevel:
		log.SetLevel(logrus.WarnLevel)
	case ErrorLevel:
		log.SetLevel(logrus.ErrorLevel)
	case FatalLevel:
		log.SetLevel(logrus.FatalLevel)
	default:
		panic("未匹配的日志级别")
	}

	// 设置在输出日志中添加文件名和方法信息
	log.SetReportCaller(true)

	// 设置自定义输出格式
	log.SetFormatter(new(myFormatter))

	return log
}
