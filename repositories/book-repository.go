package repositories

import (
	"mini_project/models"

	"gorm.io/gorm"
)

type BookRepository interface {
	GetBooksRepository() ([]*models.Book, error)
	GetBookRepository(id string) (*models.Book, error)
	CreateRepository(Book models.Book) (*models.Book, error)
	UpdateRepository(id string, BookBody models.Book) (*models.Book, error)
	DeleteRepository(id string) error
}

type bookRepository struct {
	DB *gorm.DB
}

func NewBookRepository(DB *gorm.DB) BookRepository {
	return &bookRepository{
		DB: DB,
	}
}

func (b *bookRepository) GetBooksRepository() ([]*models.Book, error) {
	var Books []*models.Book

	if err := b.DB.Find(&Books).Error; err != nil {
		return nil, err
	}

	return Books, nil
}

func (b *bookRepository) GetBookRepository(id string) (*models.Book, error) {
	var Book *models.Book

	if err := b.DB.Where("ID = ?", id).Take(&Book).Error; err != nil {
		return nil, err
	}

	return Book, nil
}

func (b *bookRepository) CreateRepository(Book models.Book) (*models.Book, error) {
	if err := b.DB.Save(&Book).Error; err != nil {
		return nil, err
	}

	return &Book, nil
}

func (b *bookRepository) UpdateRepository(id string, BookBody models.Book) (*models.Book, error) {
	Book, err := b.GetBookRepository(id)
	if err != nil {
		return nil, err
	}

	err = b.DB.Where("ID = ?", id).Updates(models.Book{Judul: BookBody.Judul, Penulis: BookBody.Penulis, Penerbit: BookBody.Penerbit, Harga: BookBody.Harga, Stok: BookBody.Stok}).Error
	if err != nil {
		return nil, err
	}

	Book.Judul = BookBody.Judul
	Book.Penulis = BookBody.Penulis
	Book.Penerbit = BookBody.Penerbit
	Book.Harga = BookBody.Harga
	Book.Stok = BookBody.Stok

	return Book, nil
}

func (b *bookRepository) DeleteRepository(id string) error {
	_, err := b.GetBookRepository(id)
	if err != nil {
		return err
	}

	if err := b.DB.Delete(&models.Book{}, id).Error; err != nil {
		return err
	}

	return nil
}
