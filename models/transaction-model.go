package models

import "gorm.io/gorm"

type Transaction struct {
	gorm.Model
	ID_Keranjang	string `json:"id_keranjang" form:"id_keranjang"`
	Status			string `json:"status" form:"status"`
}
