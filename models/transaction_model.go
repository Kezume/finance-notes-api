package models

import "time"

type Transaction struct {
	Id        int       `json:"id" gorm:"primaryKey"`
	Title     string    `json:"title"`
	Amount    float64   `json:"amount"`
	Type      string    `json:"type"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
