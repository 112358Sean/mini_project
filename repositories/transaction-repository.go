package repositories

import (
	"mini_project/models"

	"gorm.io/gorm"
)

type TransactionRepository interface {
	GetTransactionsRepository() ([]*models.Transaction, error)
	GetTransactionRepository(id string) (*models.Transaction, error)
	CreateRepository(Transaction models.Transaction) (*models.Transaction, error)
	UpdateRepository(id string, TransactionBody models.Transaction) (*models.Transaction, error)
	DeleteRepository(id string) error
}

type transactionRepository struct {
	DB *gorm.DB
}

func NewTransactionRepository(DB *gorm.DB) TransactionRepository {
	return &transactionRepository{
		DB: DB,
	}
}

func (t *transactionRepository) GetTransactionsRepository() ([]*models.Transaction, error) {
	var Transactions []*models.Transaction

	if err := t.DB.Find(&Transactions).Error; err != nil {
		return nil, err
	}

	return Transactions, nil
}

func (t *transactionRepository) GetTransactionRepository(id string) (*models.Transaction, error) {
	var Transaction *models.Transaction

	if err := t.DB.Where("id = ?", id).Take(&Transaction).Error; err != nil {
		return nil, err
	}

	return Transaction, nil
}

func (t *transactionRepository) CreateRepository(Transaction models.Transaction) (*models.Transaction, error) {
	if err := t.DB.Save(&Transaction).Error; err != nil {
		return nil, err
	}

	return &Transaction, nil
}

func (t *transactionRepository) UpdateRepository(id string, TransactionBody models.Transaction) (*models.Transaction, error) {
	Transaction, err := t.GetTransactionRepository(id)
	if err != nil {
		return nil, err
	}

	err = t.DB.Where("ID = ?", id).Updates(models.Transaction{ID_User: TransactionBody.ID_User, ID_Keranjang: TransactionBody.ID_Keranjang, Status: TransactionBody.Status}).Error
	if err != nil {
		return nil, err
	}

	Transaction.ID_User = TransactionBody.ID_User
	Transaction.ID_Keranjang = TransactionBody.ID_Keranjang
	Transaction.Status = TransactionBody.Status

	return Transaction, nil
}

func (t *transactionRepository) DeleteRepository(id string) error {
	_, err := t.GetTransactionRepository(id)
	if err != nil {
		return err
	}

	if err := t.DB.Delete(&models.Transaction{}, id).Error; err != nil {
		return err
	}

	return nil
}
