package validate

import (
	"reflect"
	"time"

	"WhaMan/pkg/datetime"
	"WhaMan/pkg/log"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

// DatetimeFormat 自定义验证，标签为datetime，用于验证字符串形式的日期时间格式
func DatetimeFormat(fl validator.FieldLevel) bool {
	if fl.Field().String() == "" {
		return true
	}
	if _, err := time.Parse(fl.Param(), fl.Field().String()); err != nil {
		return false
	}
	return true
}

// MyDatetimeValidate 用于MyDatetime类型验证，空值返回nil
func MyDatetimeValidate(field reflect.Value) interface{} {
	if field.Type() == reflect.TypeOf(datetime.MyDatetime{}) {
		t := field.Interface().(datetime.MyDatetime)
		if t.Valid == false {
			return nil
		}
		// 此处不能再返回models.MyDatetime类型，不然会陷入死循环，应返回内置类型
		return t.Valid
	}
	return nil
}

func Init() {
	// 注册验证器
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		if err := v.RegisterValidation("datetime", DatetimeFormat); err != nil {
			log.Logger.Fatal(err)
		}
		v.RegisterCustomTypeFunc(MyDatetimeValidate, datetime.MyDatetime{})
	}
}
