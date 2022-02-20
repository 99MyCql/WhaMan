package rsp

import (
	"fmt"
	"net/http"
	"time"

	_const "WhaMan/const"
	myErr "WhaMan/pkg/error"
)

const (
	sucCode = 200
	sucMsg  = "成功"
)

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func New(err error) (int, *Response) {
	if err != nil {
		return Err(err.(*myErr.Error))
	}
	return Suc(nil)
}

func NewWithData(data interface{}, err error) (int, *Response) {
	if err != nil {
		return Err(err.(*myErr.Error))
	}
	return Suc(data)
}

func Suc(data interface{}) (int, *Response) {
	return http.StatusOK, &Response{
		Code: sucCode,
		Msg:  sucMsg,
		Data: data,
	}
}

func Err(e *myErr.Error) (int, *Response) {
	return e.Code / 1000, &Response{
		Code: e.Code,
		Msg: fmt.Sprintf("[%s][%d] %s",
			time.Now().Format(_const.DatetimeFormat), e.Code, e.Msg),
	}
}

// ResponseExample 用于接口文档的示例（interface{}类型解析会出错）
type ResponseExample struct {
	Code int      `json:"code"`
	Msg  string   `json:"msg"`
	Data struct{} `json:"data"`
}
