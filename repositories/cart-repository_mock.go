package repositories

import (
	"mini_project/models"
	"fmt"

	"github.com/stretchr/testify/mock"
)

type CartRepositoryMock interface {
	GetCartsRepository() ([]*models.Cart, error)
	GetCartRepository(id string) (*models.Cart, error)
	CreateRepository(Cart models.Cart) (*models.Cart, error)
	UpdateRepository(id string, CartBody models.Cart) (*models.Cart, error)
	DeleteRepository(id string) error
	GetCartByUserRepository(id_user string) ([]*models.Cart, error)
}

type IcartRepositoryMock struct {
	Mock mock.Mock
}

func NewCartRepositoryMock(mock mock.Mock) CartRepositoryMock {
	return &IcartRepositoryMock{
		Mock: mock,
	}
}

func (u *IcartRepositoryMock) GetCartsRepository() ([]*models.Cart, error) {
	args := u.Mock.Called()
	if args.Get(0) == nil {
		return nil, args.Get(1).(error)
	}

	carts := args.Get(0).([]*models.Cart)

	return carts, nil
}

func (u *IcartRepositoryMock) GetCartRepository(id string) (*models.Cart, error) {
	args := u.Mock.Called(id)
	if args.Get(0) == nil {
		return nil, args.Get(1).(error)
	}

	cart := args.Get(0).(models.Cart)

	return &cart, nil
}

func (u *IcartRepositoryMock) CreateRepository(cartData models.Cart) (*models.Cart, error) {
	args := u.Mock.Called(cartData)
	if args.Get(0) == nil {
		return nil, args.Get(1).(error)
	}

	cart := args.Get(0).(models.Cart)

	return &cart, nil
}

func (u *IcartRepositoryMock) UpdateRepository(id string, cartData models.Cart) (*models.Cart, error) {
	args := u.Mock.Called(id, cartData)
	if args.Get(0) == nil {
		return nil,  args.Get(1).(error)
	}

	cart := args.Get(0).(models.Cart)

	return &cart, nil
}

func (u *IcartRepositoryMock) DeleteRepository(id string) error {
	args := u.Mock.Called(id)
	if args.Get(0) != nil {
		return fmt.Errorf("must nil")
	}

	return nil
}

func (u *IcartRepositoryMock) GetCartByUserRepository(id_user string) ([]*models.Cart, error) {
	args := u.Mock.Called(id_user)
	if args.Get(0) == nil {
		return nil, args.Get(0).(error)
	}

	carts := args.Get(0).([]*models.Cart)

	return carts, nil
}