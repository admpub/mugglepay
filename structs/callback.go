package structs

// Callback 回调数据
type Callback struct {
	MerchantOrderID string  `json:"merchant_order_id"`
	OrderID         string  `json:"order_id"`
	Status          string  `json:"status"`         // 状态
	PriceAmount     float64 `json:"price_amount"`   // muggalepay 实际交易金额
	PriceCurrency   string  `json:"price_currency"` // muggalepay 实际交易币种
	PayAmount       float64 `json:"pay_amount"`     // 我方原始金额(即我们自己提交过去的金额)
	PayCurrency     string  `json:"pay_currency"`   // 我方原始金额币种
	CreatedAt       string  `json:"created_at"`
	CreatedAtT      int64   `json:"created_at_t"`
	Token           string  `json:"token"`
	Meta            Meta    `json:"meta"`
}
