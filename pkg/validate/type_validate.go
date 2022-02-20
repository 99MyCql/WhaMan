package validate

import (
	"reflect"

	"WhaMan/pkg/datetime"
)

// MyDatetimeValidate MyDatetime类型验证，空值返回nil
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
