package repositories

import (
	"mini_project/models"

	"gorm.io/gorm"
)

type CartRepository interface {
	GetCartsRepository() ([]*models.Cart, error)
	GetCartRepository(id string) (*models.Cart, error)
	CreateRepository(Cart models.Cart) (*models.Cart, error)
	UpdateRepository(id string, CartBody models.Cart) (*models.Cart, error)
	DeleteRepository(id string) error
	GetCartByUserRepository(id_user string) ([]*models.Cart, error)
}

type cartRepository struct {
	DB *gorm.DB
}

func NewCartRepository(DB *gorm.DB) CartRepository {
	return &cartRepository{
		DB: DB,
	}
}

func (a *cartRepository) GetCartsRepository() ([]*models.Cart, error) {
	var Carts []*models.Cart

	if err := a.DB.Find(&Carts).Error; err != nil {
		return nil, err
	}

	return Carts, nil
}

func (a *cartRepository) GetCartRepository(id string) (*models.Cart, error) {
	var Cart *models.Cart

	if err := a.DB.Where("id = ?", id).Take(&Cart).Error; err != nil {
		return nil, err
	}

	return Cart, nil
}

func (a *cartRepository) CreateRepository(Cart models.Cart) (*models.Cart, error) {
	if err := a.DB.Save(&Cart).Error; err != nil {
		return nil, err
	}

	return &Cart, nil
}

func (a *cartRepository) UpdateRepository(id string, CartBody models.Cart) (*models.Cart, error) {
	Cart, err := a.GetCartRepository(id)
	if err != nil {
		return nil, err
	}

	err = a.DB.Where("ID = ?", id).Updates(models.Cart{ID_User: CartBody.ID_User, ID_Buku: CartBody.ID_Buku, Jumlah: CartBody.Jumlah}).Error
	if err != nil {
		return nil, err
	}

	Cart.ID_User = CartBody.ID_User
	Cart.ID_Buku = CartBody.ID_Buku
	Cart.Jumlah = CartBody.Jumlah

	return Cart, nil
}

func (a *cartRepository) DeleteRepository(id string) error {
	_, err := a.GetCartRepository(id)
	if err != nil {
		return err
	}

	if err := a.DB.Delete(&models.Cart{}, id).Error; err != nil {
		return err
	}

	return nil
}

func (a *cartRepository) GetCartByUserRepository(id_user string) ([]*models.Cart, error) {
    var Carts []*models.Cart

    if err := a.DB.Where("id_user = ?", id_user).Find(&Carts).Error; err != nil {
        return nil, err
    }
    return Carts, nil
}
