package model

import "gorm.io/gorm"

type Customer struct {
	gorm.Model
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
	Phone     uint64 `json:"phone,omitempty"`
	Email     string `json:"email,omitempty" gorm:"unique"`
}
