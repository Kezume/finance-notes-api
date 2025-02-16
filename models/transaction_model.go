package models

import "time"

type Transaction struct {
	ID        int     `json:"id" gorm:"primaryKey"`
	Title     string  `json:"title"`
	Amount    float64 `json:"amount"`
	Type      string  `json:"type"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type TransactionRequest struct {
	Title  string  `json:"title" validate:"required"`
	Amount float64 `json:"amount" validate:"required,gt=0"`
	Type   string  `json:"type" validate:"required,oneof=income expense"`
}

type TransactionResponse struct {
	ID     uint    `json:"id"`
	Title  string  `json:"title"`
	Amount float64 `json:"amount"`
	Type   string  `json:"type"`
}
