package structs

import (
	"fmt"
)

type ServerOrder struct {
	Status     int      `json:"status"`
	Order      Order    `json:"order"`
	Merchant   Merchant `json:"merchant"`
	PaymentURL string   `json:"payment_url"`
	Invoice    Invoice  `json:"invoice"`
	Permission string   `json:"permission"`

	// Failed:
	// MugglePay Server will always return status 400. If API failed, it will return error_code and error as its object.
	Error     string `json:"error"`
	ErrorCode string `json:"error_code"`
}

// QRCodeURLParsers 二维码网址解析器
var QRCodeURLParsers = map[string]func(*ServerOrder){
	`EOS`: func(s *ServerOrder) {
		s.Invoice.Address = "mgtestflight"
		s.Invoice.Memo = fmt.Sprintf("MP:%s", s.Invoice.OrderID)
	},
}

// ParseInvoiceAddress 解析扫码地址
func (s *ServerOrder) ParseInvoiceAddress() *ServerOrder {
	ps, ok := QRCodeURLParsers[s.Invoice.PayCurrency]
	if ok {
		ps(s)
	}
	return s
}
