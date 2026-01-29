package services

import (
	"log"
	"shipping-api/db"
	"shipping-api/lib/helpers"
	"shipping-api/models"
	"shipping-api/models/tracking"
	"sync"
	"time"
)

func PollTracking() ([]tracking.TrackingResponse, error) {
	log.Println("[TrackingPoller] Starting tracking poll...")

	var orders []models.Order
	if err := db.DB.Where("status != ?", "delivered").
		Limit(10).
		Find(&orders).Error; err != nil {
		log.Printf("[TrackingPoller] Error fetching orders: %v", err)
		return nil, err
	}

	if len(orders) == 0 {
		log.Println("[TrackingPoller] No orders to track")
		return []tracking.TrackingResponse{}, nil
	}

	log.Printf("[TrackingPoller] Found %d orders to track", len(orders))

	resultsChan := make(chan tracking.TrackingResponse, len(orders))
	var wg sync.WaitGroup

	for _, order := range orders {
		wg.Add(1)
		go func(o models.Order) {
			defer wg.Done()

			events, err := pollOrderTracking(o)
			if err != nil {
				log.Printf("[TrackingPoller] Error polling order %s: %v", o.OrderID, err)
				return
			}

			if len(events) > 0 {
				response := tracking.TrackingResponse{
					TrackingID: o.TrackingID,
					Carrier:    o.Carrier,
					Events:     events,
					Status:     getLatestStatus(events),
					LastUpdate: time.Now(),
				}
				resultsChan <- response
			}
		}(order)
	}

	go func() {
		wg.Wait()
		close(resultsChan)
	}()

	var responses []tracking.TrackingResponse
	for response := range resultsChan {
		responses = append(responses, response)
	}

	log.Printf("[TrackingPoller] Polling complete. Updated %d orders", len(responses))
	return responses, nil
}

func pollOrderTracking(order models.Order) ([]tracking.TrackingEvent, error) {
	var existingEvents []tracking.TrackingEvent
	db.DB.Where("tracking_id = ?", order.TrackingID).
		Find(&existingEvents)

	existingEventMap := make(map[string]bool)
	for _, event := range existingEvents {
		key := event.Event + "|" + event.EventTime.Format("2006-01-02 15:04")
		existingEventMap[key] = true
	}

	events, err := helpers.PollCarrierTracking(order.TrackingID, order.Carrier)
	if err != nil {
		return nil, err
	}

	var newEvents []tracking.TrackingEvent
	for _, event := range events {
		key := event.Event + "|" + event.EventTime.Format("2006-01-02 15:04")

		if !existingEventMap[key] {
			if err := db.DB.Create(&event).Error; err != nil {
				log.Printf("[TrackingPoller] Error saving event: %v", err)
				continue
			}
			newEvents = append(newEvents, event)
			log.Printf("[TrackingPoller] New event added: %s - %s", order.TrackingID, event.Event)
		}
	}

	if len(newEvents) > 0 {
		latestStatus := getLatestStatus(events)
		db.DB.Model(&order).Update("status", latestStatus)
	}

	return newEvents, nil
}

func getLatestStatus(events []tracking.TrackingEvent) string {
	if len(events) == 0 {
		return "unknown"
	}

	latest := events[0]
	for _, event := range events {
		if event.EventTime.After(latest.EventTime) {
			latest = event
		}
	}
	return latest.Status
}
