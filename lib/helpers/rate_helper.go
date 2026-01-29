package helpers

import (
	"math/rand"
	"shipping-api/models/carriers"
	"time"
)

func FetchCarrierRates(carrier string, fromPincode, toPincode string, weight float64) carriers.RateResponse {
	rand.Seed(time.Now().UnixNano() + int64(len(carrier)))

	randomCondition := rand.Intn(100)

	if randomCondition < 10 {
		return carriers.RateResponse{
			Carrier:       carrier,
			IsError:       true,
			ErrorMessage:  "Carrier API failed to respond",
			IsServiceable: false,
		}
	}

	if randomCondition < 30 {
		delay := 1 + rand.Intn(3)
		time.Sleep(time.Duration(delay) * time.Second)
	}

	if randomCondition < 45 {
		return carriers.RateResponse{
			Carrier:       carrier,
			IsError:       true,
			ErrorMessage:  "Location not serviceable",
			IsServiceable: false,
		}
	}

	basePrice := weight * 10.0
	distanceMultiplier := 1.0 + (rand.Float64() * 2.0)
	finalPrice := basePrice * distanceMultiplier

	finalPrice = float64(int(finalPrice*100)) / 100

	days := 1 + rand.Intn(7)
	etd := time.Now().AddDate(0, 0, days).Format("2006-01-02")

	serviceTypes := []string{"Standard", "Express", "Premium", "Next Day"}
	serviceType := serviceTypes[rand.Intn(len(serviceTypes))]

	return carriers.RateResponse{
		Carrier:       carrier,
		Price:         finalPrice,
		ETD:           etd,
		ServiceType:   serviceType,
		IsError:       false,
		IsServiceable: true,
	}
}
