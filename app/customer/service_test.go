package customer

import (
	"testing"

	"WhaMan/app/customer/dto"
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
		Name:     "A",
		Phone:    "",
		Contacts: "",
		Note:     "",
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
