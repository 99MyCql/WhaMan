package sell

import (
	"strconv"

	"WhaMan/app/sell/dto"
	myErr "WhaMan/pkg/error"
	"WhaMan/pkg/log"
	"WhaMan/pkg/rsp"

	"github.com/gin-gonic/gin"
)

var service = new(Service)

// @Summary Create
// @Tags Sell
// @Accept json
// @Param data body dto.ComReq true "出货信息"
// @Success 200 {object} rsp.ResponseExample "code=200"
// @Failure 400 {object} rsp.ResponseExample "code=4xxxxx"
// @Failure 500 {object} rsp.ResponseExample "code=5xxxxx"
// @Router /sell/create [post]
func Create(c *gin.Context) {
	var req *dto.ComReq
	if err := c.ShouldBind(&req); err != nil {
		log.Logger.Error(err)
		c.JSON(rsp.Err(myErr.ParamErr))
		return
	}
	log.Logger.Infof("%+v", req)

	c.JSON(rsp.New(service.Create(req)))
}

// @Summary Get
// @Tags Sell
// @Accept json
// @Param id path uint true "id"
// @Success 200 {object} rsp.ResponseExample "code=200"
// @Failure 400 {object} rsp.ResponseExample "code=4xxxxx"
// @Failure 500 {object} rsp.ResponseExample "code=5xxxxx"
// @Router /sell/get/{id} [get]
func Get(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Logger.Error(err)
		c.JSON(rsp.Err(myErr.ParamErr))
		return
	}
	log.Logger.Info(id)

	c.JSON(rsp.NewWithData(service.Get(uint(id))))
}

// @Summary List
// @Tags Sell
// @Accept json
// @Param data body dto.ListReq true "选项"
// @Success 200 {object} rsp.ResponseExample "code=200"
// @Failure 400 {object} rsp.ResponseExample "code=4xxxxx"
// @Failure 500 {object} rsp.ResponseExample "code=5xxxxx"
// @Router /sell/list [post]
func List(c *gin.Context) {
	var req *dto.ListReq
	if err := c.ShouldBind(&req); err != nil {
		log.Logger.Error(err)
		c.JSON(rsp.Err(myErr.ParamErr))
		return
	}
	log.Logger.Infof("%+v", req)

	c.JSON(rsp.NewWithData(service.List(req)))
}

// @Summary Update
// @Tags Sell
// @Accept json
// @Param id path uint true "id"
// @Param data body dto.ComReq true "出货订单信息"
// @Success 200 {object} rsp.ResponseExample "code=200"
// @Failure 400 {object} rsp.ResponseExample "code=4xxxxx"
// @Failure 500 {object} rsp.ResponseExample "code=5xxxxx"
// @Router /sell/update/{id} [post]
func Update(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Logger.Error(err)
		c.JSON(rsp.Err(myErr.ParamErr))
		return
	}
	log.Logger.Info(id)
	var req *dto.ComReq
	if err := c.ShouldBind(&req); err != nil {
		log.Logger.Error(err)
		c.JSON(rsp.Err(myErr.ParamErr))
		return
	}
	log.Logger.Infof("%+v", req)

	c.JSON(rsp.New(service.Update(uint(id), req)))
}

// @Summary Delete
// @Tags Sell
// @Accept json
// @Param id path uint true "id"
// @Success 200 {object} rsp.ResponseExample "code=200"
// @Failure 400 {object} rsp.ResponseExample "code=4xxxxx"
// @Failure 500 {object} rsp.ResponseExample "code=5xxxxx"
// @Router /sell/delete/{id} [get]
func Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Logger.Error(err)
		c.JSON(rsp.Err(myErr.ParamErr))
		return
	}
	log.Logger.Info(id)

	c.JSON(rsp.New(service.Delete(uint(id))))
}
