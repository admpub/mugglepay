package structs

import (
	"fmt"
)

// ResponseOrder 接口返回订单信息
type ResponseOrder struct {
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
	CreatedAt       string  `json:"created_at"`
	UpdatedAt       string  `json:"updated_at"`
	PayAmount       float64 `json:"pay_amount"`   // 我方原始金额(即我们自己提交过去的金额)
	PayCurrency     string  `json:"pay_currency"` // 我方原始金额币种
	IsSelf          bool    `json:"is_self"`
	Mobile          bool    `json:"mobile"`
	Fast            bool    `json:"fast"`
	Token           string  `json:"token"`
	PaidAt          string  `json:"paid_at"`
	ReceiveCurrency string  `json:"receive_currency"`
}

// ServerOrder 接口返回数据
type ServerOrder struct {
	Status     int           `json:"status"`
	Order      ResponseOrder `json:"order"`
	Merchant   Merchant      `json:"merchant"`
	PaymentURL string        `json:"payment_url"`
	Invoice    Invoice       `json:"invoice"`
	Permission string        `json:"permission"`

	// Failed:
	// MugglePay Server will always return status 400. If API failed, it will return error_code and error as its object.
	Error     string `json:"error"`
	ErrorCode string `json:"error_code"`
}

// ServerOrderParsers 响应订单数据解析器
var ServerOrderParsers = map[string]func(*ServerOrder){
	`EOS`: func(s *ServerOrder) {
		s.Invoice.Address = "mgtestflight"
		s.Invoice.Memo = fmt.Sprintf("MP:%s", s.Invoice.OrderID)
	},
}

// Parse 解析响应订单数据
func (s *ServerOrder) Parse() *ServerOrder {
	ps, ok := ServerOrderParsers[s.Invoice.PayCurrency]
	if ok {
		ps(s)
	}
	return s
}
