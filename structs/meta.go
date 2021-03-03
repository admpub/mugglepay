package structs

// Meta 支付宝、微信等法币平台交易信息
type Meta struct {
	Payment     string `json:"payment"`
	TotalMmount string `json:"total_amount"`
	TradeNo     string `json:"trade_no"`
	OutTradeNo  string `json:"out_trade_no"`
}
