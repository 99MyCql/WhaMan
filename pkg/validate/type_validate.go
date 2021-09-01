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
		return t
	}
	return nil
}
