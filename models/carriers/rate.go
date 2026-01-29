package carriers

type RateResponse struct {
	Carrier       string  `json:"carrier"`
	Price         float64 `json:"price"`
	ETD           string  `json:"estimated_delivery"`
	ServiceType   string  `json:"service_type"`
	IsError       bool    `json:"is_error,omitempty"`
	ErrorMessage  string  `json:"error_message,omitempty"`
	IsServiceable bool    `json:"is_serviceable"`
}
