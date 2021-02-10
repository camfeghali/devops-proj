package model

import (

	"gorm.io/gorm"

)

type User struct {
	gorm.Model
	ID int `json:"id"`
	Username string `json:"name"`
	Password string `json:"password"`
}

type NewUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
}