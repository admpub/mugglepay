package structs

type Meta struct {
	Payment     string `json:"payment"`
	TotalMmount string `json:"total_amount"`
	TradeNo     string `json:"trade_no"`
	OutTradeNo  string `json:"out_trade_no"`
}
