package services

import (
	"mini_project/models"
	"mini_project/repositories"
)

type TransactionService interface {
	GetTransactionsService() ([]*models.Transaction, error)
	GetTransactionService(id string) (*models.Transaction, error)
	CreateService(Transaction models.Transaction) (*models.Transaction, error)
	UpdateService(id string, TransactionBody models.Transaction) (*models.Transaction, error)
	DeleteService(id string) error
}

type transactionService struct {
	TransactionR repositories.TransactionRepository
}

func NewTransactionService(TransactionR repositories.TransactionRepository) TransactionService {
	return &transactionService{
		TransactionR: TransactionR,
	}
}

func (t *transactionService) GetTransactionsService() ([]*models.Transaction, error) {
	Transactions, err := t.TransactionR.GetTransactionsRepository()
	if err != nil {
		return nil, err
	}

	return Transactions, nil
}

func (t *transactionService) GetTransactionService(id string) (*models.Transaction, error) {
	Transaction, err := t.TransactionR.GetTransactionRepository(id)
	if err != nil {
		return nil, err
	}

	return Transaction, nil
}

func (t *transactionService) CreateService(Transaction models.Transaction) (*models.Transaction, error) {
	TransactionR, err := t.TransactionR.CreateRepository(Transaction)
	if err != nil {
		return nil, err
	}

	return TransactionR, nil
}

func (t *transactionService) UpdateService(id string, TransactionBody models.Transaction) (*models.Transaction, error) {
	Transaction, err := t.TransactionR.UpdateRepository(id, TransactionBody)
	if err != nil {
		return Transaction, err
	}

	return Transaction, nil
}

func (t *transactionService) DeleteService(id string) error {
	err := t.TransactionR.DeleteRepository(id)
	if err != nil {
		return err
	}

	return nil
}
