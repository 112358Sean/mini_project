package repositories

import (
	"mini_project/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	GetUsersRepository() ([]*models.User, error)
	GetUserRepository(id string) (*models.User, error)
	CreateRepository(user models.User) (*models.User, error)
	UpdateRepository(id string, userBody models.User) (*models.User, error)
	DeleteRepository(id string) error
	LoginRepository(login models.User) (*models.User, error)
}

type userRepository struct {
	DB *gorm.DB
}

func NewUserRepository(DB *gorm.DB) UserRepository {
	return &userRepository{
		DB: DB,
	}
}

func (u *userRepository) GetUsersRepository() ([]*models.User, error) {
	var users []*models.User

	if err := u.DB.Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

func (u *userRepository) GetUserRepository(id string) (*models.User, error) {
	var user models.User

	if err := u.DB.Where("ID = ?", id).Take(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *userRepository) CreateRepository(user models.User) (*models.User, error) {
	if err := u.DB.Save(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *userRepository) UpdateRepository(id string, userBody models.User) (*models.User, error) {
	user, err := u.GetUserRepository(id)
	if err != nil {
		return nil, err
	}

	err = u.DB.Where("ID = ?", id).Updates(models.User{Nama: userBody.Nama, Email: userBody.Email, Password: userBody.Password, Alamat: userBody.Alamat, Role: userBody.Role}).Error
	if err != nil {
		return nil, err
	}

	user.Nama = userBody.Nama
	user.Email = userBody.Email
	user.Password = userBody.Password
	user.Alamat = userBody.Alamat
	user.Role = userBody.Role

	return user, nil
}

func (u *userRepository) DeleteRepository(id string) error {
	_, err := u.GetUserRepository(id)
	if err != nil {
		return err
	}

	if err := u.DB.Delete(&models.User{}, id).Error; err != nil {
		return err
	}

	return nil
}

func (u *userRepository) LoginRepository(login models.User) (*models.User, error) {
	if err := u.DB.Where("email = ? AND password = ?", login.Email, login.Password).First(&login).Error; err != nil {
		return nil, err
	}

	return &login, nil
}