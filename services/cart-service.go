package services

import (
	"mini_project/models"
	"mini_project/repositories"
)

type CartService interface {
	GetCartsService() ([]*models.Cart, error)
	GetCartService(id string) (*models.Cart, error)
	CreateService(Cart models.Cart) (*models.Cart, error)
	UpdateService(id string, CartBody models.Cart) (*models.Cart, error)
	DeleteService(id string) error
}

type cartService struct {
	CartR repositories.CartRepository
}

func NewCartService(CartR repositories.CartRepository) CartService {
	return &cartService{
		CartR: CartR,
	}
}

func (a *cartService) GetCartsService() ([]*models.Cart, error) {
	Carts, err := a.CartR.GetCartsRepository()
	if err != nil {
		return nil, err
	}

	return Carts, nil
}

func (a *cartService) GetCartService(id string) (*models.Cart, error) {
	Cart, err := a.CartR.GetCartRepository(id)
	if err != nil {
		return nil, err
	}

	return Cart, nil
}

func (a *cartService) CreateService(Cart models.Cart) (*models.Cart, error) {
	CartR, err := a.CartR.CreateRepository(Cart)
	if err != nil {
		return nil, err
	}

	return CartR, nil
}

func (a *cartService) UpdateService(id string, CartBody models.Cart) (*models.Cart, error) {
	Cart, err := a.CartR.UpdateRepository(id, CartBody)
	if err != nil {
		return Cart, err
	}

	return Cart, nil
}

func (a *cartService) DeleteService(id string) error {
	err := a.CartR.DeleteRepository(id)
	if err != nil {
		return err
	}

	return nil
}
