package model

type Item struct {
	ItemName string  `json:"item_name,omitempty"`
	Rate     float32 `json:"rate,omitempty"`
	Qty      float32 `json:"qty,omitempty"`
}
