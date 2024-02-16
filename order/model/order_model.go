package model

import "time"

// Requests

type OrderCreateOrUpdateModel struct {
	ID    uint    `json:"id"`
	Total float64 `json:"total"`
}

// Responses
type OrderModel struct {
	ID        string    `json:"id"`
	Total     float64   `json:"total"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
