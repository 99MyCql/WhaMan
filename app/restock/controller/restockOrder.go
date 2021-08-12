package controller

import (
	"net/http"
	"time"

	"WhaMan/app/restock/model"
	"WhaMan/app/restock/service"
	"WhaMan/app/restock/service/impl"
	"WhaMan/pkg/global"
	"WhaMan/pkg/rsp"

	"github.com/gin-gonic/gin"
)

var restockOrderService service.RestockOrder

func init() {
	restockOrderService = new(impl.RestockOrder)
}

type createReq struct {
	Date          time.Time // 日期
	Model         string    // 型号
	Specification string    // 规格
	Quantity      float64   // 数量
	Unit          string    // 单位
	UnitPrice     float64   // 单价
	SumMoney      float64   // 金额
	StockID       string    // 库存号(外键)
	SupplierName  string    // 供应商(外键)
	PaidMoney     float64   // 已付金额
	PayMethod     string    // 付款方式
	Note          string    // 备注
}

// @Summary Create
// @Tags RestockOrder
// @Accept json
// @Param restockOrder body createReq true "restockOrder"
// @Success 200 {string} json "{"code":0,"data":{},"msg":""}"
// @Failure 200 {string} json "{"code":非0,"data":{},"msg":""}"
// @Router /api/v1/restock/order/create [post]
func Create(c *gin.Context) {
	// 定义请求数据结构
	var req createReq
	// 解析请求数据
	if err := c.ShouldBind(&req); err != nil {
		global.Log.Error(err)
		c.JSON(http.StatusOK, rsp.Err(rsp.ParamError))
		return
	}
	global.Log.Debug(req)

	err := restockOrderService.Create(&model.RestockOrder{
		Date:          req.Date,
		Model:         req.Model,
		Specification: req.Specification,
		Quantity:      req.Quantity,
		Unit:          req.Unit,
		UnitPrice:     req.UnitPrice,
		SumMoney:      req.SumMoney,
		StockID:       req.StockID,
		SupplierName:  req.SupplierName,
		PaidMoney:     req.PaidMoney,
		PayMethod:     req.PayMethod,
		Note:          req.Note,
	})

	if err != nil {
		global.Log.Errorf("%+v", err)
		c.JSON(http.StatusOK, rsp.Err(rsp.CreateFailed))
		return
	}

	c.JSON(http.StatusOK, rsp.Suc())
}
