package repositories

import (
	"mini_project/models"
	"fmt"

	"github.com/stretchr/testify/mock"
)

type PaymentRepositoryMock interface {
	GetPaymentsRepository() ([]*models.Payment, error)
	GetPaymentRepository(id string) (*models.Payment, error)
	CreateRepository(Payment models.Payment) (*models.Payment, error)
	UpdateRepository(id string, PaymentBody models.Payment) (*models.Payment, error)
	DeleteRepository(id string) error
	CheckOut(id_user string) (*models.Payment, error)
}

type IpaymentRepositoryMock struct {
	Mock mock.Mock
}

func NewPaymentRepositoryMock(mock mock.Mock) PaymentRepositoryMock {
	return &IpaymentRepositoryMock{
		Mock: mock,
	}
}

func (b *IpaymentRepositoryMock) GetPaymentsRepository() ([]*models.Payment, error) {
	args := b.Mock.Called()
	if args.Get(0) == nil {
		return nil, args.Get(1).(error)
	}

	payments := args.Get(0).([]*models.Payment)

	return payments, nil
}

func (b *IpaymentRepositoryMock) GetPaymentRepository(id string) (*models.Payment, error) {
	args := b.Mock.Called(id)
	if args.Get(0) == nil {
		return nil, args.Get(1).(error)
	}

	payment := args.Get(0).(models.Payment)

	return &payment, nil
}

func (u *IpaymentRepositoryMock) CreateRepository(paymentData models.Payment) (*models.Payment, error) {
	args := u.Mock.Called(paymentData)
	if args.Get(0) == nil {
		return nil, args.Get(1).(error)
	}

	payment := args.Get(0).(models.Payment)

	return &payment, nil
}

func (u *IpaymentRepositoryMock) UpdateRepository(id string, paymentData models.Payment) (*models.Payment, error) {
	args := u.Mock.Called(id, paymentData)
	if args.Get(0) == nil {
		return nil,  args.Get(1).(error)
	}

	payment := args.Get(0).(models.Payment)

	return &payment, nil
}

func (u *IpaymentRepositoryMock) DeleteRepository(id string) error {
	args := u.Mock.Called(id)
	if args.Get(0) != nil {
		return fmt.Errorf("must nil")
	}

	return nil
}

func (b *IpaymentRepositoryMock) CheckOut(id_user string) (*models.Payment, error) {
	args := b.Mock.Called(id_user)
	if args.Get(0) == nil {
		return nil, args.Get(1).(error)
	}

	payment := args.Get(0).(models.Payment)

	return &payment, nil
}