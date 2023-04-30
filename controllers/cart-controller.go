package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"

	h "mini_project/helpers"
	"mini_project/models"
	"mini_project/services"
)

type CartController interface {
	GetCartsController(c echo.Context) error
	GetCartController(c echo.Context) error
	CreateController(c echo.Context) error
	UpdateController(c echo.Context) error
	DeleteController(c echo.Context) error
}

type cartController struct {
	CartS services.CartService
}

func NewCartController(CartS services.CartService) CartController {
	return &cartController{
		CartS: CartS,
	}
}

func (a *cartController) GetCartsController(c echo.Context) error {
	Carts, err := a.CartS.GetCartsService()
	if err != nil {
		return h.Response(c, http.StatusBadRequest, h.ResponseModel{
			Data:    nil,
			Message: err.Error(),
			Status:  false,
		})
	}

	return h.Response(c, http.StatusOK, h.ResponseModel{
		Data:    Carts,
		Message: "Get all Carts success",
		Status:  true,
	})
}

func (a *cartController) GetCartController(c echo.Context) error {
	id := c.Param("id")

	err := h.IsNumber(id)
	if err != nil {
		return h.Response(c, http.StatusBadRequest, h.ResponseModel{
			Data:    nil,
			Message: err.Error(),
			Status:  false,
		})
	}

	var Cart *models.Cart

	Cart, err = a.CartS.GetCartService(id)
	if err != nil {
		return h.Response(c, http.StatusNotFound, h.ResponseModel{
			Data:    nil,
			Message: err.Error(),
			Status:  false,
		})
	}

	return h.Response(c, http.StatusOK, h.ResponseModel{
		Data:    Cart,
		Message: "Get Cart success",
		Status:  true,
	})
}

func (a *cartController) CreateController(c echo.Context) error {
	var Cart *models.Cart

	err := c.Bind(&Cart)
	if err != nil {
		return h.Response(c, http.StatusBadRequest, h.ResponseModel{
			Data:    nil,
			Message: err.Error(),
			Status:  false,
		})
	}

	Cart, err = a.CartS.CreateService(*Cart)
	if err != nil {
		return h.Response(c, http.StatusBadRequest, h.ResponseModel{
			Data:    nil,
			Message: err.Error(),
			Status:  false,
		})
	}

	return h.Response(c, http.StatusOK, h.ResponseModel{
		Data:    Cart,
		Message: "Create Cart success",
		Status:  true,
	})
}

func (a *cartController) UpdateController(c echo.Context) error {
	id := c.Param("id")

	err := h.IsNumber(id)
	if err != nil {
		return h.Response(c, http.StatusBadRequest, h.ResponseModel{
			Data:    nil,
			Message: err.Error(),
			Status:  false,
		})
	}

	var Cart *models.Cart

	err = c.Bind(&Cart)
	if err != nil {
		return h.Response(c, http.StatusBadRequest, h.ResponseModel{
			Data:    nil,
			Message: err.Error(),
			Status:  false,
		})
	}

	Cart, err = a.CartS.UpdateService(id, *Cart)
	if err != nil {
		return h.Response(c, http.StatusBadRequest, h.ResponseModel{
			Data:    nil,
			Message: err.Error(),
			Status:  false,
		})
	}

	return h.Response(c, http.StatusOK, h.ResponseModel{
		Data:    Cart,
		Message: "Update Cart success",
		Status:  true,
	})
}

func (a *cartController) DeleteController(c echo.Context) error {
	id := c.Param("id")

	err := h.IsNumber(id)
	if err != nil {
		return h.Response(c, http.StatusBadRequest, h.ResponseModel{
			Data:    nil,
			Message: err.Error(),
			Status:  false,
		})
	}

	err = a.CartS.DeleteService(id)
	if err != nil {
		return h.Response(c, http.StatusBadRequest, h.ResponseModel{
			Data:    nil,
			Message: err.Error(),
			Status:  false,
		})
	}

	return h.Response(c, http.StatusOK, h.ResponseModel{
		Data:    nil,
		Message: "Delete Cart success",
		Status:  true,
	})
}
