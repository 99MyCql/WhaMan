package validators

import (
	"time"

	"github.com/go-playground/validator/v10"
)

// DatetimeFormat 用于验证字符串形式的日期时间格式
func DatetimeFormat(fl validator.FieldLevel) bool {
	if fl.Field().String() == "" {
		return true
	}
	if _, err := time.Parse(fl.Param(), fl.Field().String()); err != nil {
		return false
	}
	return true
}
