package models

import "gorm.io/gorm"

type Payment struct {
	gorm.Model
	Bukti_Pembayaran	string `json:"bukti_pembayaran" form:"bukti_pembayaran"`
	Total_Price			float64 `json:"total_price" form:"total_price"`
	Status				string `json:"status" form:"status"`
}
