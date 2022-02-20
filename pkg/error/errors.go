package error

var (
	// 4XXXXX 客户端错误
	ParamErr = &Error{
		Code: 400001,
		Msg:  "参数错误",
	}
	UsernamePasswdErr = &Error{
		Code: 400002,
		Msg:  "用户名或密码错误",
	}
	FieldDuplicate = &Error{
		Code: 400003,
		Msg:  "字段重复",
	}
	NotFound = &Error{
		Code: 400004,
		Msg:  "不存在",
	}
	CannotDelete = &Error{
		Code: 400005,
		Msg:  "无法删除",
	}
	NotLogin = &Error{
		Code: 400006,
		Msg:  "未登录",
	}

	// 5XXXXX 服务端错误
	ServerErr = &Error{
		Code: 500001,
		Msg:  "服务端出错",
	}
)
