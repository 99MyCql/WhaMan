package rsp

import (
	"net/http"

	"WhaMan/consts"
	myErr "WhaMan/pkg/error"
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
		Code: consts.SucCode,
		Msg:  consts.CodeMsgMap[consts.SucCode],
		Data: data,
	}
}

func Err(e *myErr.Error) (int, *Response) {
	return e.Code / 1000, &Response{
		Code: e.Code,
		Msg:  e.Error(),
	}
}

// ResponseExample 用于接口文档的示例（interface{}类型被swag解析会出错）
type ResponseExample struct {
	Code int      `json:"code"`
	Msg  string   `json:"msg"`
	Data struct{} `json:"data"`
}
