package rsp

const (
	UnknownError = -1
	Success      = 0

	// 4XXXXX 客户端错误
	ParamError          = 400001
	UsernamePasswdError = 400002
	NotLoginErr         = 400003

	// 5XXXXX 服务的错误
	CreateFailed  = 500001
	GetFailed     = 500002
	ListFailed    = 500003
	UpdateFailed  = 500004
	DeleteFailed  = 500005
	RestockFailed = 500006
	SellFailed    = 500007
	LoginFailed   = 500008
)

var codeMsgMap = map[int]string{
	-1:     "未知错误",
	0:      "成功",
	400001: "参数错误",
	400002: "用户名或密码错误",
	400003: "未登录",
	500001: "创建失败",
	500002: "查询失败",
	500003: "获取列表失败",
	500004: "更新失败",
	500005: "删除失败",
	500006: "进货失败",
	500007: "出货失败",
	500008: "登录失败",
}
