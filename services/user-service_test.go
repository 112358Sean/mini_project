package services

import (
	"mini_project/models"
	"mini_project/repositories"
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

var (
	userRMock = &repositories.IuserRepositoryMock{Mock: mock.Mock{}}
	userSMock = NewUserService(userRMock)
)

func TestGetUsersService_Success(t *testing.T) {
	usersMP := []*models.User{
		{
			Nama:     "Sean 1",
			Email:    "sean1@gmail.com",
			Password: "123456",
			Alamat: "Jln. 123",
			Role: "admin",
		},
		{
			Nama:     "Sean 1",
			Email:    "sean1@gmail.com",
			Password: "123456",
			Alamat: "Jln. 123",
			Role: "admin",
		},
	}

	usersM := []models.User{
		{
			Nama:     "Sean 1",
			Email:    "sean1@gmail.com",
			Password: "123456",
			Alamat: "Jln. 123",
			Role: "admin",
		},
		{
			Nama:     "Sean 1",
			Email:    "sean1@gmail.com",
			Password: "123456",
			Alamat: "Jln. 123",
			Role: "admin",
		},
	}

	userRMock.Mock.On("GetUsersRepository").Return(usersMP, nil)
	users, err := userSMock.GetUsersService()

	assert.Nil(t, err)
	assert.NotNil(t, users)

	assert.Equal(t, usersM[0].Nama, users[0].Nama)
	assert.Equal(t, usersM[0].Password, users[0].Password)
	assert.Equal(t, usersM[0].Email, users[0].Email)
	assert.Equal(t, usersM[0].Alamat, users[0].Alamat)
	assert.Equal(t, usersM[0].Role, users[0].Role)
}

func TestGetUsersService_Failure(t *testing.T) {
	userRMock = &repositories.IuserRepositoryMock{Mock: mock.Mock{}}
	userSMock = NewUserService(userRMock)
	userRMock.Mock.On("GetUsersRepository").Return(nil, errors.New("get all users failed"))
	users, err := userSMock.GetUsersService()

	assert.Nil(t, users)
	assert.NotNil(t, err)
}

func TestGetUserService_Success(t *testing.T) {
	user := models.User{
		Nama:     "Sean 1",
		Email:    "sean1@gmail.com",
		Password: "123456",
		Alamat: "Jln. 123",
		Role: "admin",
	}

	userRMock.Mock.On("GetUserRepository", "1").Return(user, nil)
	users, err := userSMock.GetUserService("1")

	assert.Nil(t, err)
	assert.NotNil(t, users)

	assert.Equal(t, user.Nama, users.Nama)
	assert.Equal(t, user.Password, users.Password)
	assert.Equal(t, user.Email, users.Email)
	assert.Equal(t, user.Alamat, users.Alamat)
	assert.Equal(t, user.Role, users.Role)
}

func TestGetUserService_Failure(t *testing.T) {
	userRMock.Mock.On("GetUserRepository", "3").Return(nil, fmt.Errorf("user not found"))
	user, err := userSMock.GetUserService("3")

	assert.NotNil(t, err)
	assert.Nil(t, user)
}

func TestCreateUserService_Success(t *testing.T) {
	user := models.User{
		Nama:     "Sean 1",
		Email:    "sean1@gmail.com",
		Password: "123456",
		Alamat: "Jln. 123",
		Role: "admin",
	}

	userRMock.Mock.On("CreateRepository", user).Return(user, nil)
	users, err := userSMock.CreateService(user)

	assert.Nil(t, err)
	assert.NotNil(t, users)

	assert.Equal(t, user.Nama, users.Nama)
	assert.Equal(t, user.Password, users.Password)
	assert.Equal(t, user.Email, users.Email)
	assert.Equal(t, user.Alamat, users.Alamat)
	assert.Equal(t, user.Role, users.Role)
}

func TestCreateUserService_Failure(t *testing.T) {
	user := models.User{
		Nama:     "Sean 1",
		Email:    "sean1@gmail.com",
		Password: "123456",
		Alamat: "Jln. 123",
		Role: "admin",
	}

	userRMock.Mock.On("CreateRepository", user).Return(nil, fmt.Errorf("create user failed"))
	users, err := userSMock.CreateService(user)

	assert.Nil(t, users)
	assert.NotNil(t, err)
}

func TestUpdateUserService_Success(t *testing.T) {
	user := models.User{
		Model: gorm.Model{
			ID: 1,
		},
		Nama:     "Sean 1",
		Email:    "sean1@gmail.com",
		Password: "123456",
		Alamat: "Jln. 123",
		Role: "admin",
	}

	userRMock.Mock.On("UpdateRepository", "1", user).Return(user, nil)
	users, err := userSMock.UpdateService("1", user)

	assert.Nil(t, err)
	assert.NotNil(t, users)

	assert.Equal(t, uint(1), users.ID)
	assert.Equal(t, user.Nama, users.Nama)
	assert.Equal(t, user.Password, users.Password)
	assert.Equal(t, user.Email, users.Email)
	assert.Equal(t, user.Alamat, users.Alamat)
	assert.Equal(t, user.Role, users.Role)
}

func TestUpdateUserService_Failure(t *testing.T) {
	user := models.User{
		Model: gorm.Model{
			ID: 2,
		},
		Nama:     "Sean 1",
		Email:    "sean1@gmail.com",
		Password: "123456",
		Alamat: "Jln. 123",
		Role: "admin",
	}

	userRMock.Mock.On("UpdateRepository", "2", user).Return(nil, fmt.Errorf("user not found"))
	users, err := userSMock.UpdateService("2", user)

	assert.Nil(t, users)
	assert.NotNil(t, err)
}

func TestDeleteUserService_Success(t *testing.T) {
	userRMock.Mock.On("DeleteRepository", "1").Return(nil)
	err := userSMock.DeleteService("1")

	assert.Nil(t, err)
}

func TestDeleteUserService_Failure(t *testing.T) {
	userRMock.Mock.On("DeleteRepository", "2").Return(fmt.Errorf("user not found"))
	err := userSMock.DeleteService("2")

	assert.NotNil(t, err)
}
