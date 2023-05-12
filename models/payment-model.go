package models

import "gorm.io/gorm"

type Payment struct {
	gorm.Model
	ID_Transaksi		string `json:"id_transaksi" form:"id_transaksi"`
	Bukti_Pembayaran	string `json:"bukti_pembayaran" form:"bukti_pembayaran"`
	Total_Price			float64 `json:"total_price" form:"total_price"`
	Status				string `json:"status" form:"status"`
}
