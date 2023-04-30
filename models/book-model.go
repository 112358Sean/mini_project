package models

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	Judul		string `json:"judul" form:"judul"`
	Penulis		string `json:"penulis" form:"penulis"`
	Penerbit	string `json:"penerbit" form:"penerbit"`
	Harga		int64 `json:"harga" form:"harga"`
	Stok		int64 `json:"stok" form:"stok"`
}
