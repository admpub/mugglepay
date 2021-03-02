package structs

import (
	"fmt"

	"github.com/admpub/mugglepay/utils"
)

type ServerOrder struct {
	Status     int      `json:"status"`
	Order      Order    `json:"order"`
	Merchant   Merchant `json:"merchant"`
	PaymentURL string   `json:"payment_url"`
	Invoice    Invoice  `json:"invoice"`
	Permission string   `json:"permission"`
}

// GetPaymentURL 获取支付地址
func (s *ServerOrder) GetPaymentURL() {
	switch s.Invoice.PayCurrency {
	case "ALIPAY":
		if rurl := utils.URLQueryValueGetter(s.Invoice.Qrcode)("url"); rurl != "" {
			s.Invoice.Address = rurl
		} else {
			s.Invoice.Address = utils.URLQueryValueGetter(s.Invoice.QrcodeLg)("mpurl")
		}
	case "WECHAT":
		s.Invoice.Address = s.Invoice.Qrcode
	case "EOS":
		s.Invoice.Address = "mgtestflight"
		s.Invoice.Memo = fmt.Sprintf("MP:%s", s.Invoice.OrderID)
	}
}
