package rsp

import (
	"fmt"
	"time"
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
			time.Now().Format("2006-01-02 15:04:05")),
	}
}

func ErrWithMsg(code int, msg string) Response {
	return Response{
		Code: code,
		Msg: fmt.Sprintf("[%d][%s][%s] %s", code, codeMsgMap[code],
			time.Now().Format("2006-01-02 15:04:05"), msg),
	}
}
