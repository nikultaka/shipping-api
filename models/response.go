package models

type APIResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

type CreateOrderResponse struct {
	OrderID      string `json:"order_id"`
	TrackingID   string `json:"tracking_id"`
	ExpectedDate string `json:"expected_date"`
	Carrier      string `json:"carrier"`
	Status       string `json:"status"`
}
