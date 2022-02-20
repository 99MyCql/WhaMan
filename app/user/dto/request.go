package dto

// LoginReq Login 接口请求参数
type LoginReq struct {
	Username string `binding:"required"`
	Password string `binding:"required"`
}
