package model

import "time"

// RestockInfo 进货信息，用于HTTP请求传递数据
type RestockInfo struct {
	DateStr       string    `json:"date" binding:"required,excludes= "` // 日期(字符串形式)
	Date          time.Time `binding:"-"`                               // 日期
	ModelNum      string    `binding:"required,excludes= "`             // 型号
	Specification string    `binding:"required,excludes= "`             // 规格
	Quantity      float64   `binding:"required"`                        // 数量
	Unit          string    `binding:"excludes= "`                      // 单位
	UnitPrice     float64   `binding:"required"`                        // 单价
	SumMoney      float64   `binding:"required"`                        // 金额
	SupplierID    uint      `binding:"required"`                        // 供应商(外键)
	PaidMoney     float64   `binding:"required"`                        // 已付金额
	PayMethod     string    `binding:"excludes= "`                      // 付款方式
	Note          string    `binding:"excludes= "`                      // 备注
	Location      string    `binding:"excludes= "`                      // 存放地点
}
