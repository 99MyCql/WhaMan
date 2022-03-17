package supplier

import (
	"testing"
	"time"

	"WhaMan/app/restock"
	restockDTO "WhaMan/app/restock/dto"
	"WhaMan/app/supplier/dto"
	"WhaMan/pkg/datetime"
	"WhaMan/pkg/test"

	"github.com/stretchr/testify/require"
)

func init() {
	test.Init()
}

func TestCreate(t *testing.T) {
	assert := require.New(t) // require 碰到错误直接终止；assert 碰到错误继续
	service := new(Service)
	req := &dto.ComReq{
		Name:     "A",
		Phone:    "",
		Contacts: "",
		Note:     "",
	}
	id, err := service.Create(req)
	assert.Nil(err)
	defer func() {
		assert.Nil(service.Delete(id))
	}()
	_, err = service.Create(req)
	assert.NotNil(err) // 名称重复应报错

	supplier, err := service.Get(id)
	assert.Nil(err)
	t.Log(supplier)
	assert.Equal(supplier.Name, "A")
}

func TestUpdate(t *testing.T) {
	assert := require.New(t)
	service := new(Service)
	req1 := &dto.ComReq{
		Name:     "A",
		Phone:    "",
		Contacts: "",
		Note:     "",
	}
	req2 := &dto.ComReq{
		Name:     "B",
		Phone:    "",
		Contacts: "",
		Note:     "",
	}
	id1, err := service.Create(req1)
	assert.Nil(err)
	defer func() {
		assert.Nil(service.Delete(id1))
	}()
	id2, err := service.Create(req2)
	assert.Nil(err)
	defer func() {
		assert.Nil(service.Delete(id2))
	}()

	req1.Phone = "123"
	assert.Nil(service.Update(id1, req1))
	customer, err := service.Get(id1)
	assert.Nil(err)
	assert.Equal(customer.Phone, "123")

	// 不能更新为已存在的名字
	req2.Name = "A"
	assert.NotNil(service.Update(id2, req2))
}

func TestGet(t *testing.T) {
	assert := require.New(t)
	service := new(Service)
	restockService := new(restock.Service)
	supplierID, err := service.Create(&dto.ComReq{
		Name: "A",
	})
	assert.Nil(err)
	defer func() {
		assert.Nil(service.Delete(supplierID))
	}()

	// 创建进货订单
	restockOrderID, err := restockService.Create(&restockDTO.ComReq{
		Date: datetime.MyDatetime{
			Time:  time.Now(),
			Valid: true,
		},
		ModelNum:      "PC",
		Specification: "PC",
		Quantity:      100,
		UnitPrice:     13,
		SupplierID:    supplierID,
		PaidMoney:     1300,
	})
	assert.Nil(err)
	defer func() {
		assert.Nil(restockService.Delete(restockOrderID))
	}()

	// 验证
	supplier, err := service.Get(supplierID)
	assert.Nil(err)
	t.Logf("%+v", supplier)
	assert.Equal(supplier.Name, "A")
	// 验证是否返回关联的进货订单
	assert.NotNil(supplier.RestockOrders)
	assert.Equal(len(supplier.RestockOrders), 1)
	t.Logf("%+v", supplier.RestockOrders[0])
	assert.Equal(supplier.RestockOrders[0].ID, restockOrderID)
}

func TestList(t *testing.T) {
	assert := require.New(t)
	service := new(Service)
	restockService := new(restock.Service)
	supplierID, err := service.Create(&dto.ComReq{
		Name: "A",
	})
	assert.Nil(err)
	defer func() {
		assert.Nil(service.Delete(supplierID))
	}()

	// 创建进货订单
	restockOrderID, err := restockService.Create(&restockDTO.ComReq{
		Date: datetime.MyDatetime{
			Time:  time.Now(),
			Valid: true,
		},
		ModelNum:      "PC",
		Specification: "PC",
		Quantity:      100,
		UnitPrice:     13,
		SupplierID:    supplierID,
		PaidMoney:     1300,
	})
	assert.Nil(err)
	defer func() {
		assert.Nil(restockService.Delete(restockOrderID))
	}()

	suppliers, err := service.List(&dto.ListReq{
		Where:              nil,
		OrderBy:            "",
		RestockOrdersWhere: nil,
	})
	assert.Nil(err)
	t.Logf("%+v", suppliers[0])
}
