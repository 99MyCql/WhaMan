package restock

import (
	"testing"
	"time"

	"WhaMan/app/customer"
	customerDTO "WhaMan/app/customer/dto"
	restockDTO "WhaMan/app/restock/dto"
	"WhaMan/app/sell"
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

func TestCreateAndGet(t *testing.T) {
	assert := require.New(t)
	service := new(Service)
	supplierService := new(supplier.Service)

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
	assert.Equal("PC", restockOrder.ModelNum)
	assert.Equal("A", restockOrder.SupplierName) // 验证是否返回供应商名字
}

func TestUpdate(t *testing.T) {
	assert := require.New(t)
	restockService := new(Service)
	supplierService := new(supplier.Service)
	sellService := new(sell.Service)
	customerService := new(customer.Service)

	// 创建供应商
	supplierID, err := supplierService.Create(&supplierDTO.ComReq{
		Name: "A",
	})
	assert.Nil(err)
	defer func() {
		assert.Nil(supplierService.Delete(supplierID))
	}()

	// 创建客户
	customerID, err := customerService.Create(&customerDTO.ComReq{
		Name: "A",
	})
	assert.Nil(err)
	defer func() {
		assert.Nil(customerService.Delete(customerID))
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
		SupplierID: supplierID,
		PaidMoney:  1000,
	}
	restockOrderID, err := restockService.Create(req)
	assert.Nil(err)
	defer func() {
		restockService.Delete(restockOrderID) // 下文会进行删除，不检查是否返回 Nil
	}()

	// 创建出货订单
	sellOrderID, err := sellService.Create(&sellDTO.ComReq{
		Date: datetime.MyDatetime{
			Time:  time.Now(),
			Valid: true,
		},
		CustomerOrderID: "123",
		ModelNum:        "PC",
		Quantity:        50,
		UnitPrice:       13,
		RestockOrderID:  &restockOrderID,
		CustomerID:      customerID,
	})
	defer func() {
		assert.Nil(sellService.Delete(sellOrderID))
	}()

	// 更新进货单价
	req.UnitPrice = 11
	err = restockService.Update(restockOrderID, req)
	assert.Nil(err)
	restockOrder, err := restockService.Get(restockOrderID)
	assert.Nil(err)
	assert.Equal(float64(11), restockOrder.UnitPrice) // 验证进货订单是否更新
	sellOrder, err := sellService.Get(sellOrderID)
	assert.Nil(err)
	assert.Equal(float64(11), sellOrder.RestockUnitPrice) // 验证出货订单是否更新
}

func TestDelete(t *testing.T) {
	assert := require.New(t)
	restockService := new(Service)
	supplierService := new(supplier.Service)
	sellService := new(sell.Service)
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
	req := &restockDTO.ComReq{
		Date: datetime.MyDatetime{
			Time:  time.Now(),
			Valid: true,
		},
		ModelNum:   "PC",
		Quantity:   100,
		UnitPrice:  10,
		SupplierID: supplierID,
		PaidMoney:  1000,
	}
	restockOrderID, err := restockService.Create(req)
	assert.Nil(err)
	defer func() {
		restockService.Delete(restockOrderID) // 下文会进行删除，不检查是否返回 Nil
	}()

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
		CustomerOrderID: "123",
		ModelNum:        "PC",
		Quantity:        50,
		UnitPrice:       13,
		RestockOrderID:  &restockOrderID,
		CustomerID:      customerID,
	})
	assert.Nil(err)
	defer func() {
		sellService.Delete(sellOrderID) // 下文会进行删除，不检查是否返回 Nil
	}()

	// 验证删除
	assert.NotNil(restockService.Delete(restockOrderID)) // 已出货的订单不能删除
	assert.Nil(sellService.Delete(sellOrderID))
	assert.Nil(restockService.Delete(restockOrderID))
}

func TestGet(t *testing.T) {
	assert := require.New(t)
	service := new(Service)

	restockOrder, err := service.Get(185)
	assert.Nil(err)
	t.Logf("%+v", restockOrder)
}

func TestList(t *testing.T) {
	assert := require.New(t)
	service := new(Service)

	minQuantity := float64(1)
	req := &restockDTO.ListReq{
		Where: &restockDTO.Where{
			Date: &restockDTO.Date{
				StartDate: "2021-01-01",
				EndDate:   "",
			},
			SupplierID: 0,
			ModelNum:   "",
			CurQuantity: &restockDTO.CurQuantity{
				Start: &minQuantity,
				End:   nil,
			},
		},
		OrderBy: "",
	}
	restockOrders, err := service.List(req)
	assert.Nil(err)
	t.Logf("%+v", restockOrders[0])
}

func TestListGroupByModelNum(t *testing.T) {
	assert := require.New(t)
	service := new(Service)
	data, err := service.ListGroupByModelNum(&restockDTO.ListGroupByModelNumReq{})
	assert.Nil(err)
	t.Log(data[0])
}
