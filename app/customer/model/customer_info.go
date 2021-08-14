package model

// CustomerInfo 客户信息结构体，用于HTTP请求等过程中传递数据
type CustomerInfo struct {
	Name     string `binding:"required,excludes= "` // 客户名
	Contacts string `binding:"excludes= "`          // 联系人
	Phone    string `binding:"excludes= ,len=11"`   // 联系电话
	Note     string `binding:"excludes= "`          // 备注
}
