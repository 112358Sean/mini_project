package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Nama     string `json:"nama" form:"nama"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	Alamat 	 string `json:"alamat" form:"alamat"`
	Role	 string `json:"role" form:"role"`
}

type CreateUser struct {
	User  *User   `json:"user" form:"user"`
	Token string `json:"token" form:"token"`
}

