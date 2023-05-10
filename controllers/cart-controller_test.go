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
	cartRMock = &repositories.IcartRepositoryMock{Mock: mock.Mock{}}
	cartSMock = services.NewCartService(cartRMock)
	cartCTest = NewCartController(cartSMock)
)

func TestGetCartsController_Success(t *testing.T) {
	carts := []*models.Cart{
		{
			ID_User: "1",
			ID_Buku: "1",
			Jumlah: 1,
		},
		{
			ID_User: "1",
			ID_Buku: "1",
			Jumlah: 1,
		},
	}

	cartRMock.Mock.On("GetCartsRepository").Return(carts, nil)

	rec := httptest.NewRecorder()

	req := httptest.NewRequest(http.MethodGet, "/", nil)

	e := echo.New()

	c := e.NewContext(req, rec)

	err := cartCTest.GetCartsController(c)
	assert.Nil(t, err)
}

func TestGetCartsController_Failure(t *testing.T) {
	cartRMock = &repositories.IcartRepositoryMock{Mock: mock.Mock{}}
	cartSMock = services.NewCartService(cartRMock)
	cartCTest = NewCartController(cartSMock)
	cartRMock.Mock.On("GetCartsRepository").Return(nil, errors.New("get all carts failed"))

	rec := httptest.NewRecorder()

	req := httptest.NewRequest(http.MethodGet, "/", nil)

	e := echo.New()

	c := e.NewContext(req, rec)

	err := cartCTest.GetCartsController(c)
	assert.Nil(t, err)
}

func TestGetCartController_Success(t *testing.T) {
	cart := models.Cart{
		Model: gorm.Model{
			ID: 2,
		},
		ID_User: "1",
		ID_Buku: "1",
		Jumlah: 1,
	}

	cartRMock.Mock.On("GetCartRepository", "2").Return(cart, nil)

	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/carts/2", nil)

	e := echo.New()

	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("2")

	err := cartCTest.GetCartController(c)
	assert.Nil(t, err)
}

func TestGetCartController_Failure1(t *testing.T) {
	cartRMock.Mock.On("GetCartRepository", "qwe").Return(nil, errors.New("get cart failed"))

	rec := httptest.NewRecorder()

	req := httptest.NewRequest(http.MethodGet, "/carts/qwe", nil)

	e := echo.New()

	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("qwe")

	err := cartCTest.GetCartController(c)
	assert.Nil(t, err)
}

func TestGetCartController_Failure2(t *testing.T) {
	cartRMock.Mock.On("GetCartRepository", "3").Return(nil, fmt.Errorf("cart not found"))

	rec := httptest.NewRecorder()

	req := httptest.NewRequest(http.MethodGet, "/carts/3", nil)

	e := echo.New()

	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("3")

	err := cartCTest.GetCartController(c)
	assert.Nil(t, err)
}

func TestCreateCartController_Success(t *testing.T) {
	cart := models.Cart{
		ID_User: "1",
		ID_Buku: "1",
		Jumlah: 1,
	}

	cartRMock.Mock.On("CreateRepository", cart).Return(cart, nil)

	rec := httptest.NewRecorder()

	cartByte, err := json.Marshal(cart)
	if err != nil {
		t.Error(err)
	}

	reqBody := strings.NewReader(string(cartByte))

	req := httptest.NewRequest(http.MethodPost, "/carts", reqBody)
	req.Header.Add("Content-type", "application/json")
	e := echo.New()
	c := e.NewContext(req, rec)

	err = cartCTest.CreateController(c)
	assert.Nil(t, err)
}

func TestCreateCartController_Failure1(t *testing.T) {
	cart := models.Cart{}

	cartRMock.Mock.On("CreateRepository", cart).Return(nil, errors.New("something wrong"))

	rec := httptest.NewRecorder()

	cartByte, err := json.Marshal(cart)
	if err != nil {
		t.Error(err)
	}

	reqBody := strings.NewReader(string(cartByte))

	req := httptest.NewRequest(http.MethodPost, "/carts", reqBody)
	req.Header.Add("Content-type", "application/json")
	e := echo.New()
	c := e.NewContext(req, rec)

	err = cartCTest.CreateController(c)
	assert.Nil(t, err)
}

func TestCreateCartController_Failure2(t *testing.T) {
	cart := models.Cart{}

	cartRMock.Mock.On("CreateRepository", cart).Return(nil, errors.New("something wrong"))

	rec := httptest.NewRecorder()

	reqBody := strings.NewReader(string([]byte("qwe")))

	req := httptest.NewRequest(http.MethodPost, "/carts", reqBody)
	req.Header.Add("Content-type", "application/json")
	e := echo.New()
	c := e.NewContext(req, rec)

	err := cartCTest.CreateController(c)
	assert.Nil(t, err)
}

func TestUpdateCartController_Success(t *testing.T) {
	cart := models.Cart{
		Model: gorm.Model{
			ID: 1,
		},
		ID_User: "1",
		ID_Buku: "1",
		Jumlah: 1,
	}

	cartRMock.Mock.On("UpdateRepository", "1", cart).Return(cart, nil)

	rec := httptest.NewRecorder()

	cartByte, err := json.Marshal(cart)
	if err != nil {
		t.Error(err)
	}

	reqBody := strings.NewReader(string(cartByte))

	req := httptest.NewRequest(http.MethodPut, "/carts/1", reqBody)
	req.Header.Add("Content-type", "application/json")
	e := echo.New()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("1")

	err = cartCTest.UpdateController(c)
	assert.Nil(t, err)
}

func TestUpdateCartController_Failure1(t *testing.T) {
	cart := models.Cart{
		Model: gorm.Model{
			ID: 1,
		},
		ID_User: "1",
		ID_Buku: "1",
		Jumlah: 1,
	}

	cartRMock.Mock.On("UpdateRepository", "1", cart).Return(cart, nil)

	rec := httptest.NewRecorder()

	cartByte, err := json.Marshal(cart)
	if err != nil {
		t.Error(err)
	}

	reqBody := strings.NewReader(string(cartByte))

	req := httptest.NewRequest(http.MethodPut, "/carts/qwe", reqBody)
	req.Header.Add("Content-type", "application/json")
	e := echo.New()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("qwe")

	err = cartCTest.UpdateController(c)
	assert.Nil(t, err)
}

func TestUpdateCartController_Failure2(t *testing.T) {
	cart := models.Cart{}

	cartRMock.Mock.On("UpdateRepository", "1", cart).Return(cart, nil)

	rec := httptest.NewRecorder()

	_, err := json.Marshal(cart)
	if err != nil {
		t.Error(err)
	}

	reqBody := strings.NewReader(string([]byte("qwe")))

	req := httptest.NewRequest(http.MethodPut, "/carts/qwe", reqBody)
	req.Header.Add("Content-type", "application/json")
	e := echo.New()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("1")

	err = cartCTest.UpdateController(c)
	assert.Nil(t, err)
}

func TestUpdateCartController_Failure3(t *testing.T) {
	cartRMock = &repositories.IcartRepositoryMock{Mock: mock.Mock{}}
	cartSMock = services.NewCartService(cartRMock)
	cartCTest = NewCartController(cartSMock)
	cart := models.Cart{
		Model: gorm.Model{
			ID: 1,
		},
		ID_User: "1",
		ID_Buku: "1",
		Jumlah: 1,
	}

	cartRMock.Mock.On("UpdateRepository", "1", cart).Return(nil, errors.New("something wrong"))

	rec := httptest.NewRecorder()

	cartByte, err := json.Marshal(cart)
	if err != nil {
		t.Error(err)
	}

	reqBody := strings.NewReader(string(cartByte))

	req := httptest.NewRequest(http.MethodPut, "/carts/1", reqBody)
	req.Header.Add("Content-type", "application/json")
	e := echo.New()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("1")

	err = cartCTest.UpdateController(c)
	assert.Nil(t, err)
}

func TestDeleteCartController_Success(t *testing.T) {
	cartRMock.Mock.On("DeleteRepository", "2").Return(nil)
	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodDelete, "/carts/2", nil)

	e := echo.New()

	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("2")

	err := cartCTest.DeleteController(c)

	assert.Nil(t, err)
}

func TestDeleteCartController_Failure1(t *testing.T) {
	cartRMock = &repositories.IcartRepositoryMock{Mock: mock.Mock{}}
	cartSMock = services.NewCartService(cartRMock)
	cartCTest = NewCartController(cartSMock)
	cartRMock.Mock.On("DeleteRepository", "2").Return(fmt.Errorf("cart not found"))
	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodDelete, "/carts/2", nil)

	e := echo.New()

	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("2")

	err := cartCTest.DeleteController(c)

	assert.Nil(t, err)
}

func TestDeleteCartController_Failure2(t *testing.T) {
	cartRMock.Mock.On("DeleteRepository", "2").Return(fmt.Errorf("cart not found"))
	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/carts/qwe", nil)

	e := echo.New()

	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("qwe")

	err := cartCTest.DeleteController(c)

	assert.Nil(t, err)
}

func TestGetCartByUserController_Failure1(t *testing.T) {
	cartRMock = &repositories.IcartRepositoryMock{Mock: mock.Mock{}}
	cartSMock = services.NewCartService(cartRMock)
	cartCTest = NewCartController(cartSMock)
	cartRMock.Mock.On("GetCartByUserController", "2").Return(fmt.Errorf("cart not found"))
	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/users/2/carts", nil)

	e := echo.New()

	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("3")

	err := cartCTest.GetCartByUserController(c)

	assert.Nil(t, err)
}

func TestGetCartByUserController_Failure2(t *testing.T) {
	cartRMock.Mock.On("GetCartByUserController", "3").Return(fmt.Errorf("cart not found"))
	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/carts/qwe", nil)

	e := echo.New()

	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("qwe")

	err := cartCTest.GetCartByUserController(c)

	assert.Nil(t, err)
}
