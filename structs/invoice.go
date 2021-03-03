package structs

// Invoice 客户在mugglepay上的金额信息
type Invoice struct {
	InvoiceID       string  `json:"invoice_id"`
	OrderID         string  `json:"order_id"`
	PayAmount       float64 `json:"pay_amount"`   // 客户支付的金额
	PayCurrency     string  `json:"pay_currency"` // 客户支付的币种
	Status          string  `json:"status"`
	CreatedAt       string  `json:"created_at"`
	CreatedAtT      int64   `json:"created_at_t"`
	ExpiredAt       string  `json:"expired_at"`
	ExpiredAtT      int64   `json:"expired_at_t"`
	MerchantOrderID string  `json:"merchant_order_id"`
	ReceiveAmount   float64 `json:"receive_amount"`   // 实际收款金额
	ReceiveCurrency string  `json:"receive_currency"` // 实际收款币种
	Qrcode          string  `json:"qrcode"`
	QrcodeLg        string  `json:"qrcodeLg"`
	Address         string  `json:"address"`
	Memo            string  `json:"memo"`
}
