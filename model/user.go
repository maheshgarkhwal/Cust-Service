package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email    string `json:"email,omitempty" gorm:"unique"`
	Password string `json:"password,omitempty"`
}
