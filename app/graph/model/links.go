package model

import (

	"gorm.io/gorm"

)

type Link struct {
	gorm.Model
	ID      int `json:"id"`
	Title   string `json:"title"`
	Address string `json:"address"`
	UserID	int	   `json:"userId"`
	User    *User  `json:"user"`
}

type NewLink struct {
	Title   string `json:"title"`
	Address string `json:"address"`
}

