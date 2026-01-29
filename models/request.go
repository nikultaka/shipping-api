package models

type CreateOrderRequest struct {
	OrderID     string  `json:"order_id" binding:"required"`
	Carrier     string  `json:"carrier" binding:"required"`
	FromAddress string  `json:"from_address" binding:"required"`
	ToAddress   string  `json:"to_address" binding:"required"`
	Weight      float64 `json:"weight" binding:"required"`
}

type RateRequest struct {
	FromPincode string  `json:"from_pincode" binding:"required"`
	ToPincode   string  `json:"to_pincode" binding:"required"`
	Weight      float64 `json:"weight" binding:"required"`
}
