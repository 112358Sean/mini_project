package controllers

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"

	h "mini_project/helpers"
	"mini_project/models"
	"mini_project/services"
)

type PaymentController interface {
	GetPaymentsController(c echo.Context) error
	GetPaymentController(c echo.Context) error
	CreateController(c echo.Context) error
	UpdateController(c echo.Context) error
	DeleteController(c echo.Context) error
	CheckOut(c echo.Context) error
}

type paymentController struct {
	PaymentS services.PaymentService
}

func NewPaymentController(PaymentS services.PaymentService) PaymentController {
	return &paymentController{
		PaymentS: PaymentS,
	}
}

func (p *paymentController) GetPaymentsController(c echo.Context) error {
	Payments, err := p.PaymentS.GetPaymentsService()
	if err != nil {
		return h.Response(c, http.StatusBadRequest, h.ResponseModel{
			Data:    nil,
			Message: err.Error(),
			Status:  false,
		})
	}

	return h.Response(c, http.StatusOK, h.ResponseModel{
		Data:    Payments,
		Message: "Get all Payments success",
		Status:  true,
	})
}

func (p *paymentController) GetPaymentController(c echo.Context) error {
	id := c.Param("id")

	err := h.IsNumber(id)
	if err != nil {
		return h.Response(c, http.StatusBadRequest, h.ResponseModel{
			Data:    nil,
			Message: err.Error(),
			Status:  false,
		})
	}

	var Payment *models.Payment

	Payment, err = p.PaymentS.GetPaymentService(id)
	if err != nil {
		return h.Response(c, http.StatusNotFound, h.ResponseModel{
			Data:    nil,
			Message: err.Error(),
			Status:  false,
		})
	}

	return h.Response(c, http.StatusOK, h.ResponseModel{
		Data:    Payment,
		Message: "Get Payment success",
		Status:  true,
	})
}

func (p *paymentController) CreateController(c echo.Context) error {
	var Payment *models.Payment

	err := c.Bind(&Payment)
	if err != nil {
		return h.Response(c, http.StatusBadRequest, h.ResponseModel{
			Data:    nil,
			Message: err.Error(),
			Status:  false,
		})
	}

	Payment, err = p.PaymentS.CreateService(*Payment)
	if err != nil {
		return h.Response(c, http.StatusBadRequest, h.ResponseModel{
			Data:    nil,
			Message: err.Error(),
			Status:  false,
		})
	}

	return h.Response(c, http.StatusOK, h.ResponseModel{
		Data:    Payment,
		Message: "Create Payment success",
		Status:  true,
	})
}

func (p *paymentController) UpdateController(c echo.Context) error {
	id := c.Param("id_user")

	err := h.IsNumber(id)
	if err != nil {
		return h.Response(c, http.StatusBadRequest, h.ResponseModel{
			Data:    nil,
			Message: err.Error(),
			Status:  false,
		})
	}

	var Payment *models.Payment

	err = c.Bind(&Payment)
	if err != nil {
		return h.Response(c, http.StatusBadRequest, h.ResponseModel{
			Data:    nil,
			Message: err.Error(),
			Status:  false,
		})
	}

	Payment, err = p.PaymentS.UpdateService(id, *Payment)
	if err != nil {
		return h.Response(c, http.StatusBadRequest, h.ResponseModel{
			Data:    nil,
			Message: err.Error(),
			Status:  false,
		})
	}

	return h.Response(c, http.StatusOK, h.ResponseModel{
		Data: map[string]interface{}{
			"Payment":    Payment,
			"Ongkos Kirim": 10000,
			"Total Pembayaran" : Payment.Total_Price + 10000,
		},
		Message: "Payment success",
		Status:  true,
	})
}

func (p *paymentController) DeleteController(c echo.Context) error {
	id := c.Param("id")

	err := h.IsNumber(id)
	if err != nil {
		return h.Response(c, http.StatusBadRequest, h.ResponseModel{
			Data:    nil,
			Message: err.Error(),
			Status:  false,
		})
	}

	err = p.PaymentS.DeleteService(id)
	if err != nil {
		return h.Response(c, http.StatusBadRequest, h.ResponseModel{
			Data:    nil,
			Message: err.Error(),
			Status:  false,
		})
	}

	return h.Response(c, http.StatusOK, h.ResponseModel{
		Data:    nil,
		Message: "Delete Payment success",
		Status:  true,
	})
}

func (p *paymentController) CheckOut(c echo.Context) error {
	var Payment models.Payment

	err := c.Bind(&Payment)
	if err != nil {
		return h.Response(c, http.StatusBadRequest, h.ResponseModel{
			Data:    nil,
			Message: err.Error(),
			Status:  false,
		})
	}

	PaymentR, err := p.PaymentS.CheckOut(Payment, c.Param("id_user"))

	fmt.Println(err)

	PaymentR.ID_Transaksi = Payment.ID_Transaksi
	PaymentR.Bukti_Pembayaran = "-"
	PaymentR.Status = "Belum Terbayar"

	return h.Response(c, http.StatusOK, h.ResponseModel{
		Data: map[string]interface{}{
			"Payment":    PaymentR,
			"Ongkos Kirim": 10000,
			"Total Pembayaran" : PaymentR.Total_Price + 10000,
		},
		Message: "Create Payment and CheckOut success",
		Status:  true,
	})
}
