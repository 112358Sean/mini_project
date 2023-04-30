package models

import "gorm.io/gorm"

type Transaction struct {
	gorm.Model
	ID_User			string `json:"id_user" form:"id_user"`
	ID_Keranjang	string `json:"id_keranjang" form:"id_keranjang"`
	Status			string `json:"status" form:"status"`
}
