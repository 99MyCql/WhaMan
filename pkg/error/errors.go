package error

import "WhaMan/consts"

var (
	// 4XXXXX 客户端错误
	ParamErr = &Error{
		Code: consts.ParamErrCode,
		Msg:  consts.CodeMsgMap[consts.ParamErrCode],
	}
	UsernamePasswdErr = &Error{
		Code: consts.UsernamePasswdErrCode,
		Msg:  consts.CodeMsgMap[consts.UsernamePasswdErrCode],
	}
	FieldDuplicate = &Error{
		Code: consts.FieldDuplicateCode,
		Msg:  consts.CodeMsgMap[consts.FieldDuplicateCode],
	}
	NotFound = &Error{
		Code: consts.NotFoundCode,
		Msg:  consts.CodeMsgMap[consts.NotFoundCode],
	}
	CannotDelete = &Error{
		Code: consts.CannotDeleteCode,
		Msg:  consts.CodeMsgMap[consts.CannotDeleteCode],
	}
	NotLogin = &Error{
		Code: consts.NotLoginCode,
		Msg:  consts.CodeMsgMap[consts.NotLoginCode],
	}

	// 5XXXXX 服务端错误
	ServerErr = &Error{
		Code: consts.ServerErrCode,
		Msg:  consts.CodeMsgMap[consts.ServerErrCode],
	}
)
