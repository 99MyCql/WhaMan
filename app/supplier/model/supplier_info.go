package model

type SupplierInfo struct {
	Name     string `binding:"required,excludes= "` // 供应商名
	Contacts string `binding:"excludes= "`          // 联系人
	Phone    string `binding:"excludes= "`          // 联系电话
	Note     string `binding:"excludes= "`          // 备注
}
