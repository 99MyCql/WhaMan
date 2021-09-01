package validate

import (
	"reflect"

	"WhaMan/pkg/global/models"
)

// MyDatetimeValidate MyDatetime类型验证，空值返回nil
func MyDatetimeValidate(field reflect.Value) interface{} {
	if field.Type() == reflect.TypeOf(models.MyDatetime{}) {
		t := field.Interface().(models.MyDatetime)
		if t.Valid == false {
			return nil
		}
		// 此处不能再返回models.MyDatetime类型，不然会陷入死循环，应返回内置类型
		return t.Valid
	}
	return nil
}
