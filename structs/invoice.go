package structs

type Invoice struct {
	InvoiceID       string  `json:"invoice_id"`
	OrderID         string  `json:"order_id"`
	PayAmount       float64 `json:"pay_amount"`
	PayCurrency     string  `json:"pay_currency"`
	Status          string  `json:"status"`
	CreatedAt       string  `json:"created_at"`
	CreatedAtT      int64   `json:"created_at_t"`
	ExpiredAt       string  `json:"expired_at"`
	ExpiredAtT      int64   `json:"expired_at_t"`
	MerchantOrderID string  `json:"merchant_order_id"`
	ReceiveAmount   float64 `json:"receive_amount"`
	ReceiveCurrency string  `json:"receive_currency"`
	Qrcode          string  `json:"qrcode"`
	QrcodeLg        string  `json:"qrcodeLg"`
	Address         string  `json:"address"`
	Memo            string  `json:"memo"`
}
