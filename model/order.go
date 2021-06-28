package model

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	Item     string `json:"item,omitempty"`
	Customer string `json:"customer,omitempty"`
	Qty      int    `json:"qty,omitempty"`
	Rate     int    `json:"rate,omitempty"`
}
