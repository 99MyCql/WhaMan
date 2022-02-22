package customer

import (
	"strconv"

	"WhaMan/app/customer/dto"
	myErr "WhaMan/pkg/error"
	"WhaMan/pkg/log"
	"WhaMan/pkg/rsp"

	"github.com/gin-gonic/gin"
)

var service = new(Service)

// @Summary Create
// @Tags Customer
// @Accept json
// @Param data body dto.ComReq true "客户信息"
// @Success 200 {object} rsp.ResponseExample "code=200"
// @Failure 400 {object} rsp.ResponseExample "code=4xxxxx"
// @Failure 500 {object} rsp.ResponseExample "code=5xxxxx"
// @Router /customer/create [post]
func Create(c *gin.Context) {
	// 解析请求数据
	var req *dto.ComReq
	if err := c.ShouldBind(&req); err != nil {
		log.Logger.Error(err)
		c.JSON(rsp.Err(myErr.ParamErr))
		return
	}

	c.JSON(rsp.NewWithData(service.Create(req)))
}

// @Summary Get
// @Tags Customer
// @Accept json
// @Param id path uint true "id"
// @Success 200 {object} rsp.ResponseExample "code=200"
// @Failure 400 {object} rsp.ResponseExample "code=4xxxxx"
// @Failure 500 {object} rsp.ResponseExample "code=5xxxxx"
// @Router /customer/get/{id} [get]
func Get(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Logger.Error(err)
		c.JSON(rsp.Err(myErr.ParamErr))
		return
	}

	c.JSON(rsp.NewWithData(service.Get(uint(id))))
}

// @Summary List
// @Tags Customer
// @Accept json
// @Success 200 {object} rsp.ResponseExample "code=200"
// @Failure 400 {object} rsp.ResponseExample "code=4xxxxx"
// @Failure 500 {object} rsp.ResponseExample "code=5xxxxx"
// @Router /customer/list [post]
func List(c *gin.Context) {
	c.JSON(rsp.NewWithData(service.List()))
}

// @Summary Update
// @Tags Customer
// @Accept json
// @Param id path uint true "id"
// @Param data body dto.ComReq true "客户信息"
// @Success 200 {object} rsp.ResponseExample "code=200"
// @Failure 400 {object} rsp.ResponseExample "code=4xxxxx"
// @Failure 500 {object} rsp.ResponseExample "code=5xxxxx"
// @Router /customer/update/{id} [post]
func Update(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Logger.Error(err)
		c.JSON(rsp.Err(myErr.ParamErr))
		return
	}
	var req *dto.ComReq
	if err := c.ShouldBind(&req); err != nil {
		log.Logger.Error(err)
		c.JSON(rsp.Err(myErr.ParamErr))
		return
	}

	c.JSON(rsp.New(service.Update(uint(id), req)))
}

// @Summary Delete
// @Tags Customer
// @Accept json
// @Param id path uint true "id"
// @Success 200 {object} rsp.ResponseExample "code=200"
// @Failure 400 {object} rsp.ResponseExample "code=4xxxxx"
// @Failure 500 {object} rsp.ResponseExample "code=5xxxxx"
// @Router /customer/delete/{id} [get]
func Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Logger.Error(err)
		c.JSON(rsp.Err(myErr.ParamErr))
		return
	}

	c.JSON(rsp.New(service.Delete(uint(id))))
}
