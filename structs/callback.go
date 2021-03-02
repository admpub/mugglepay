package structs

type Callback struct {
	MerchantOrderID string  `json:"merchant_order_id"`
	OrderID         string  `json:"order_id"`
	Status          string  `json:"status"`
	PriceAmount     float64 `json:"price_amount"`
	PriceCurrency   string  `json:"price_currency"`
	PayAmount       float64 `json:"pay_amount"`
	PayCurrency     string  `json:"pay_currency"`
	CreatedAt       string  `json:"created_at"`
	CreatedAtT      int64   `json:"created_at_t"`
	Token           string  `json:"token"`
	Meta            Meta    `json:"meta"`
}
