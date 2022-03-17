package supplier

import (
	"strconv"

	"WhaMan/app/supplier/dto"
	myErr "WhaMan/pkg/error"
	"WhaMan/pkg/log"
	"WhaMan/pkg/rsp"

	"github.com/gin-gonic/gin"
)

var service = new(Service)

// @Summary Create
// @Tags Supplier
// @Accept json
// @Param data body dto.ComReq true "供应商信息"
// @Success 200 {object} rsp.ResponseExample "code=200"
// @Failure 400 {object} rsp.ResponseExample "code=4xxxxx"
// @Failure 500 {object} rsp.ResponseExample "code=5xxxxx"
// @Router /supplier/create [post]
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
// @Tags Supplier
// @Accept json
// @Param id path uint true "id"
// @Success 200 {object} rsp.ResponseExample "code=200"
// @Failure 400 {object} rsp.ResponseExample "code=4xxxxx"
// @Failure 500 {object} rsp.ResponseExample "code=5xxxxx"
// @Router /supplier/get/{id} [get]
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
// @Tags Supplier
// @Accept json
// @Param data body dto.ListReq true "请求参数"
// @Success 200 {object} rsp.ResponseExample "code=200"
// @Failure 400 {object} rsp.ResponseExample "code=4xxxxx"
// @Failure 500 {object} rsp.ResponseExample "code=5xxxxx"
// @Router /supplier/list [post]
func List(c *gin.Context) {
	var req *dto.ListReq
	if err := c.ShouldBind(&req); err != nil {
		log.Logger.Error(err)
		c.JSON(rsp.Err(myErr.ParamErr))
		return
	}
	c.JSON(rsp.NewWithData(service.List(req)))
}

// @Summary Update
// @Tags Supplier
// @Accept json
// @Param id path uint true "id"
// @Param data body dto.ComReq true "供应商信息"
// @Success 200 {object} rsp.ResponseExample "code=200"
// @Failure 400 {object} rsp.ResponseExample "code=4xxxxx"
// @Failure 500 {object} rsp.ResponseExample "code=5xxxxx"
// @Router /supplier/update/{id} [post]
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
// @Tags Supplier
// @Accept json
// @Param id path uint true "id"
// @Success 200 {object} rsp.ResponseExample "code=200"
// @Failure 400 {object} rsp.ResponseExample "code=4xxxxx"
// @Failure 500 {object} rsp.ResponseExample "code=5xxxxx"
// @Router /supplier/delete/{id} [get]
func Delete(c *gin.Context) {
	// 解析请求数据
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Logger.Error(err)
		c.JSON(rsp.Err(myErr.ParamErr))
		return
	}

	c.JSON(rsp.New(service.Delete(uint(id))))
}
