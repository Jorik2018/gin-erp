package models

import (
	"gorm.io/gorm"
)
	
type Student struct {
	gorm.Model
	Name      string   `json:"name,omitempty"`
	Age       int      `json:"age,omitempty"`
	Address   *Address `json:"address,omitempty" gorm:"embedded"`
}

type Address struct {
	StreetAddress  string `json:"streetAddress,omitempty"`
	StreetAddress2 string `json:"streetAddress2,omitempty"`
	City           string `json:"city,omitempty"`
	State          string `json:"state,omitempty"`
	ZipCode        string `json:"zipCode,omitempty"`
	Country        string `json:"country,omitempty"`
}
