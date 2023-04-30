package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"

	h "mini_project/helpers"
	"mini_project/models"
	"mini_project/services"
)

type TransactionController interface {
	GetTransactionsController(c echo.Context) error
	GetTransactionController(c echo.Context) error
	CreateController(c echo.Context) error
	UpdateController(c echo.Context) error
	DeleteController(c echo.Context) error
}

type transactionController struct {
	TransactionS services.TransactionService
}

func NewTransactionController(TransactionS services.TransactionService) TransactionController {
	return &transactionController{
		TransactionS: TransactionS,
	}
}

func (t *transactionController) GetTransactionsController(c echo.Context) error {
	Transactions, err := t.TransactionS.GetTransactionsService()
	if err != nil {
		return h.Response(c, http.StatusBadRequest, h.ResponseModel{
			Data:    nil,
			Message: err.Error(),
			Status:  false,
		})
	}

	return h.Response(c, http.StatusOK, h.ResponseModel{
		Data:    Transactions,
		Message: "Get all Transactions success",
		Status:  true,
	})
}

func (t *transactionController) GetTransactionController(c echo.Context) error {
	id := c.Param("id")

	err := h.IsNumber(id)
	if err != nil {
		return h.Response(c, http.StatusBadRequest, h.ResponseModel{
			Data:    nil,
			Message: err.Error(),
			Status:  false,
		})
	}

	var Transaction *models.Transaction

	Transaction, err = t.TransactionS.GetTransactionService(id)
	if err != nil {
		return h.Response(c, http.StatusNotFound, h.ResponseModel{
			Data:    nil,
			Message: err.Error(),
			Status:  false,
		})
	}

	return h.Response(c, http.StatusOK, h.ResponseModel{
		Data:    Transaction,
		Message: "Get Transaction success",
		Status:  true,
	})
}

func (t *transactionController) CreateController(c echo.Context) error {
	var Transaction *models.Transaction

	err := c.Bind(&Transaction)
	if err != nil {
		return h.Response(c, http.StatusBadRequest, h.ResponseModel{
			Data:    nil,
			Message: err.Error(),
			Status:  false,
		})
	}

	Transaction, err = t.TransactionS.CreateService(*Transaction)
	if err != nil {
		return h.Response(c, http.StatusBadRequest, h.ResponseModel{
			Data:    nil,
			Message: err.Error(),
			Status:  false,
		})
	}

	return h.Response(c, http.StatusOK, h.ResponseModel{
		Data:    Transaction,
		Message: "Create Transaction success",
		Status:  true,
	})
}

func (t *transactionController) UpdateController(c echo.Context) error {
	id := c.Param("id")

	err := h.IsNumber(id)
	if err != nil {
		return h.Response(c, http.StatusBadRequest, h.ResponseModel{
			Data:    nil,
			Message: err.Error(),
			Status:  false,
		})
	}

	var Transaction *models.Transaction

	err = c.Bind(&Transaction)
	if err != nil {
		return h.Response(c, http.StatusBadRequest, h.ResponseModel{
			Data:    nil,
			Message: err.Error(),
			Status:  false,
		})
	}

	Transaction, err = t.TransactionS.UpdateService(id, *Transaction)
	if err != nil {
		return h.Response(c, http.StatusBadRequest, h.ResponseModel{
			Data:    nil,
			Message: err.Error(),
			Status:  false,
		})
	}

	return h.Response(c, http.StatusOK, h.ResponseModel{
		Data:    Transaction,
		Message: "Update Transaction success",
		Status:  true,
	})
}

func (t *transactionController) DeleteController(c echo.Context) error {
	id := c.Param("id")

	err := h.IsNumber(id)
	if err != nil {
		return h.Response(c, http.StatusBadRequest, h.ResponseModel{
			Data:    nil,
			Message: err.Error(),
			Status:  false,
		})
	}

	err = t.TransactionS.DeleteService(id)
	if err != nil {
		return h.Response(c, http.StatusBadRequest, h.ResponseModel{
			Data:    nil,
			Message: err.Error(),
			Status:  false,
		})
	}

	return h.Response(c, http.StatusOK, h.ResponseModel{
		Data:    nil,
		Message: "Delete Transaction success",
		Status:  true,
	})
}
