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

// QRCodeURLParsers 二维码网址解析器
var QRCodeURLParsers = map[string]func(*ServerOrder){
	`ALIPAY`: func(s *ServerOrder) {
		if rurl := utils.URLQueryValueGetter(s.Invoice.Qrcode)("url"); len(rurl) > 0 {
			s.Invoice.Address = rurl
		} else {
			s.Invoice.Address = utils.URLQueryValueGetter(s.Invoice.QrcodeLg)("mpurl")
		}
	},
	`WECHAT`: func(s *ServerOrder) {
		s.Invoice.Address = s.Invoice.Qrcode
	},
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
