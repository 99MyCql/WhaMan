package customer

import (
	"testing"
	"time"

	"WhaMan/app/customer/dto"
	"WhaMan/app/sell"
	sellDTO "WhaMan/app/sell/dto"
	"WhaMan/pkg/datetime"
	"WhaMan/pkg/test"

	"github.com/stretchr/testify/require"
)

func init() {
	test.Init()
}

func TestCreate(t *testing.T) {
	assert := require.New(t)
	service := new(Service)
	req := &dto.ComReq{
		Name: "A",
	}
	id, err := service.Create(req)
	assert.Nil(err)
	t.Log(id)
	defer func() { // 不使用 func(){}() 形式， assert 会被直接调用
		assert.Nil(service.Delete(id))
	}()

	// 名称重复应返回错误
	_, err = service.Create(req)
	assert.NotNil(err)

	// 检查
	customer, err := service.Get(id)
	assert.Nil(err)
	t.Log(customer)
	assert.Equal(customer.Name, "A")
}

func TestUpdate(t *testing.T) {
	assert := require.New(t)
	service := new(Service)
	req1 := &dto.ComReq{
		Name: "A",
	}
	req2 := &dto.ComReq{
		Name: "B",
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
	sellService := new(sell.Service)

	req := &dto.ComReq{
		Name: "A",
	}
	id, err := service.Create(req)
	assert.Nil(err)
	t.Log(id)
	defer func() {
		assert.Nil(service.Delete(id))
	}()

	sellOrderID, err := sellService.Create(&sellDTO.ComReq{
		Date: datetime.MyDatetime{
			Time:  time.Now(),
			Valid: true,
		},
		ModelNum:         "PC",
		Quantity:         10000,
		UnitPrice:        1,
		PaidMoney:        0,
		RestockUnitPrice: 0.9,
		RestockOrderID:   nil,
		CustomerID:       id,
	})
	assert.Nil(err)
	defer func() {
		assert.Nil(sellService.Delete(sellOrderID))
	}()

	customer, err := service.Get(id)
	assert.Nil(err)
	t.Logf("%+v", customer)
	assert.Equal(1, len(customer.SellOrders))
	t.Logf("%+v", customer.SellOrders[0])
}

func TestList(t *testing.T) {
	assert := require.New(t)
	service := new(Service)
	sellService := new(sell.Service)

	req := &dto.ComReq{
		Name: "A",
	}
	id, err := service.Create(req)
	assert.Nil(err)
	t.Log(id)
	defer func() {
		assert.Nil(service.Delete(id))
	}()

	sellOrderAID, err := sellService.Create(&sellDTO.ComReq{
		Date: datetime.MyDatetime{
			Time:  time.Now(),
			Valid: true,
		},
		ModelNum:         "PC",
		Quantity:         10000,
		UnitPrice:        1,
		PaidMoney:        0,
		RestockUnitPrice: 0.9,
		RestockOrderID:   nil,
		CustomerID:       id,
	})
	assert.Nil(err)
	defer func() {
		assert.Nil(sellService.Delete(sellOrderAID))
	}()
	sellOrderBID, err := sellService.Create(&sellDTO.ComReq{
		Date: datetime.MyDatetime{
			Time:  time.Now(),
			Valid: true,
		},
		ModelNum:         "PC",
		Quantity:         20000,
		UnitPrice:        1,
		PaidMoney:        10000,
		RestockUnitPrice: 0.8,
		OtherCost:        100,
		RestockOrderID:   nil,
		CustomerID:       id,
	})
	assert.Nil(err)
	defer func() {
		assert.Nil(sellService.Delete(sellOrderBID))
	}()

	customers, err := service.List(&dto.ListReq{
		Where:   nil,
		OrderBy: "",
		SellOrdersWhere: &dto.SellOrdersWhere{
			Date: &dto.Date{
				StartDate: "2021-01-01",
				EndDate:   "",
			},
		},
	})
	assert.Nil(err)
	t.Logf("%+v", customers[0])
}
