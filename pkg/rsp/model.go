package rsp

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
		Msg:  codeMsgMap[code],
	}
}

func ErrWithData(code int, data interface{}) Response {
	return Response{
		Code: code,
		Msg:  codeMsgMap[code],
		Data: data,
	}
}
