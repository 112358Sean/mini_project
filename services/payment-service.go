package services

import (
	"mini_project/models"
	"mini_project/repositories"
)

type PaymentService interface {
	GetPaymentsService() ([]*models.Payment, error)
	GetPaymentService(id string) (*models.Payment, error)
	CreateService(Payment models.Payment) (*models.Payment, error)
	UpdateService(id string, PaymentBody models.Payment) (*models.Payment, error)
	DeleteService(id string) error
	CheckOut(Payment models.Payment, id_user string) (*models.Payment, error)
}

type paymentService struct {
	PaymentR repositories.PaymentRepository
}

func NewPaymentService(PaymentR repositories.PaymentRepository) PaymentService {
	return &paymentService{
		PaymentR: PaymentR,
	}
}

func (p *paymentService) GetPaymentsService() ([]*models.Payment, error) {
	Payments, err := p.PaymentR.GetPaymentsRepository()
	if err != nil {
		return nil, err
	}

	return Payments, nil
}

func (p *paymentService) GetPaymentService(id string) (*models.Payment, error) {
	Payment, err := p.PaymentR.GetPaymentRepository(id)
	if err != nil {
		return nil, err
	}

	return Payment, nil
}

func (p *paymentService) CreateService(Payment models.Payment) (*models.Payment, error) {
	PaymentR, err := p.PaymentR.CreateRepository(Payment)
	if err != nil {
		return nil, err
	}

	return PaymentR, nil
}

func (p *paymentService) UpdateService(id_user string, PaymentBody models.Payment) (*models.Payment, error) {
	Payment, err := p.PaymentR.UpdateRepository(id_user, PaymentBody)
	if err != nil {
		return Payment, err
	}

	return Payment, nil
}

func (p *paymentService) DeleteService(id string) error {
	err := p.PaymentR.DeleteRepository(id)
	if err != nil {
		return err
	}

	return nil
}

func (p *paymentService) CheckOut(Payment models.Payment, id_user string) (*models.Payment, error) {
	PaymentR, err := p.PaymentR.CheckOut(id_user)
	if err != nil {
		return nil, err
	}

	return PaymentR, nil
}
