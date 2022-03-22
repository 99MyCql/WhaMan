package sell

import (
	"testing"
	"time"

	"WhaMan/app/customer"
	customerDTO "WhaMan/app/customer/dto"
	"WhaMan/app/restock"
	restockDTO "WhaMan/app/restock/dto"
	sellDTO "WhaMan/app/sell/dto"
	"WhaMan/app/supplier"
	supplierDTO "WhaMan/app/supplier/dto"
	"WhaMan/pkg/datetime"
	"WhaMan/pkg/test"

	"github.com/stretchr/testify/require"
)

func init() {
	test.Init()
}

func TestCreateAndGetAndUpdateAndDelete(t *testing.T) {
	assert := require.New(t)
	restockService := new(restock.Service)
	supplierService := new(supplier.Service)
	sellService := new(Service)
	customerService := new(customer.Service)

	// 创建供应商
	supplierID, err := supplierService.Create(&supplierDTO.ComReq{
		Name: "A",
	})
	assert.Nil(err)
	defer func() {
		assert.Nil(supplierService.Delete(supplierID))
	}()

	// 创建进货订单
	restockOrderID, err := restockService.Create(&restockDTO.ComReq{
		Date: datetime.MyDatetime{
			Time:  time.Now(),
			Valid: true,
		},
		ModelNum:   "PC",
		Quantity:   100,
		UnitPrice:  10,
		SupplierID: supplierID,
		PaidMoney:  1000,
	})
	assert.Nil(err)
	defer func() {
		assert.Nil(restockService.Delete(restockOrderID))
	}()

	// 创建客户A、B
	customerAID, err := customerService.Create(&customerDTO.ComReq{
		Name: "A",
	})
	assert.Nil(err)
	defer func() {
		assert.Nil(customerService.Delete(customerAID))
	}()
	customerBID, err := customerService.Create(&customerDTO.ComReq{
		Name: "B",
	})
	assert.Nil(err)
	defer func() {
		assert.Nil(customerService.Delete(customerBID))
	}()

	// 创建出货订单
	req := &sellDTO.ComReq{
		Date: datetime.MyDatetime{
			Time:  time.Now(),
			Valid: true,
		},
		CustomerOrderID: "123",
		ModelNum:        "PC",
		Quantity:        50,
		UnitPrice:       13,
		RestockOrderID:  &restockOrderID,
		CustomerID:      customerAID,
	}
	sellOrderID, err := sellService.Create(req)
	assert.Nil(err)
	defer func() {
		assert.Nil(sellService.Delete(sellOrderID))
	}()

	// 验证创建是否成功
	sellOrder, err := sellService.Get(sellOrderID)
	assert.Nil(err)
	assert.Equal("PC", sellOrder.ModelNum)
	assert.Equal(req.Quantity, sellOrder.Quantity)

	// 验证更新
	req.UnitPrice = 12           // 更新单价
	req.RestockOrderID = nil     // 更新库存
	req.CustomerID = customerBID // 更新客户
	assert.Nil(sellService.Update(sellOrderID, req))
	sellOrder, err = sellService.Get(sellOrderID)
	assert.Nil(err)
	t.Logf("%+v", sellOrder)
	assert.Equal(float64(12), sellOrder.UnitPrice)
	assert.Nil(sellOrder.RestockOrderID)
	assert.Equal(customerBID, sellOrder.CustomerID)
}

func TestList(t *testing.T) {
	assert := require.New(t)
	service := new(Service)

	req := &sellDTO.ListReq{
		Where:   nil,
		OrderBy: "date desc",
	}
	sellOrders, err := service.List(req)
	assert.Nil(err)
	t.Logf("%+v", sellOrders[0])
}
