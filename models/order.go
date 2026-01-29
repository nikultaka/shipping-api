package models

import (
	"time"

	"github.com/google/uuid"
)

type Order struct {
	ID           uuid.UUID `gorm:"type:char(36);primaryKey" json:"id"`
	OrderID      string    `gorm:"type:varchar(100);uniqueIndex;not null" json:"order_id"`
	Carrier      string    `gorm:"type:varchar(50);not null" json:"carrier"`
	TrackingID   string    `gorm:"type:varchar(100)" json:"tracking_id"`
	FromAddress  string    `gorm:"type:text;not null" json:"from_address"`
	ToAddress    string    `gorm:"type:text;not null" json:"to_address"`
	Weight       float64   `gorm:"type:decimal(10,2)" json:"weight"`
	Status       string    `gorm:"type:varchar(50);default:'created'" json:"status"`
	ExpectedDate time.Time `json:"expected_date"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
