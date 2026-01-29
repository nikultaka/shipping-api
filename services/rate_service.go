package services

import (
	"shipping-api/lib/helpers"
	"shipping-api/models"
	"shipping-api/models/carriers"
	"sync"
	"time"
)

func CalculateRatesService(req models.RateRequest) models.RateResponse {
	ratesChan := make(chan carriers.RateResponse, 3)
	var wg sync.WaitGroup

	carriersList := []string{"bluedart", "delhivery", "xpressbees"}

	for _, carrier := range carriersList {
		wg.Add(1)
		go func(c string) {
			defer wg.Done()

			rate := helpers.FetchCarrierRates(
				c,
				req.FromPincode,
				req.ToPincode,
				req.Weight,
			)
			ratesChan <- rate
		}(carrier)
	}

	go func() {
		wg.Wait()
		close(ratesChan)
	}()

	var allRates []carriers.RateResponse
	var serviceableRates []carriers.RateResponse

	timeout := time.After(5 * time.Second)

collectLoop:
	for {
		select {
		case rate, ok := <-ratesChan:
			if !ok {
				break collectLoop
			}
			allRates = append(allRates, rate)

			if rate.IsServiceable && !rate.IsError {
				serviceableRates = append(serviceableRates, rate)
			}

		case <-timeout:
			break collectLoop
		}
	}

	return models.RateResponse{
		Success: true,
		Rates:   serviceableRates,
		Total:   len(serviceableRates),
	}
}
