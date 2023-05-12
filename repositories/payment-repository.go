package repositories

import (
	"mini_project/models"

	"gorm.io/gorm"
)

type PaymentRepository interface {
	GetPaymentsRepository() ([]*models.Payment, error)
	GetPaymentRepository(id string) (*models.Payment, error)
	CreateRepository(Payment models.Payment) (*models.Payment, error)
	UpdateRepository(id string, PaymentBody models.Payment) (*models.Payment, error)
	DeleteRepository(id string) error
	CheckOut(id_user string) (*models.Payment, error)
}

type paymentRepository struct {
	DB *gorm.DB
}

func NewPaymentRepository(DB *gorm.DB) PaymentRepository {
	return &paymentRepository{
		DB: DB,
	}
}

func (p *paymentRepository) GetPaymentsRepository() ([]*models.Payment, error) {
	var Payments []*models.Payment

	if err := p.DB.Find(&Payments).Error; err != nil {
		return nil, err
	}

	return Payments, nil
}

func (p *paymentRepository) GetPaymentRepository(id string) (*models.Payment, error) {
	var Payment *models.Payment

	if err := p.DB.Where("ID = ?", id).Take(&Payment).Error; err != nil {
		return nil, err
	}

	return Payment, nil
}

func (p *paymentRepository) CreateRepository(Payment models.Payment) (*models.Payment, error) {
	if err := p.DB.Save(&Payment).Error; err != nil {
		return nil, err
	}

	return &Payment, nil
}

func (p *paymentRepository) UpdateRepository(id string, PaymentBody models.Payment) (*models.Payment, error) {
	Payment, err := p.GetPaymentRepository(id)
	if err != nil {
		return nil, err
	}

	err = p.DB.Where("ID = ?", id).Updates(models.Payment{ID_Transaksi: PaymentBody.ID_Transaksi, Bukti_Pembayaran: PaymentBody.Bukti_Pembayaran, Status: PaymentBody.Status}).Error
	if err != nil {
		return nil, err
	}

	Payment.ID_Transaksi = PaymentBody.ID_Transaksi
	Payment.Bukti_Pembayaran = PaymentBody.Bukti_Pembayaran
	Payment.Status = "Terbayar"

	return Payment, nil
}

func (p *paymentRepository) DeleteRepository(id string) error {
	_, err := p.GetPaymentRepository(id)
	if err != nil {
		return err
	}

	if err := p.DB.Delete(&models.Payment{}, id).Error; err != nil {
		return err
	}

	return nil
}

func (p *paymentRepository) CheckOut(id_user string) (*models.Payment, error) {
	var totalPrice float64
	CheckOut := models.Payment{}

	row := p.DB.
		Table("books").
		Select("SUM(books.harga * carts.jumlah) as totalPrice").
		Joins("JOIN carts ON carts.id_buku = books.id").
		Where("carts.id_user = ?", id_user).
		Row()

	err := row.Scan(&totalPrice)
	if err != nil {
		return nil, err
	}

	CheckOut.Total_Price = totalPrice

	if err := p.DB.Save(&CheckOut).Error; err != nil {
		return nil, err
	}

	return &CheckOut, nil
}
