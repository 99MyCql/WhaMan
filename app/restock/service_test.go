package restock

import (
	"testing"
	"time"

	"WhaMan/app/customer"
	customerDTO "WhaMan/app/customer/dto"
	restockDTO "WhaMan/app/restock/dto"
	"WhaMan/app/sell"
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

func TestCreateAndGet(t *testing.T) {
	assert := require.New(t)
	service := new(Service)
	supplierService := new(supplier.Service)
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
	restockOrderID, err := service.Create(&restockDTO.ComReq{
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
		assert.Nil(service.Delete(restockOrderID))
	}()

	// 验证进货订单是否创建成功
	restockOrder, err := service.Get(restockOrderID)
	assert.Nil(err)
	t.Logf("%+v", restockOrder)
	assert.Equal(restockOrder.ModelNum, "PC")
	assert.Equal(restockOrder.SupplierName, "A") // 验证是否返回供应商名字

	// 验证库存是否创建成功
	_stock, err := stockService.Get(restockOrder.StockID)
	assert.Nil(err)
	t.Logf("%+v", _stock)
	assert.Equal(_stock.CurQuantity, float64(100))

	// 验证是否更新供应商交易额
	_supplier, err := supplierService.Get(supplierID)
	assert.Nil(err)
	t.Logf("%+v", _supplier)
	assert.Equal(_supplier.Turnover, float64(1000))
}

func TestUpdateAndDelete(t *testing.T) {
	assert := require.New(t)
	restockService := new(Service)
	supplierService := new(supplier.Service)
	sellService := new(sell.Service)
	customerService := new(customer.Service)
	stockService := new(stock.Service)

	// 创建供应商
	supplierAID, err := supplierService.Create(&supplierDTO.ComReq{
		Name: "A",
	})
	assert.Nil(err)
	defer func() {
		assert.Nil(supplierService.Delete(supplierAID))
	}()
	// 创建另一个供应商
	supplierBID, err := supplierService.Create(&supplierDTO.ComReq{
		Name: "B",
	})
	assert.Nil(err)
	defer func() {
		assert.Nil(supplierService.Delete(supplierBID))
	}()

	// 创建进货订单
	req := &restockDTO.ComReq{
		Date: datetime.MyDatetime{
			Time:  time.Now(),
			Valid: true,
		},
		ModelNum:   "PC",
		Quantity:   100,
		UnitPrice:  10,
		SupplierID: supplierAID,
		PaidMoney:  1000,
	}
	restockOrderID, err := restockService.Create(req)
	assert.Nil(err)
	defer func() {
		restockService.Delete(restockOrderID) // 下文会进行删除，不检查是否返回 Nil
	}()
	restockOrder, err := restockService.Get(restockOrderID)
	assert.Nil(err)
	t.Logf("%+v", restockOrder)

	// 创建客户
	customerID, err := customerService.Create(&customerDTO.ComReq{
		Name: "A",
	})
	assert.Nil(err)
	defer func() {
		assert.Nil(customerService.Delete(customerID))
	}()

	// 创建出货订单
	sellOrderID, err := sellService.Create(&sellDTO.ComReq{
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
		CustomerID:       customerID,
	})
	assert.Nil(err)
	defer func() {
		sellService.Delete(sellOrderID) // 下文会进行删除，不检查是否返回 Nil
	}()

	// 更新进货数量
	req.Quantity = 80
	err = restockService.Update(restockOrderID, req)
	assert.Nil(err)
	_stock, err := stockService.Get(restockOrder.StockID)
	assert.Nil(err)
	t.Logf("%+v", _stock)
	assert.Equal(_stock.CurQuantity, float64(30)) // 验证库存数量是否更新

	// 更新进货单价
	req.UnitPrice = 11
	err = restockService.Update(restockOrderID, req)
	assert.Nil(err)
	sellOrder, err := sellService.Get(sellOrderID)
	assert.Nil(err)
	assert.Equal(sellOrder.RestockUnitPrice, float64(11))
	assert.Equal(sellOrder.Profit, float64(50*(13-11))) // 验证出货订单利润是否更新

	// 更换供应商
	req.SupplierID = supplierBID
	err = restockService.Update(restockOrderID, req)
	assert.Nil(err)
	supplierA, err := supplierService.Get(supplierAID)
	assert.Nil(err)
	assert.Equal(supplierA.Turnover, float64(0)) // 验证旧供应商交易额是否更新
	supplierB, err := supplierService.Get(supplierBID)
	assert.Nil(err)
	assert.Equal(supplierB.Turnover, float64(80*11)) // 验证新供应商交易额是否更新

	// 验证删除
	assert.NotNil(restockService.Delete(restockOrderID)) // 已出货的订单不能删除
	assert.Nil(sellService.Delete(sellOrderID))
	assert.Nil(restockService.Delete(restockOrderID))
	supplierB, err = supplierService.Get(supplierBID)
	assert.Nil(err)
	assert.Equal(supplierB.Turnover, float64(0)) // 供应商交易额应为0
}
