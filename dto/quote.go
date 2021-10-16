package dto

import "time"

type QuoteRequest struct {
	SourceCurrency string  `json:"source_currency" binding:"required"`
	TargetCurrency string  `json:"target_currency" binding:"required"`
	Amount         float64 `json:"amount" binding:"required"`
	UserID         uint    `json:"user_id" binding:"required"`
}

type QuoteResponse struct {
	ID                    uint      `json:"id"`
	Fee                   float64   `json:"fee"`
	EstimatedDeliveryTime time.Time `json:"estimated_delivery_time"`
}
