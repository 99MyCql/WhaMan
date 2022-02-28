package sell

import (
	"testing"
	"time"

	"WhaMan/app/customer"
	customerDTO "WhaMan/app/customer/dto"
	"WhaMan/app/restock"
	restockDTO "WhaMan/app/restock/dto"
	sellDTO "WhaMan/app/sell/dto"
	"WhaMan/app/stock"
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
	stockService := new(stock.Service)

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
	restockOrder, err := restockService.Get(restockOrderID)
	assert.Nil(err)

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
		CustomerOrderID:  "123",
		ModelNum:         "PC",
		Quantity:         50,
		RestockUnitPrice: restockOrder.UnitPrice,
		UnitPrice:        13,
		StockID:          &restockOrder.StockID,
		CustomerID:       customerAID,
	}
	sellOrderID, err := sellService.Create(req)
	assert.Nil(err)
	defer func() {
		sellService.Delete(sellOrderID) // 下文会进行删除，此处不检查是否返回 nil
	}()

	// 验证创建是否成功
	sellOrder, err := sellService.Get(sellOrderID)
	assert.Nil(err)
	assert.Equal("PC", sellOrder.ModelNum)
	assert.Equal(float64(50), sellOrder.Quantity)

	// 验证库存是否更新
	_stock, err := stockService.Get(restockOrder.StockID)
	assert.Nil(err)
	assert.Equal(float64(50), _stock.CurQuantity)

	// 验证客户交易额是否更新
	_customer, err := customerService.Get(customerAID)
	assert.Nil(err)
	assert.Equal(float64(650), _customer.Turnover)

	// 验证更新
	req.UnitPrice = 12
	assert.Nil(sellService.Update(sellOrderID, req))
	sellOrder, err = sellService.Get(sellOrderID)
	assert.Nil(err)
	assert.Equal(float64(12), sellOrder.UnitPrice)

	// 更新库存
	req.StockID = nil
	assert.Nil(sellService.Update(sellOrderID, req))
	sellOrder, err = sellService.Get(sellOrderID)
	assert.Nil(err)
	assert.Nil(sellOrder.StockID)
	_stock, err = stockService.Get(restockOrder.StockID)
	assert.Nil(err)
	assert.Equal(float64(100), _stock.CurQuantity)

	// 更新客户
	req.CustomerID = customerBID
	assert.Nil(sellService.Update(sellOrderID, req))
	sellOrder, err = sellService.Get(sellOrderID)
	assert.Nil(err)
	assert.Equal(customerBID, sellOrder.CustomerID)
	_customer, err = customerService.Get(customerAID)
	assert.Nil(err)
	assert.Equal(float64(0), _customer.Turnover)
	_customer, err = customerService.Get(customerBID)
	assert.Nil(err)
	assert.Equal(float64(12*50), _customer.Turnover)

	// 验证删除
	assert.Nil(sellService.Delete(sellOrderID))
	_customer, err = customerService.Get(customerBID)
	assert.Nil(err)
	assert.Equal(float64(0), _customer.Turnover)
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
