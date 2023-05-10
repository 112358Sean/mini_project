package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"

	h "mini_project/helpers"
	m "mini_project/middlewares"
	"mini_project/models"
	"mini_project/services"
)

type UserController interface {
	GetUsersController(c echo.Context) error
	GetUserController(c echo.Context) error
	CreateController(c echo.Context) error
	UpdateController(c echo.Context) error
	DeleteController(c echo.Context) error
	LoginController(c echo.Context) error
}

type userController struct {
	userS services.UserService
	jwt m.JWTS
}

func NewUserController(userS services.UserService, jwtS m.JWTS) UserController {
	return &userController{
		userS: userS,
		jwt: jwtS,
	}
}

func (u *userController) GetUsersController(c echo.Context) error {
	users, err := u.userS.GetUsersService()
	if err != nil {
		return h.Response(c, http.StatusBadRequest, h.ResponseModel{
			Data:    nil,
			Message: err.Error(),
			Status:  false,
		})
	}

	for _,user := range users {
		user.Password = "-"
	}

	return h.Response(c, http.StatusOK, h.ResponseModel{
		Data:    users,
		Message: "Get all users success",
		Status:  true,
	})
}

func (u *userController) GetUserController(c echo.Context) error {
	id := c.Param("id")

	err := h.IsNumber(id)
	if err != nil {
		return h.Response(c, http.StatusBadRequest, h.ResponseModel{
			Data:    nil,
			Message: err.Error(),
			Status:  false,
		})
	}

	var user *models.User

	user, err = u.userS.GetUserService(id)
	if err != nil {
		return h.Response(c, http.StatusNotFound, h.ResponseModel{
			Data:    nil,
			Message: err.Error(),
			Status:  false,
		})
	}

	user.Password = "-"

	return h.Response(c, http.StatusOK, h.ResponseModel{
		Data:    user,
		Message: "Get user success",
		Status:  true,
	})
}

func (u *userController) CreateController(c echo.Context) error {
	var user models.CreateUser

	err := c.Bind(&user.User)
	if err != nil {
		return h.Response(c, http.StatusBadRequest, h.ResponseModel{
			Data:    nil,
			Message: err.Error(),
			Status:  false,
		})
	}

	user.User, err = u.userS.CreateService(*user.User)
	if err != nil {
		return h.Response(c, http.StatusBadRequest, h.ResponseModel{
			Data:    nil,
			Message: err.Error(),
			Status:  false,
		})
	}

	token, err := u.jwt.CreateJWTToken(user.User.ID, user.User.Nama)
	if err != nil {
		return h.Response(c, http.StatusBadRequest, h.ResponseModel{
			Data:    nil,
			Message: err.Error(),
			Status:  false,
		})
	}

	user.User.Password = "-"

	user.Token = token
	return h.Response(c, http.StatusOK, h.ResponseModel{
		Data:    user,
		Message: "Create user success",
		Status:  true,
	})
}

func (u *userController) UpdateController(c echo.Context) error {
	id := c.Param("id")

	err := h.IsNumber(id)
	if err != nil {
		return h.Response(c, http.StatusBadRequest, h.ResponseModel{
			Data:    nil,
			Message: err.Error(),
			Status:  false,
		})
	}

	var user *models.User

	err = c.Bind(&user)
	if err != nil {
		return h.Response(c, http.StatusBadRequest, h.ResponseModel{
			Data:    nil,
			Message: err.Error(),
			Status:  false,
		})
	}

	user, err = u.userS.UpdateService(id, *user)
	if err != nil {
		return h.Response(c, http.StatusBadRequest, h.ResponseModel{
			Data:    nil,
			Message: err.Error(),
			Status:  false,
		})
	}

	user.Password = "-"

	return h.Response(c, http.StatusOK, h.ResponseModel{
		Data:    user,
		Message: "Update user success",
		Status:  true,
	})
}

func (u *userController) DeleteController(c echo.Context) error {
	id := c.Param("id")

	err := h.IsNumber(id)
	if err != nil {
		return h.Response(c, http.StatusBadRequest, h.ResponseModel{
			Data:    nil,
			Message: err.Error(),
			Status:  false,
		})
	}

	err = u.userS.DeleteService(id)
	if err != nil {
		return h.Response(c, http.StatusBadRequest, h.ResponseModel{
			Data:    nil,
			Message: err.Error(),
			Status:  false,
		})
	}

	return h.Response(c, http.StatusOK, h.ResponseModel{
		Data:    nil,
		Message: "Delete user success",
		Status:  true,
	})
}

func (u *userController) LoginController(c echo.Context) error {
	var user models.CreateUser

	err := c.Bind(&user.User)
	if err != nil {
		return h.Response(c, http.StatusBadRequest, h.ResponseModel{
			Data:    nil,
			Message: err.Error(),
			Status:  false,
		})
	}

	user.User, err = u.userS.LoginService(*user.User)
	if err != nil {
		return h.Response(c, http.StatusBadRequest, h.ResponseModel{
			Data:    nil,
			Message: err.Error(),
			Status:  false,
		})
	}

	token, err := u.jwt.CreateJWTToken(user.User.ID, user.User.Nama)
	if err != nil {
		return h.Response(c, http.StatusBadRequest, h.ResponseModel{
			Data:    nil,
			Message: err.Error(),
			Status:  false,
		})
	}

	user.User.Password = "-"

	user.Token = token
	return h.Response(c, http.StatusOK, h.ResponseModel{
		Data:    user,
		Message: "Login success",
		Status:  true,
	})
}