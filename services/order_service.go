package services

import (
	"errors"
	"fmt"
	"math/rand"
	"time"

	"shipping-api/models"
	"shipping-api/repository"
)

func CreateOrderService(req models.CreateOrderRequest) (*models.CreateOrderResponse, error) {
	exists, err := repository.OrderExists(req.OrderID)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, errors.New("order ID already exists")
	}

	if !isCarrierServiceable(req.Carrier, req.FromAddress, req.ToAddress) {
		return nil, errors.New("carrier not serviceable for this route")
	}

	trackingID := generateTrackingID(req.Carrier)

	rand.Seed(time.Now().UnixNano())
	days := 3 + rand.Intn(5) // 3 to 7 days
	expectedDate := time.Now().AddDate(0, 0, days)

	order := &models.Order{
		OrderID:      req.OrderID,
		Carrier:      req.Carrier,
		TrackingID:   trackingID,
		FromAddress:  req.FromAddress,
		ToAddress:    req.ToAddress,
		Weight:       req.Weight,
		Status:       "created",
		ExpectedDate: expectedDate,
	}

	err = repository.CreateOrder(order)
	if err != nil {
		return nil, err
	}

	response := &models.CreateOrderResponse{
		OrderID:      order.OrderID,
		TrackingID:   order.TrackingID,
		ExpectedDate: order.ExpectedDate.Format("2006-01-02"),
		Carrier:      order.Carrier,
		Status:       order.Status,
	}

	return response, nil
}

func generateTrackingID(carrier string) string {
	rand.Seed(time.Now().UnixNano())
	randomNum := 100000 + rand.Intn(900000)

	switch carrier {
	case "bluedart":
		return fmt.Sprintf("BLD%d", randomNum)
	case "delhivery":
		return fmt.Sprintf("DLV%d", randomNum)
	case "xpressbees":
		return fmt.Sprintf("XPB%d", randomNum)
	default:
		return fmt.Sprintf("TRK%d", randomNum)
	}
}

func isCarrierServiceable(carrier, from, to string) bool {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(100) < 90
}
