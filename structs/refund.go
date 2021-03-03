package structs

type ResponseRefundOrder struct {
	OrderID         string  `json:"order_id"`          // muggalepay 订单ID
	UserID          int64   `json:"user_id"`           // muggalepay 付款用户ID
	MerchantOrderID string  `json:"merchant_order_id"` // 自己系统的订单ID
	Title           string  `json:"title"`
	Description     string  `json:"description"`
	CallbackURL     string  `json:"callback_url"`
	CancelURL       string  `json:"cancel_url"`
	SuccessURL      string  `json:"success_url"`
	PriceAmount     float64 `json:"price_amount"`   // muggalepay 实际交易金额
	PriceCurrency   string  `json:"price_currency"` // muggalepay 实际交易币种
	Status          string  `json:"status"`
	Notified        string  `json:"notified"`
	IsSelf          bool    `json:"is_self"`
	PaidAt          string  `json:"paid_at"`
	ReceiveCurrency string  `json:"receive_currency"`
	CreatedAt       string  `json:"created_at"`
	UpdatedAt       string  `json:"updated_at"`
}

type ServerRefund struct {
	Status int                 `json:"status"`
	Order  ResponseRefundOrder `json:"order"`
	// Failed:
	// MugglePay Server will always return status 400. If API failed, it will return error_code and error as its object.
	Error     string `json:"error"`
	ErrorCode string `json:"error_code"`
}
