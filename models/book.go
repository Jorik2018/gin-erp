package models

import (
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Name 	string `json:"name,omitempty" gorm:"not null" binding:"required"`
	Author 	string `json:"author,omitempty" gorm:"not null" binding:"required"`
}
