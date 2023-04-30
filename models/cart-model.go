package models

import "gorm.io/gorm"

type Cart struct {
	gorm.Model
	ID_User			string 	`json:"id_user" form:"id_user"`
	ID_Buku			string 	`json:"id_buku" form:"id_buku"`
	Jumlah			int64 	`json:"jumlah" form:"jumlah"`	
}
