package consts

// 返回数据体中code字段的取值，HTTP状态码取code前三位
const (
	// 成功
	SucCode = 200

	// 客户端错误 4XXXXX
	ParamErrCode          = 400001
	UsernamePasswdErrCode = 400002
	FieldDuplicateCode    = 400003
	NotFoundCode          = 400004
	CannotDeleteCode      = 400005
	NotLoginCode          = 400006

	// 服务端错误 5XXXXX
	ServerErrCode = 500001
)

var (
	// CodeMsgMap code对应的消息概要
	CodeMsgMap = map[int]string{
		SucCode:               "成功",
		ParamErrCode:          "参数错误",
		UsernamePasswdErrCode: "用户名或密码错误",
		FieldDuplicateCode:    "字段重复",
		NotFoundCode:          "不存在",
		CannotDeleteCode:      "无法删除",
		NotLoginCode:          "未登录",
		ServerErrCode:         "服务端出错",
	}
)
