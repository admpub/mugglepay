package structs

import (
	"fmt"
)

type ResponseOrder struct {
	OrderID         string  `json:"order_id"`
	UserID          int64   `json:"user_id"`
	MerchantOrderID string  `json:"merchant_order_id"`
	Title           string  `json:"title"`
	Description     string  `json:"description"`
	CallbackURL     string  `json:"callback_url"`
	CancelURL       string  `json:"cancel_url"`
	SuccessURL      string  `json:"success_url"`
	PriceAmount     float64 `json:"price_amount"`
	PriceCurrency   string  `json:"price_currency"`
	Status          string  `json:"status"`
	Notified        string  `json:"notified"`
	CreatedAt       string  `json:"created_at"`
	UpdatedAt       string  `json:"updated_at"`
	PayAmount       float64 `json:"pay_amount"`
	PayCurrency     string  `json:"pay_currency"`
	IsSelf          bool    `json:"is_self"`
	Mobile          bool    `json:"mobile"`
	Fast            bool    `json:"fast"`
	Token           string  `json:"token"`
	PaidAt          string  `json:"paid_at"`
	ReceiveCurrency string  `json:"receive_currency"`
}

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
