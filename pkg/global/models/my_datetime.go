package models

import (
	"database/sql/driver"
	"fmt"
	"time"

	"WhaMan/pkg/global"
)

const (
	location       = "Asia/Shanghai"
	dateFormat     = "\"" + global.DateFormat + "\""
	datetimeFormat = "\"" + global.DatetimeFormat + "\""
)

// MyDatetime 自定义时间格式，用于json序列化与非序列化、MySQL读写
type MyDatetime struct {
	Time  time.Time
	Valid bool
}

// UnmarshalJSON 从json中解析
func (t *MyDatetime) UnmarshalJSON(data []byte) error {
	dataStr := string(data)

	// 空值不进行解析（如果是空字符串，从json解析出的data字节数组会包含双引号）
	if dataStr == "\"\"" || dataStr == "null" {
		t.Time, t.Valid = time.Time{}, false
		return nil
	}

	// 按指定格式解析时间
	var err error
	local, _ := time.LoadLocation(location)
	if len(dataStr) == len(dateFormat) {
		t.Time, err = time.ParseInLocation(dateFormat, dataStr, local)
	} else {
		t.Time, err = time.ParseInLocation(dateFormat, dataStr, local)
	}
	if err != nil {
		return err
	}
	t.Valid = true
	return nil
}

// MarshalJSON 解析为json
func (t *MyDatetime) MarshalJSON() ([]byte, error) {
	if !t.Valid {
		return []byte("\"\""), nil
	}
	return []byte(fmt.Sprintf("\"%s\"", t.Time.Format(global.DatetimeFormat))), nil
}

// Value 写入 mysql 时调用
func (t MyDatetime) Value() (driver.Value, error) {
	// 遇到空值解析成 null
	if !t.Valid {
		return nil, nil
	}
	return t.Time, nil
}

// Scan 检出 mysql 时调用
func (t *MyDatetime) Scan(v interface{}) error {
	if v == nil {
		t.Time, t.Valid = time.Time{}, false
		return nil
	}

	t.Time = v.(time.Time)
	t.Valid = true
	return nil
}

// 用于 fmt.Println 和后续验证场景
func (t *MyDatetime) String() string {
	return t.Time.Format(global.DatetimeFormat)
}
