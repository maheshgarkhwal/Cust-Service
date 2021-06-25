package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserName        string `json:"user_name,omitempty" gorm:"unique"`
	Password        string `json:"password,omitempty"`
	ConfirmPassword string `json:"confirm_password,omitempty"`
}
