package models

import "time"

type Quote struct {
	ID                    uint      `json:"id"`
	CurrencySource        string    `json:"currency_source"`
	CurrencyTarget        string    `json:"currency_target"`
	Amount                float64   `json:"amount"`
	Fee                   float64   `json:"fee"`
	EstimatedDeliveryTime time.Time `json:"estimated_delivery_time"`
	UserID                uint      `json:"user_id"`
}

func (t *Quote) TableName() string {
	return "quotes"
}
