package repositories

import (
	"mini_project/models"
	"fmt"

	"github.com/stretchr/testify/mock"
)

type TransactionRepositoryMock interface {
	GetTransactionsRepository() ([]*models.Transaction, error)
	GetTransactionRepository(id string) (*models.Transaction, error)
	CreateRepository(Transaction models.Transaction) (*models.Transaction, error)
	UpdateRepository(id string, TransactionBody models.Transaction) (*models.Transaction, error)
	DeleteRepository(id string) error
}

type ItransactionRepositoryMock struct {
	Mock mock.Mock
}

func NewTransactionRepositoryMock(mock mock.Mock) TransactionRepositoryMock {
	return &ItransactionRepositoryMock{
		Mock: mock,
	}
}

func (b *ItransactionRepositoryMock) GetTransactionsRepository() ([]*models.Transaction, error) {
	args := b.Mock.Called()
	if args.Get(0) == nil {
		return nil, args.Get(1).(error)
	}

	transactions := args.Get(0).([]*models.Transaction)

	return transactions, nil
}

func (b *ItransactionRepositoryMock) GetTransactionRepository(id string) (*models.Transaction, error) {
	args := b.Mock.Called(id)
	if args.Get(0) == nil {
		return nil, args.Get(1).(error)
	}

	transaction := args.Get(0).(models.Transaction)

	return &transaction, nil
}

func (u *ItransactionRepositoryMock) CreateRepository(transactionData models.Transaction) (*models.Transaction, error) {
	args := u.Mock.Called(transactionData)
	if args.Get(0) == nil {
		return nil, args.Get(1).(error)
	}

	transaction := args.Get(0).(models.Transaction)

	return &transaction, nil
}

func (u *ItransactionRepositoryMock) UpdateRepository(id string, transactionData models.Transaction) (*models.Transaction, error) {
	args := u.Mock.Called(id, transactionData)
	if args.Get(0) == nil {
		return nil,  args.Get(1).(error)
	}

	transaction := args.Get(0).(models.Transaction)

	return &transaction, nil
}

func (u *ItransactionRepositoryMock) DeleteRepository(id string) error {
	args := u.Mock.Called(id)
	if args.Get(0) != nil {
		return fmt.Errorf("must nil")
	}

	return nil
}
