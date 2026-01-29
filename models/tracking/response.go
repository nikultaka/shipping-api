package tracking

import "time"

type TrackingResponse struct {
	TrackingID string          `json:"tracking_id"`
	Carrier    string          `json:"carrier"`
	Events     []TrackingEvent `json:"events"`
	Status     string          `json:"status"`
	LastUpdate time.Time       `json:"last_update"`
}
