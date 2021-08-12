package controller

import (
	"net/http"

	"WhaMan/app/customer/model"
	"WhaMan/app/customer/service"
	"WhaMan/app/customer/service/impl"
	"WhaMan/pkg/global"
	"WhaMan/pkg/rsp"

	"github.com/gin-gonic/gin"
)

var customerService service.Customer

func init() {
	customerService = new(impl.CustomerImpl)
}

type createReq struct {
	Name        string  `json:"name" binding:"required,excludes= "` // 客户名
	Contacts    string  `json:"contacts" binding:"excludes= "`      // 联系人
	Phone       string  `json:"phone" binding:"excludes= "`         // 联系电话
	Turnover    float64 `json:"turnover" binding:"excludes= "`      // 交易额
	UnpaidMoney float64 `json:"unpaid_money" binding:"excludes= "`  // 未付款
	Note        string  `json:"note" binding:"excludes= "`          // 备注
}

// @Summary Create
// @Tags Customer
// @Accept json
// @Param customer body createReq true "customer"
// @Success 200 {string} json "{"code":0,"data":{},"msg":""}"
// @Failure 200 {string} json "{"code":非0,"data":{},"msg":""}"
// @Router /api/v1/customer/create [post]
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

	err := customerService.Create(&model.Customer{
		Name:        req.Name,
		Contacts:    req.Contacts,
		Phone:       req.Phone,
		Turnover:    req.Turnover,
		UnpaidMoney: req.UnpaidMoney,
		Note:        req.Note,
	})
	if err != nil {
		global.Log.Errorf("%+v", err)
		c.JSON(http.StatusOK, rsp.Err(rsp.CreateFailed))
		return
	}

	c.JSON(http.StatusOK, rsp.Suc())
}
