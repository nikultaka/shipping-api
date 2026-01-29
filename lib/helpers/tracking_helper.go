package helpers

import (
	"errors"
	"math/rand"
	"shipping-api/models/tracking"
	"time"

	"github.com/google/uuid"
)

func PollCarrierTracking(trackingID, carrier string) ([]tracking.TrackingEvent, error) {
	rand.Seed(time.Now().UnixNano() + int64(len(trackingID)))

	if rand.Intn(100) < 20 {
		return nil, Error("carrier tracking API failed")
	}

	events := []struct {
		event    string
		location string
		status   string
	}{
		{"Order Created", "Mumbai", "created"},
		{"Picked Up", "Mumbai", "in_transit"},
		{"In Transit", "Delhi", "in_transit"},
		{"Out for Delivery", "Delhi", "out_for_delivery"},
		{"Delivered", "Delhi", "delivered"},
		{"Exception", "Mumbai", "exception"},
		{"Returned", "Mumbai", "returned"},
	}

	numEvents := 1 + rand.Intn(3)
	var trackingEvents []tracking.TrackingEvent

	for i := 0; i < numEvents; i++ {
		eventIndex := rand.Intn(len(events))
		selectedEvent := events[eventIndex]

		daysAgo := 1 + rand.Intn(5)
		hoursAgo := rand.Intn(24)
		eventTime := time.Now().Add(-time.Duration(daysAgo)*24*time.Hour - time.Duration(hoursAgo)*time.Hour)

		trackingEvent := tracking.TrackingEvent{
			ID:         uuid.New(),
			OrderID:    "ORD" + trackingID[3:6],
			TrackingID: trackingID,
			Carrier:    carrier,
			Event:      selectedEvent.event,
			Location:   selectedEvent.location,
			Status:     selectedEvent.status,
			EventTime:  eventTime,
			CreatedAt:  time.Now(),
		}

		trackingEvents = append(trackingEvents, trackingEvent)
	}

	return trackingEvents, nil
}

func Error(msg string) error {
	return errors.New(msg)
}
