package rsp

import (
	"fmt"
	"time"

	"WhaMan/pkg/global"
)

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func Suc() Response {
	return Response{
		Code: Success,
		Msg:  codeMsgMap[Success],
	}
}

func SucWithData(data interface{}) Response {
	return Response{
		Code: Success,
		Msg:  codeMsgMap[Success],
		Data: data,
	}
}

func Err(code int) Response {
	return Response{
		Code: code,
		Msg: fmt.Sprintf("[%d][%s][%s]", code, codeMsgMap[code],
			time.Now().Format(global.DatetimeFormat)),
	}
}

func ErrWithMsg(code int, msg string) Response {
	return Response{
		Code: code,
		Msg: fmt.Sprintf("[%d][%s][%s] %s", code, codeMsgMap[code],
			time.Now().Format(global.DatetimeFormat), msg),
	}
}
