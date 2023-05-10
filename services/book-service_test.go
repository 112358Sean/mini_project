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
	bookRMock = &repositories.IbookRepositoryMock{Mock: mock.Mock{}}
	bookSMock = NewBookService(bookRMock)
)

func TestGetBooksService_Success(t *testing.T) {
	booksMP := []*models.Book{
		{
			Judul:       "Buku 1",
			Penulis:      "Boy",
			Penerbit: "Gramedia",
			Harga: 20000,
			Stok: 15,
		},
		{
			Judul:       "Buku 1",
			Penulis:      "Boy",
			Penerbit: "Gramedia",
			Harga: 20000,
			Stok: 15,
		},
	}

	booksM := []models.Book{
		{
			Judul:       "Buku 1",
			Penulis:      "Boy",
			Penerbit: "Gramedia",
			Harga: 20000,
			Stok: 15,
		},
		{
			Judul:       "Buku 1",
			Penulis:      "Boy",
			Penerbit: "Gramedia",
			Harga: 20000,
			Stok: 15,
		},
	}

	bookRMock.Mock.On("GetBooksRepository").Return(booksMP)
	books, err := bookSMock.GetBooksService()

	assert.Nil(t, err)
	assert.NotNil(t, books)

	assert.Equal(t, booksM[0].Judul, books[0].Judul)
	assert.Equal(t, booksM[0].Penulis, books[0].Penulis)
	assert.Equal(t, booksM[0].Penerbit, books[0].Penerbit)
	assert.Equal(t, booksM[0].Harga, books[0].Harga)
	assert.Equal(t, booksM[0].Stok, books[0].Stok)
}

func TestGetBooksService_Failure(t *testing.T) {
	bookRMock = &repositories.IbookRepositoryMock{Mock: mock.Mock{}}
	bookSMock = NewBookService(bookRMock)
	bookRMock.Mock.On("GetBooksRepository").Return(nil, errors.New("get all books failed"))
	books, err := bookSMock.GetBooksService()

	assert.Nil(t, books)
	assert.NotNil(t, err)
}

func TestGetBookService_Success(t *testing.T) {
	book := models.Book{
		Judul:       "Buku 1",
		Penulis:      "Boy",
		Penerbit: "Gramedia",
		Harga: 20000,
		Stok: 15,
	}

	bookRMock.Mock.On("GetBookRepository", "1").Return(book, nil)
	books, err := bookSMock.GetBookService("1")

	assert.Nil(t, err)
	assert.NotNil(t, books)

	assert.Equal(t, book.Judul, books.Judul)
	assert.Equal(t, book.Penulis, books.Penulis)
	assert.Equal(t, book.Penerbit, books.Penerbit)
	assert.Equal(t, book.Harga, books.Harga)
	assert.Equal(t, book.Stok, books.Stok)
}

func TestGetBookService_Failure(t *testing.T) {
	bookRMock.Mock.On("GetBookRepository", "3").Return(nil, fmt.Errorf("book not found"))
	book, err := bookSMock.GetBookService("3")

	assert.NotNil(t, err)
	assert.Nil(t, book)
}

func TestCreateBookService_Success(t *testing.T) {
	book := models.Book{
		Judul:       "Buku 1",
		Penulis:      "Boy",
		Penerbit: "Gramedia",
		Harga: 20000,
		Stok: 15,
	}

	bookRMock.Mock.On("CreateRepository", book).Return(book, nil)
	books, err := bookSMock.CreateService(book)

	assert.Nil(t, err)
	assert.NotNil(t, books)

	assert.Equal(t, book.Judul, books.Judul)
	assert.Equal(t, book.Penulis, books.Penulis)
	assert.Equal(t, book.Penerbit, books.Penerbit)
	assert.Equal(t, book.Harga, books.Harga)
	assert.Equal(t, book.Stok, books.Stok)
}

func TestCreateBookService_Failure(t *testing.T) {
	book := models.Book{
		Judul:       "Buku 1",
		Penulis:      "Boy",
		Penerbit: "Gramedia",
		Harga: 20000,
		Stok: 15,
	}

	bookRMock.Mock.On("CreateRepository", book).Return(nil, fmt.Errorf("create book failed"))
	books, err := bookSMock.CreateService(book)

	assert.Nil(t, books)
	assert.NotNil(t, err)
}

func TestUpdateBookService_Success(t *testing.T) {
	book := models.Book{
		Judul:       "Buku 1",
			Penulis:      "Boy",
			Penerbit: "Gramedia",
			Harga: 20000,
			Stok: 15,
	}

	bookRMock.Mock.On("UpdateRepository", "1", book).Return(book, nil)
	books, err := bookSMock.UpdateService("1", book)

	assert.Nil(t, err)
	assert.NotNil(t, books)

	assert.Equal(t, book.ID, books.ID)
	assert.Equal(t, book.Judul, books.Judul)
	assert.Equal(t, book.Penulis, books.Penulis)
	assert.Equal(t, book.Penerbit, books.Penerbit)
	assert.Equal(t, book.Harga, books.Harga)
	assert.Equal(t, book.Stok, books.Stok)
}

func TestUpdateBookService_Failure(t *testing.T) {
	book := models.Book{
		Judul:       "Buku 1",
		Penulis:      "Boy",
		Penerbit: "Gramedia",
		Harga: 20000,
		Stok: 15,
	}

	bookRMock.Mock.On("UpdateRepository", "2", book).Return(nil, fmt.Errorf("book not found"))
	books, err := bookSMock.UpdateService("2", book)

	assert.Nil(t, books)
	assert.NotNil(t, err)
}

func TestDeleteBookService_Success(t *testing.T) {
	bookRMock.Mock.On("DeleteRepository", "1").Return(nil)
	err := bookSMock.DeleteService("1")

	assert.Nil(t, err)
}

func TestDeleteBookService_Failure(t *testing.T) {
	bookRMock.Mock.On("DeleteRepository", "2").Return(fmt.Errorf("book not found"))
	err := bookSMock.DeleteService("2")

	assert.NotNil(t, err)
}
