package model

type Item struct {
	ItemName string `json:"item_name,omitempty" gorm:"unique"`
	Rate     int    `json:"rate,omitempty"`
	Qty      int    `json:"qty,omitempty"`
}
