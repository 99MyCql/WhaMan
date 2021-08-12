package rsp

const (
	UnknownError = -1
	Success      = 0

	// 4XXXXX 客户端错误
	ParamError = 400001

	// 5XXXXX 服务的错误
	CreateFailed = 500001
)

var codeMsgMap = map[int]string{
	-1:     "未知错误",
	0:      "成功",
	400001: "参数错误",
	500001: "创建失败",
}
