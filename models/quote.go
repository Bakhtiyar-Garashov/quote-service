package models

import "time"

type Quote struct {
	ID                    uint      `gorm:"column:id;primary_key"`
	CurrencySource        string    `gorm:"column:currency_source"`
	CurrencyTarget        string    `gorm:"column:currency_target"`
	Amount                float64   `gorm:"column:amount"`
	Fee                   float64   `gorm:"column:fee"`
	EstimatedDeliveryTime time.Time `gorm:"column:estimated_delivery_time"`
	UserID                uint      `gorm:"column:user_id"`
}

func (t *Quote) TableName() string {
	return "quotes"
}
