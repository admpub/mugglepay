package structs

import (
	"crypto/md5"
	"fmt"
	"net/url"
	"strings"
)

type Order struct {
	MerchantOrderID string  `json:"merchant_order_id"` // required. 自己系统的订单号
	PriceAmount     float64 `json:"price_amount"`      // required. 金额
	PriceCurrency   string  `json:"price_currency"`    // required. 金额币种
	PayCurrency     string  `json:"pay_currency"`      // 支付方式，自己实现选择支付方式的时候提供。e.g. ALIPAY, ALIGLOBAL, WECHAT, BTC, LTC, ETH, EOS, BCH, LBTC (for Lightening BTC), CUSD (for Celo Dollars)
	Title           string  `json:"title"`
	Description     string  `json:"description"`
	CallbackURL     string  `json:"callback_url"`
	CancelURL       string  `json:"cancel_url"`
	SuccessURL      string  `json:"success_url"`
	Mobile          bool    `json:"mobile"`
	Fast            bool    `json:"fast"`
	Token           string  `json:"token"`
}

// Sign 签名
func (o *Order) Sign(secret string) {
	q := url.Values{}
	q.Set("merchant_order_id", o.MerchantOrderID)
	q.Set("secret", secret)
	q.Set("type", "FIAT")
	o.Token = strings.ToLower(fmt.Sprintf("%x", md5.Sum([]byte(fmt.Sprintf("%x", md5.Sum([]byte(q.Encode())))+secret))))
}
