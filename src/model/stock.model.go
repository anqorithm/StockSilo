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

type CreateStockRequest struct {
	Ticker       string  `json:"ticker" validate:"required,alphanum"`
	Name         string  `json:"name" validate:"required"`
	Open         float64 `json:"open" validate:"required,gt=0"`
	Close        float64 `json:"close" validate:"required,gt=0"`
	High         float64 `json:"high" validate:"required,gt=0"`
	Low          float64 `json:"low" validate:"required,gt=0"`
	Volume       int     `json:"volume" validate:"required,min=1"`
	Date         string  `json:"date" validate:"required"`
	CurrentPrice float64 `json:"current_price" validate:"required,gt=0"`
}

type UpdateStockRequest struct {
	Ticker       *string  `json:"ticker" validate:"omitempty"`
	Name         *string  `json:"name" validate:"omitempty"`
	Open         *float64 `json:"open" validate:"omitempty"`
	Close        *float64 `json:"close" validate:"omitempty"`
	High         *float64 `json:"high" validate:"omitempty"`
	Low          *float64 `json:"low" validate:"omitempty"`
	Volume       *int     `json:"volume" validate:"omitempty"`
	Date         *string  `json:"date" validate:"omitempty"`
	CurrentPrice *float64 `json:"current_price" validate:"omitempty"`
}
