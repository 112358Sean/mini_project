package models

import "gorm.io/gorm"

type Transaction struct {
	gorm.Model
	ID_Transaksi 	string `json:"id_transaksi" form:"id_transaksi"`
	ID_Keranjang	string `json:"id_keranjang" form:"id_keranjang"`
	Status			string `json:"status" form:"status"`
}
