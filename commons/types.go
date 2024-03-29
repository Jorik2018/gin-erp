package commons

import (
	"gorm.io/gorm"
)

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type Product struct {
	gorm.Model
	Code  string
	Price uint
}