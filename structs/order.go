package structs

import (
	"crypto/md5"
	"fmt"
	"net/url"
	"strings"
)

type Order struct {
	MerchantOrderID string  `json:"merchant_order_id"`
	PriceAmount     float64 `json:"price_amount"`
	PriceCurrency   string  `json:"price_currency"`
	PayCurrency     string  `json:"pay_currency"`
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
