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
	paymentRMock = &repositories.IpaymentRepositoryMock{Mock: mock.Mock{}}
	paymentSMock = services.NewPaymentService(paymentRMock)
	paymentCTest = NewPaymentController(paymentSMock)
)

func TestGetPaymentsController_Success(t *testing.T) {
	payments := []*models.Payment{
		{
			Bukti_Pembayaran: "123456",
			Total_Price: 70000,
			Status: "Belum Terbayar",
		},
		{
			Bukti_Pembayaran: "123456",
			Total_Price: 70000,
			Status: "Belum Terbayar",
		},
	}

	paymentRMock.Mock.On("GetPaymentsRepository").Return(payments, nil)

	rec := httptest.NewRecorder()

	req := httptest.NewRequest(http.MethodGet, "/", nil)

	e := echo.New()

	c := e.NewContext(req, rec)

	err := paymentCTest.GetPaymentsController(c)
	assert.Nil(t, err)
}

func TestGetPaymentsController_Failure(t *testing.T) {
	paymentRMock = &repositories.IpaymentRepositoryMock{Mock: mock.Mock{}}
	paymentSMock = services.NewPaymentService(paymentRMock)
	paymentCTest = NewPaymentController(paymentSMock)
	paymentRMock.Mock.On("GetPaymentsRepository").Return(nil, errors.New("get all payments failed"))

	rec := httptest.NewRecorder()

	req := httptest.NewRequest(http.MethodGet, "/", nil)

	e := echo.New()

	c := e.NewContext(req, rec)

	err := paymentCTest.GetPaymentsController(c)
	assert.Nil(t, err)
}

func TestGetpaymentController_Success(t *testing.T) {
	payment := models.Payment{
		Model: gorm.Model{
			ID: 2,
		},
		Bukti_Pembayaran: "123456",
		Total_Price: 70000,
		Status: "Belum Terbayar",
	}

	paymentRMock.Mock.On("GetPaymentRepository", "2").Return(payment, nil)

	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/payments/2", nil)

	e := echo.New()

	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("2")

	err := paymentCTest.GetPaymentController(c)
	assert.Nil(t, err)
}

func TestGetPaymentController_Failure1(t *testing.T) {
	paymentRMock.Mock.On("GetPaymentRepository", "qwe").Return(nil, errors.New("get payment failed"))

	rec := httptest.NewRecorder()

	req := httptest.NewRequest(http.MethodGet, "/payments/qwe", nil)

	e := echo.New()

	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("qwe")

	err := paymentCTest.GetPaymentController(c)
	assert.Nil(t, err)
}

func TestGetPaymentController_Failure2(t *testing.T) {
	paymentRMock.Mock.On("GetPaymentRepository", "3").Return(nil, fmt.Errorf("payment not found"))

	rec := httptest.NewRecorder()

	req := httptest.NewRequest(http.MethodGet, "/payments/3", nil)

	e := echo.New()

	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("3")

	err := paymentCTest.GetPaymentController(c)
	assert.Nil(t, err)
}

func TestCreatePaymentController_Success(t *testing.T) {
	payment := models.Payment{
		Bukti_Pembayaran: "123456",
		Total_Price: 70000,
		Status: "Belum Terbayar",
	}

	paymentRMock.Mock.On("CreateRepository", payment).Return(payment, nil)

	rec := httptest.NewRecorder()

	paymentByte, err := json.Marshal(payment)
	if err != nil {
		t.Error(err)
	}

	reqBody := strings.NewReader(string(paymentByte))

	req := httptest.NewRequest(http.MethodPost, "/payments", reqBody)
	req.Header.Add("Content-type", "application/json")
	e := echo.New()
	c := e.NewContext(req, rec)

	err = paymentCTest.CreateController(c)
	assert.Nil(t, err)
}

func TestCreatePaymentController_Failure1(t *testing.T) {
	payment := models.Payment{}

	paymentRMock.Mock.On("CreateRepository", payment).Return(nil, errors.New("something wrong"))

	rec := httptest.NewRecorder()

	paymentByte, err := json.Marshal(payment)
	if err != nil {
		t.Error(err)
	}

	reqBody := strings.NewReader(string(paymentByte))

	req := httptest.NewRequest(http.MethodPost, "/payments", reqBody)
	req.Header.Add("Content-type", "application/json")
	e := echo.New()
	c := e.NewContext(req, rec)

	err = paymentCTest.CreateController(c)
	assert.Nil(t, err)
}

func TestCreatePaymentController_Failure2(t *testing.T) {
	payment := models.Payment{}

	paymentRMock.Mock.On("CreateRepository", payment).Return(nil, errors.New("something wrong"))

	rec := httptest.NewRecorder()

	reqBody := strings.NewReader(string([]byte("qwe")))

	req := httptest.NewRequest(http.MethodPost, "/payments", reqBody)
	req.Header.Add("Content-type", "application/json")
	e := echo.New()
	c := e.NewContext(req, rec)

	err := paymentCTest.CreateController(c)
	assert.Nil(t, err)
}

func TestUpdatePaymentController_Success(t *testing.T) {
	payment := models.Payment{
		Model: gorm.Model{
			ID: 1,
		},
		Bukti_Pembayaran: "123456",
		Total_Price: 70000,
		Status: "Belum Terbayar",
	}

	paymentRMock.Mock.On("UpdateRepository", "1", payment).Return(payment, nil)

	rec := httptest.NewRecorder()

	paymentByte, err := json.Marshal(payment)
	if err != nil {
		t.Error(err)
	}

	reqBody := strings.NewReader(string(paymentByte))

	req := httptest.NewRequest(http.MethodPut, "/payments/1", reqBody)
	req.Header.Add("Content-type", "application/json")
	e := echo.New()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("1")

	err = paymentCTest.UpdateController(c)
	assert.Nil(t, err)
}

func TestUpdatePaymentController_Failure1(t *testing.T) {
	payment := models.Payment{
		Model: gorm.Model{
			ID: 1,
		},
		Bukti_Pembayaran: "123456",
		Total_Price: 70000,
		Status: "Belum Terbayar",
	}

	paymentRMock.Mock.On("UpdateRepository", "1", payment).Return(payment, nil)

	rec := httptest.NewRecorder()

	paymentByte, err := json.Marshal(payment)
	if err != nil {
		t.Error(err)
	}

	reqBody := strings.NewReader(string(paymentByte))

	req := httptest.NewRequest(http.MethodPut, "/payments/qwe", reqBody)
	req.Header.Add("Content-type", "application/json")
	e := echo.New()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("qwe")

	err = paymentCTest.UpdateController(c)
	assert.Nil(t, err)
}

func TestUpdatePaymentController_Failure2(t *testing.T) {
	payment := models.Payment{}

	paymentRMock.Mock.On("UpdateRepository", "1", payment).Return(payment, nil)

	rec := httptest.NewRecorder()

	_, err := json.Marshal(payment)
	if err != nil {
		t.Error(err)
	}

	reqBody := strings.NewReader(string([]byte("qwe")))

	req := httptest.NewRequest(http.MethodPut, "/payments/qwe", reqBody)
	req.Header.Add("Content-type", "application/json")
	e := echo.New()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("1")

	err = paymentCTest.UpdateController(c)
	assert.Nil(t, err)
}

func TestUpdatePaymentController_Failure3(t *testing.T) {
	paymentRMock = &repositories.IpaymentRepositoryMock{Mock: mock.Mock{}}
	paymentSMock = services.NewPaymentService(paymentRMock)
	paymentCTest = NewPaymentController(paymentSMock)
	payment := models.Payment{
		Model: gorm.Model{
			ID: 1,
		},
		Bukti_Pembayaran: "123456",
		Total_Price: 70000,
		Status: "Belum Terbayar",
	}

	paymentRMock.Mock.On("UpdateRepository", "1", payment).Return(nil, errors.New("something wrong"))

	rec := httptest.NewRecorder()

	paymentByte, err := json.Marshal(payment)
	if err != nil {
		t.Error(err)
	}

	reqBody := strings.NewReader(string(paymentByte))

	req := httptest.NewRequest(http.MethodPut, "/payments/1", reqBody)
	req.Header.Add("Content-type", "application/json")
	e := echo.New()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("1")

	err = paymentCTest.UpdateController(c)
	assert.Nil(t, err)
}

func TestDeletePaymentController_Success(t *testing.T) {
	paymentRMock.Mock.On("DeleteRepository", "2").Return(nil)
	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodDelete, "/payments/2", nil)

	e := echo.New()

	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("2")

	err := paymentCTest.DeleteController(c)

	assert.Nil(t, err)
}

func TestDeletePaymentController_Failure1(t *testing.T) {
	paymentRMock = &repositories.IpaymentRepositoryMock{Mock: mock.Mock{}}
	paymentSMock = services.NewPaymentService(paymentRMock)
	paymentCTest = NewPaymentController(paymentSMock)
	paymentRMock.Mock.On("DeleteRepository", "2").Return(fmt.Errorf("payment not found"))
	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodDelete, "/payments/2", nil)

	e := echo.New()

	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("2")

	err := paymentCTest.DeleteController(c)

	assert.Nil(t, err)
}

func TestDeletePaymentController_Failure2(t *testing.T) {
	paymentRMock.Mock.On("DeleteRepository", "2").Return(fmt.Errorf("payment not found"))
	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/payments/qwe", nil)

	e := echo.New()

	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("qwe")

	err := paymentCTest.DeleteController(c)

	assert.Nil(t, err)
}

func TestCheckOut_Failure(t *testing.T) {
	payment := models.Payment{}

	paymentRMock.Mock.On("CheckOut", payment).Return(nil, errors.New("something wrong"))

	rec := httptest.NewRecorder()

	reqBody := strings.NewReader(string([]byte("qwe")))

	req := httptest.NewRequest(http.MethodPost, "/users/3/checkout", reqBody)
	req.Header.Add("Content-type", "application/json")
	e := echo.New()
	c := e.NewContext(req, rec)

	err := paymentCTest.CheckOut(c)
	assert.Nil(t, err)
}
