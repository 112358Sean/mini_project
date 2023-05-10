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
	transactionRMock = &repositories.ItransactionRepositoryMock{Mock: mock.Mock{}}
	transactionSMock = NewTransactionService(transactionRMock)
)

func TestGetTransactionsService_Success(t *testing.T) {
	transactionsMP := []*models.Transaction{
		{
			ID_Keranjang: "1",
			Status: "pending",
		},
		{
			ID_Keranjang: "1",
			Status: "pending",
		},
	}

	transactionsM := []models.Transaction{
		{
			ID_Keranjang: "1",
			Status: "pending",
		},
		{
			ID_Keranjang: "1",
			Status: "pending",
		},
	}

	transactionRMock.Mock.On("GetTransactionsRepository").Return(transactionsMP)
	transactions, err := transactionSMock.GetTransactionsService()

	assert.Nil(t, err)
	assert.NotNil(t, transactions)

	assert.Equal(t, transactionsM[0].ID_Keranjang, transactions[0].ID_Keranjang)
	assert.Equal(t, transactionsM[0].Status, transactions[0].Status)
}

func TestGetTransactionsService_Failure(t *testing.T) {
	transactionRMock = &repositories.ItransactionRepositoryMock{Mock: mock.Mock{}}
	transactionSMock = NewTransactionService(transactionRMock)
	transactionRMock.Mock.On("GetTransactionsRepository").Return(nil, errors.New("get all transactions failed"))
	transactions, err := transactionSMock.GetTransactionsService()

	assert.Nil(t, transactions)
	assert.NotNil(t, err)
}

func TestGetTransactionService_Success(t *testing.T) {
	transaction := models.Transaction{
		ID_Keranjang: "1",
		Status: "pending",
	}

	transactionRMock.Mock.On("GetTransactionRepository", "1").Return(transaction, nil)
	transactions, err := transactionSMock.GetTransactionService("1")

	assert.Nil(t, err)
	assert.NotNil(t, transactions)

	assert.Equal(t, transaction.ID_Keranjang, transactions.ID_Keranjang)
	assert.Equal(t, transaction.Status, transactions.Status)
}

func TestGetTransactionService_Failure(t *testing.T) {
	transactionRMock.Mock.On("GetTransactionRepository", "3").Return(nil, fmt.Errorf("transaction not found"))
	transaction, err := transactionSMock.GetTransactionService("3")

	assert.NotNil(t, err)
	assert.Nil(t, transaction)
}

func TestCreatetransactionService_Success(t *testing.T) {
	transaction := models.Transaction{
		ID_Keranjang: "1",
		Status: "pending",
	}

	transactionRMock.Mock.On("CreateRepository", transaction).Return(transaction, nil)
	transactions, err := transactionSMock.CreateService(transaction)

	assert.Nil(t, err)
	assert.NotNil(t, transactions)

	assert.Equal(t, transaction.ID_Keranjang, transactions.ID_Keranjang)
	assert.Equal(t, transaction.Status, transactions.Status)
}

func TestCreateTransactionService_Failure(t *testing.T) {
	transaction := models.Transaction{
		ID_Keranjang: "1",
		Status: "pending",
	}

	transactionRMock.Mock.On("CreateRepository", transaction).Return(nil, fmt.Errorf("create transaction failed"))
	transactions, err := transactionSMock.CreateService(transaction)

	assert.Nil(t, transactions)
	assert.NotNil(t, err)
}

func TestUpdatetransactionService_Success(t *testing.T) {
	transaction := models.Transaction{
		ID_Keranjang: "1",
		Status: "pending",
	}

	transactionRMock.Mock.On("UpdateRepository", "1", transaction).Return(transaction, nil)
	transactions, err := transactionSMock.UpdateService("1", transaction)

	assert.Nil(t, err)
	assert.NotNil(t, transactions)

	assert.Equal(t, transaction.ID, transactions.ID)
	assert.Equal(t, transaction.ID_Keranjang, transactions.ID_Keranjang)
	assert.Equal(t, transaction.Status, transactions.Status)
}

func TestUpdateTransactionService_Failure(t *testing.T) {
	transaction := models.Transaction{
		ID_Keranjang: "1",
		Status: "pending",
	}

	transactionRMock.Mock.On("UpdateRepository", "2", transaction).Return(nil, fmt.Errorf("transaction not found"))
	transactions, err := transactionSMock.UpdateService("2", transaction)

	assert.Nil(t, transactions)
	assert.NotNil(t, err)
}

func TestDeleteTransactionService_Success(t *testing.T) {
	transactionRMock.Mock.On("DeleteRepository", "1").Return(nil)
	err := transactionSMock.DeleteService("1")

	assert.Nil(t, err)
}

func TestDeleteTransactionService_Failure(t *testing.T) {
	transactionRMock.Mock.On("DeleteRepository", "2").Return(fmt.Errorf("transaction not found"))
	err := transactionSMock.DeleteService("2")

	assert.NotNil(t, err)
}
