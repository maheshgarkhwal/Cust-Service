package model

import "gorm.io/gorm"

type Item struct {
	gorm.Model
	ItemName string `json:"item_name,omitempty" gorm:"unique"`
	Rate     int    `json:"rate,omitempty"`
	Qty      int    `json:"qty,omitempty"`
}
