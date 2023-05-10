package services

import (
	"mini_project/models"
	"mini_project/repositories"
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	cartRMock = &repositories.IcartRepositoryMock{Mock: mock.Mock{}}
	cartSMock = NewCartService(cartRMock)
)

func TestGetCartsService_Success(t *testing.T) {
	cartsMP := []*models.Cart{
		{
			ID_User: "1",
			ID_Buku: "1",
			Jumlah: 1,
		},
		{
			ID_User: "1",
			ID_Buku: "1",
			Jumlah: 1,
		},
	}

	cartsM := []models.Cart{
		{
			ID_User: "1",
			ID_Buku: "1",
			Jumlah: 1,
		},
		{
			ID_User: "1",
			ID_Buku: "1",
			Jumlah: 1,
		},
	}

	cartRMock.Mock.On("GetCartsRepository").Return(cartsMP)
	carts, err := cartSMock.GetCartsService()

	assert.Nil(t, err)
	assert.NotNil(t, carts)

	assert.Equal(t, cartsM[0].ID_User, carts[0].ID_User)
	assert.Equal(t, cartsM[0].ID_Buku, carts[0].ID_Buku)
	assert.Equal(t, cartsM[0].Jumlah, carts[0].Jumlah)
}

func TestGetCartsService_Failure(t *testing.T) {
	cartRMock = &repositories.IcartRepositoryMock{Mock: mock.Mock{}}
	cartSMock = NewCartService(cartRMock)
	cartRMock.Mock.On("GetCartsRepository").Return(nil, errors.New("get all carts failed"))
	carts, err := cartSMock.GetCartsService()

	assert.Nil(t, carts)
	assert.NotNil(t, err)
}

func TestGetCartService_Success(t *testing.T) {
	cart := models.Cart{
		ID_User: "1",
		ID_Buku: "1",
		Jumlah: 1,
	}

	cartRMock.Mock.On("GetCartRepository", "1").Return(cart, nil)
	carts, err := cartSMock.GetCartService("1")

	assert.Nil(t, err)
	assert.NotNil(t, carts)

	assert.Equal(t, cart.ID_Buku, carts.ID_Buku)
	assert.Equal(t, cart.ID_User, carts.ID_User)
	assert.Equal(t, cart.Jumlah, carts.Jumlah)
}

func TestGetCartService_Failure(t *testing.T) {
	cartRMock.Mock.On("GetCartRepository", "3").Return(nil, fmt.Errorf("cart not found"))
	cart, err := cartSMock.GetCartService("3")

	assert.NotNil(t, err)
	assert.Nil(t, cart)
}

func TestCreateCartService_Success(t *testing.T) {
	cart := models.Cart{
		ID_User: "1",
		ID_Buku: "1",
		Jumlah: 1,
	}

	cartRMock.Mock.On("CreateRepository", cart).Return(cart, nil)
	carts, err := cartSMock.CreateService(cart)

	assert.Nil(t, err)
	assert.NotNil(t, carts)

	assert.Equal(t, cart.ID_User, carts.ID_User)
	assert.Equal(t, cart.ID_Buku, carts.ID_Buku)
	assert.Equal(t, cart.Jumlah, carts.Jumlah)
}

func TestCreateCartService_Failure(t *testing.T) {
	cart := models.Cart{
		ID_User: "1",
		ID_Buku: "1",
		Jumlah: 1,
	}

	cartRMock.Mock.On("CreateRepository", cart).Return(nil, fmt.Errorf("create cart failed"))
	carts, err := cartSMock.CreateService(cart)

	assert.Nil(t, carts)
	assert.NotNil(t, err)
}

func TestUpdateCartService_Success(t *testing.T) {
	cart := models.Cart{
		ID_User: "1",
		ID_Buku: "1",
		Jumlah: 1,
	}

	cartRMock.Mock.On("UpdateRepository", "1", cart).Return(cart, nil)
	carts, err := cartSMock.UpdateService("1", cart)

	assert.Nil(t, err)
	assert.NotNil(t, carts)

	assert.Equal(t, cart.ID, carts.ID)
	assert.Equal(t, cart.ID_User, carts.ID_User)
	assert.Equal(t, cart.ID_Buku, carts.ID_Buku)
	assert.Equal(t, cart.Jumlah, carts.Jumlah)
}

func TestUpdateCartService_Failure(t *testing.T) {
	cart := models.Cart{
		ID_User: "1",
		ID_Buku: "1",
		Jumlah: 1,
	}

	cartRMock.Mock.On("UpdateRepository", "2", cart).Return(nil, fmt.Errorf("cart not found"))
	carts, err := cartSMock.UpdateService("2", cart)

	assert.Nil(t, carts)
	assert.NotNil(t, err)
}

func TestDeleteCartService_Success(t *testing.T) {
	cartRMock.Mock.On("DeleteRepository", "1").Return(nil)
	err := cartSMock.DeleteService("1")

	assert.Nil(t, err)
}

func TestDeleteCartService_Failure(t *testing.T) {
	cartRMock.Mock.On("DeleteRepository", "2").Return(fmt.Errorf("cart not found"))
	err := cartSMock.DeleteService("2")

	assert.NotNil(t, err)
}
