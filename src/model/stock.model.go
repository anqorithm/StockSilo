package model

type Stock struct {
	ID           string  `json:"id" gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	Ticker       string  `json:"ticker"`
	Name         string  `json:"name"`
	Open         float64 `json:"open"`
	Close        float64 `json:"close"`
	High         float64 `json:"high"`
	Low          float64 `json:"low"`
	Volume       int     `json:"volume"`
	Date         string  `json:"date"`
	CurrentPrice float64 `json:"current_price"`
}
