package services

import (
	"mini_project/models"
	"mini_project/repositories"
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	paymentRMock = &repositories.IpaymentRepositoryMock{Mock: mock.Mock{}}
	paymentSMock = NewPaymentService(paymentRMock)
)

func TestGetpaymentsService_Success(t *testing.T) {
	paymentsMP := []*models.Payment{
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

	paymentsM := []models.Payment{
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

	paymentRMock.Mock.On("GetPaymentsRepository").Return(paymentsMP)
	payments, err := paymentSMock.GetPaymentsService()

	assert.Nil(t, err)
	assert.NotNil(t, payments)

	assert.Equal(t, paymentsM[0].Bukti_Pembayaran, payments[0].Bukti_Pembayaran)
	assert.Equal(t, paymentsM[0].Total_Price, payments[0].Total_Price)
	assert.Equal(t, paymentsM[0].Status, payments[0].Status)
}

func TestGetPaymentsService_Failure(t *testing.T) {
	paymentRMock = &repositories.IpaymentRepositoryMock{Mock: mock.Mock{}}
	paymentSMock = NewPaymentService(paymentRMock)
	paymentRMock.Mock.On("GetPaymentsRepository").Return(nil, errors.New("get all payments failed"))
	payments, err := paymentSMock.GetPaymentsService()

	assert.Nil(t, payments)
	assert.NotNil(t, err)
}

func TestGetPaymentService_Success(t *testing.T) {
	payment := models.Payment{
		Bukti_Pembayaran: "123456",
		Total_Price: 70000,
		Status: "Belum Terbayar",
	}

	paymentRMock.Mock.On("GetPaymentRepository", "1").Return(payment, nil)
	payments, err := paymentSMock.GetPaymentService("1")

	assert.Nil(t, err)
	assert.NotNil(t, payments)

	assert.Equal(t, payment.Bukti_Pembayaran, payments.Bukti_Pembayaran)
	assert.Equal(t, payment.Total_Price, payments.Total_Price)
	assert.Equal(t, payment.Status, payments.Status)
}

func TestGetPaymentService_Failure(t *testing.T) {
	paymentRMock.Mock.On("GetPaymentRepository", "3").Return(nil, fmt.Errorf("payment not found"))
	payment, err := paymentSMock.GetPaymentService("3")

	assert.NotNil(t, err)
	assert.Nil(t, payment)
}

func TestUpdatePaymentService_Success(t *testing.T) {
	payment := models.Payment{
		Bukti_Pembayaran: "123456",
		Total_Price: 70000,
		Status: "Belum Terbayar",
	}

	paymentRMock.Mock.On("UpdateRepository", "1", payment).Return(payment, nil)
	payments, err := paymentSMock.UpdateService("1", payment)

	assert.Nil(t, err)
	assert.NotNil(t, payments)

	assert.Equal(t, payment.ID, payments.ID)
	assert.Equal(t, payment.Bukti_Pembayaran, payments.Bukti_Pembayaran)
	assert.Equal(t, payment.Total_Price, payments.Total_Price)
	assert.Equal(t, payment.Status, payments.Status)
}

func TestUpdatePaymentService_Failure(t *testing.T) {
	payment := models.Payment{
		Bukti_Pembayaran: "123456",
		Total_Price: 70000,
		Status: "Belum Terbayar",
	}

	paymentRMock.Mock.On("UpdateRepository", "2", payment).Return(nil, fmt.Errorf("payment not found"))
	payments, err := paymentSMock.UpdateService("2", payment)

	assert.Nil(t, payments)
	assert.NotNil(t, err)
}

func TestDeletepaymentService_Success(t *testing.T) {
	paymentRMock.Mock.On("DeleteRepository", "1").Return(nil)
	err := paymentSMock.DeleteService("1")

	assert.Nil(t, err)
}

func TestDeletepaymentService_Failure(t *testing.T) {
	paymentRMock.Mock.On("DeleteRepository", "2").Return(fmt.Errorf("payment not found"))
	err := paymentSMock.DeleteService("2")

	assert.NotNil(t, err)
}
