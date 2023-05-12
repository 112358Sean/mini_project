package controllers

import (
	"mini_project/models"
	"mini_project/repositories"
	"mini_project/services"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

var (
	transactionRMock = &repositories.ItransactionRepositoryMock{Mock: mock.Mock{}}
	transactionSMock = services.NewTransactionService(transactionRMock)
	transactionCTest = NewTransactionController(transactionSMock)
)

func TestGetTransactionsController_Success(t *testing.T) {
	transactions := []*models.Transaction{
		{
			ID_Transaksi: "1",
			ID_Keranjang: "1",
			Status: "pending",
		},
		{
			ID_Transaksi: "1",
			ID_Keranjang: "1",
			Status: "pending",
		},
	}

	transactionRMock.Mock.On("GetTransactionsRepository").Return(transactions, nil)

	rec := httptest.NewRecorder()

	req := httptest.NewRequest(http.MethodGet, "/", nil)

	e := echo.New()

	c := e.NewContext(req, rec)

	err := transactionCTest.GetTransactionsController(c)
	assert.Nil(t, err)
}

func TestGettransactionsController_Failure(t *testing.T) {
	transactionRMock = &repositories.ItransactionRepositoryMock{Mock: mock.Mock{}}
	transactionSMock = services.NewTransactionService(transactionRMock)
	transactionCTest = NewTransactionController(transactionSMock)
	transactionRMock.Mock.On("GetTransactionsRepository").Return(nil, errors.New("get all Transactions failed"))

	rec := httptest.NewRecorder()

	req := httptest.NewRequest(http.MethodGet, "/", nil)

	e := echo.New()

	c := e.NewContext(req, rec)

	err := transactionCTest.GetTransactionsController(c)
	assert.Nil(t, err)
}

func TestGetTransactionController_Success(t *testing.T) {
	transaction := models.Transaction{
		Model: gorm.Model{
			ID: 2,
		},
		ID_Transaksi: "1",
		ID_Keranjang: "1",
		Status: "pending",
	}

	transactionRMock.Mock.On("GetTransactionRepository", "2").Return(transaction, nil)

	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/transactions/2", nil)

	e := echo.New()

	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("2")

	err := transactionCTest.GetTransactionController(c)
	assert.Nil(t, err)
}

func TestGetTransactionController_Failure1(t *testing.T) {
	transactionRMock.Mock.On("GetTransactionRepository", "qwe").Return(nil, errors.New("get transaction failed"))

	rec := httptest.NewRecorder()

	req := httptest.NewRequest(http.MethodGet, "/transactions/qwe", nil)

	e := echo.New()

	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("qwe")

	err := transactionCTest.GetTransactionController(c)
	assert.Nil(t, err)
}

func TestGetTransactionController_Failure2(t *testing.T) {
	transactionRMock.Mock.On("GetTransactionRepository", "3").Return(nil, fmt.Errorf("transaction not found"))

	rec := httptest.NewRecorder()

	req := httptest.NewRequest(http.MethodGet, "/transactions/3", nil)

	e := echo.New()

	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("3")

	err := transactionCTest.GetTransactionController(c)
	assert.Nil(t, err)
}

func TestCreateTransactionController_Success(t *testing.T) {
	transaction := models.Transaction{
		ID_Keranjang: "1",
		Status: "pending",
	}

	transactionRMock.Mock.On("CreateRepository", transaction).Return(transaction, nil)

	rec := httptest.NewRecorder()

	transactionByte, err := json.Marshal(transaction)
	if err != nil {
		t.Error(err)
	}

	reqBody := strings.NewReader(string(transactionByte))

	req := httptest.NewRequest(http.MethodPost, "/transactions", reqBody)
	req.Header.Add("Content-type", "application/json")
	e := echo.New()
	c := e.NewContext(req, rec)

	err = transactionCTest.CreateController(c)
	assert.Nil(t, err)
}

func TestCreateTransactionController_Failure1(t *testing.T) {
	transaction := models.Transaction{}

	transactionRMock.Mock.On("CreateRepository", transaction).Return(nil, errors.New("something wrong"))

	rec := httptest.NewRecorder()

	transactionByte, err := json.Marshal(transaction)
	if err != nil {
		t.Error(err)
	}

	reqBody := strings.NewReader(string(transactionByte))

	req := httptest.NewRequest(http.MethodPost, "/transactions", reqBody)
	req.Header.Add("Content-type", "application/json")
	e := echo.New()
	c := e.NewContext(req, rec)

	err = transactionCTest.CreateController(c)
	assert.Nil(t, err)
}

func TestCreateTransactionController_Failure2(t *testing.T) {
	transaction := models.Transaction{}

	transactionRMock.Mock.On("CreateRepository", transaction).Return(nil, errors.New("something wrong"))

	rec := httptest.NewRecorder()

	reqBody := strings.NewReader(string([]byte("qwe")))

	req := httptest.NewRequest(http.MethodPost, "/transactions", reqBody)
	req.Header.Add("Content-type", "application/json")
	e := echo.New()
	c := e.NewContext(req, rec)

	err := transactionCTest.CreateController(c)
	assert.Nil(t, err)
}

func TestUpdateTransactionController_Success(t *testing.T) {
	transaction := models.Transaction{
		Model: gorm.Model{
			ID: 1,
		},
		ID_Keranjang: "1",
		Status: "pending",
	}

	transactionRMock.Mock.On("UpdateRepository", "1", transaction).Return(transaction, nil)

	rec := httptest.NewRecorder()

	transactionByte, err := json.Marshal(transaction)
	if err != nil {
		t.Error(err)
	}

	reqBody := strings.NewReader(string(transactionByte))

	req := httptest.NewRequest(http.MethodPut, "/transactions/1", reqBody)
	req.Header.Add("Content-type", "application/json")
	e := echo.New()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("1")

	err = transactionCTest.UpdateController(c)
	assert.Nil(t, err)
}

func TestUpdateTransactionController_Failure1(t *testing.T) {
	transaction := models.Transaction{
		Model: gorm.Model{
			ID: 1,
		},
		ID_Transaksi: "1",
		ID_Keranjang: "1",
		Status: "pending",
	}

	transactionRMock.Mock.On("UpdateRepository", "1", transaction).Return(transaction, nil)

	rec := httptest.NewRecorder()

	transactionByte, err := json.Marshal(transaction)
	if err != nil {
		t.Error(err)
	}

	reqBody := strings.NewReader(string(transactionByte))

	req := httptest.NewRequest(http.MethodPut, "/transactions/qwe", reqBody)
	req.Header.Add("Content-type", "application/json")
	e := echo.New()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("qwe")

	err = transactionCTest.UpdateController(c)
	assert.Nil(t, err)
}

func TestUpdateTransactionController_Failure2(t *testing.T) {
	transaction := models.Transaction{}

	transactionRMock.Mock.On("UpdateRepository", "1", transaction).Return(transaction, nil)

	rec := httptest.NewRecorder()

	_, err := json.Marshal(transaction)
	if err != nil {
		t.Error(err)
	}

	reqBody := strings.NewReader(string([]byte("qwe")))

	req := httptest.NewRequest(http.MethodPut, "/transactions/qwe", reqBody)
	req.Header.Add("Content-type", "application/json")
	e := echo.New()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("1")

	err = transactionCTest.UpdateController(c)
	assert.Nil(t, err)
}

func TestUpdateTransactionController_Failure3(t *testing.T) {
	transactionRMock = &repositories.ItransactionRepositoryMock{Mock: mock.Mock{}}
	transactionSMock = services.NewTransactionService(transactionRMock)
	transactionCTest = NewTransactionController(transactionSMock)
	transaction := models.Transaction{
		Model: gorm.Model{
			ID: 1,
		},
		ID_Transaksi: "1",
		ID_Keranjang: "1",
		Status: "pending",
	}

	transactionRMock.Mock.On("UpdateRepository", "1", transaction).Return(nil, errors.New("something wrong"))

	rec := httptest.NewRecorder()

	transactionByte, err := json.Marshal(transaction)
	if err != nil {
		t.Error(err)
	}

	reqBody := strings.NewReader(string(transactionByte))

	req := httptest.NewRequest(http.MethodPut, "/transactions/1", reqBody)
	req.Header.Add("Content-type", "application/json")
	e := echo.New()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("1")

	err = transactionCTest.UpdateController(c)
	assert.Nil(t, err)
}

func TestDeleteTransactionController_Success(t *testing.T) {
	transactionRMock.Mock.On("DeleteRepository", "2").Return(nil)
	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodDelete, "/transactions/2", nil)

	e := echo.New()

	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("2")

	err := transactionCTest.DeleteController(c)

	assert.Nil(t, err)
}

func TestDeleteTransactionController_Failure1(t *testing.T) {
	transactionRMock = &repositories.ItransactionRepositoryMock{Mock: mock.Mock{}}
	transactionSMock = services.NewTransactionService(transactionRMock)
	transactionCTest = NewTransactionController(transactionSMock)
	transactionRMock.Mock.On("DeleteRepository", "2").Return(fmt.Errorf("transaction not found"))
	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodDelete, "/transactions/2", nil)

	e := echo.New()

	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("2")

	err := transactionCTest.DeleteController(c)

	assert.Nil(t, err)
}

func TestDeleteTransactionController_Failure2(t *testing.T) {
	transactionRMock.Mock.On("DeleteRepository", "2").Return(fmt.Errorf("transaction not found"))
	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/transactions/qwe", nil)

	e := echo.New()

	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("qwe")

	err := transactionCTest.DeleteController(c)

	assert.Nil(t, err)
}
