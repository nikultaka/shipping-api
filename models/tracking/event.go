package tracking

import (
	"time"

	"github.com/google/uuid"
)

type TrackingEvent struct {
	ID         uuid.UUID `gorm:"type:char(36);primaryKey" json:"id"`
	OrderID    string    `gorm:"type:varchar(100);index;not null" json:"order_id"`
	TrackingID string    `gorm:"type:varchar(100);index;not null" json:"tracking_id"`
	Carrier    string    `gorm:"type:varchar(50);not null" json:"carrier"`
	Event      string    `gorm:"type:varchar(200);not null" json:"event"`
	Location   string    `gorm:"type:varchar(200)" json:"location"`
	Status     string    `gorm:"type:varchar(50)" json:"status"`
	EventTime  time.Time `json:"event_time"`
	CreatedAt  time.Time `json:"created_at"`
}

func (TrackingEvent) TableName() string {
	return "tracking_events"
}
