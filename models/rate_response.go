package models

import "shipping-api/models/carriers"

type RateResponse struct {
	Success bool                    `json:"success"`
	Rates   []carriers.RateResponse `json:"rates"`
	Total   int                     `json:"total"`
}
