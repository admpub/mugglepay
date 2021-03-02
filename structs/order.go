package structs

import (
	"crypto/md5"
	"fmt"
	"net/url"
	"strings"
)

type Order struct {
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

// Sign 签名
func (o *Order) Sign(secret string) {
	q := url.Values{}
	q.Set("merchant_order_id", o.MerchantOrderID)
	q.Set("secret", secret)
	q.Set("type", "FIAT")
	o.Token = strings.ToLower(fmt.Sprintf("%x", md5.Sum([]byte(fmt.Sprintf("%x", md5.Sum([]byte(q.Encode())))+secret))))
}
