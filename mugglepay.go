package mugglepay

import (
	"errors"
	"fmt"

	"github.com/admpub/mugglepay/structs"
)

const (
	apiBaseURL     = "https://api.mugglepay.com/v1"
	createOrderURL = "/orders"
	getOrderURL    = "/orders/%s"
	checkoutURL    = "/orders/%s/checkout"
	statusURL      = "/orders/%s/status"
	sentURL        = "/orders/%s/sent"
)

func New(key string) *Mugglepay {
	m := &Mugglepay{
		AppKey: key,
		APIURL: apiBaseURL,
	}
	return m
}

type Mugglepay struct {
	AppKey      string
	APIURL      string
	CallbackURL string
	CancelURL   string
	SuccessURL  string
}

// CreateOrder 创建订单，返回 ServerOrder
func (m *Mugglepay) CreateOrder(order *structs.Order) (structs.ServerOrder, error) {
	var sorder structs.ServerOrder
	if m.AppKey == "" {
		return sorder, errors.New("application key cannot be null")
	}
	if order.MerchantOrderID == "" {
		return sorder, errors.New("merchant order id cannot be null")
	}

	if order.PriceCurrency == "" {
		order.PriceCurrency = "CNY"
	}
	if m.CallbackURL == "" {
		// 如果没有回调地址将无法使用法币支付，默认仅可用虚拟币
		order.PayCurrency = ""
	}

	order.CallbackURL = m.CallbackURL
	order.CancelURL = m.CancelURL
	order.SuccessURL = m.SuccessURL

	// 签名
	order.Sign(m.AppKey)

	err := m.Post(createOrderURL, &sorder, order)
	return sorder, err
}

// VerifyOrder 校验订单 true: 已支付; false: 未支付/取消/欺诈
func (m *Mugglepay) VerifyOrder(callback *structs.Callback) bool {
	if m.AppKey == "" {
		return false
	}
	order := &structs.Order{
		MerchantOrderID: callback.MerchantOrderID,
	}
	order.Sign(m.AppKey)
	// 校验签名
	if order.Token != callback.Token {
		return false
	}
	if callback.Status == "PAID" {
		return true
	}
	return false
}

// GetOrder 根据网关订单编号获取 ServerOrder
func (m *Mugglepay) GetOrder(orderID string) (structs.ServerOrder, error) {
	var sorder structs.ServerOrder
	if orderID == "" {
		return sorder, errors.New("order id cannot be null")
	}
	err := m.Get(fmt.Sprintf(getOrderURL, orderID), &sorder)
	return sorder, err
}

// CheckOut 切换网关支付方式
func (m *Mugglepay) CheckOut(orderID, PayCurrency string) (structs.ServerOrder, error) {
	var sorder structs.ServerOrder
	if orderID == "" {
		return sorder, errors.New("order id cannot be null")
	}
	me := make(map[string]string)
	me["pay_currency"] = PayCurrency
	err := m.Post(fmt.Sprintf(checkoutURL, orderID), &sorder, me)
	return sorder, err
}

// GetStatus 订单查询
func (m *Mugglepay) GetStatus(orderID string) (structs.ServerOrder, error) {
	var sorder structs.ServerOrder
	if orderID == "" {
		return sorder, errors.New("order id cannot be null")
	}
	err := m.Get(fmt.Sprintf(statusURL, orderID), &sorder)
	return sorder, err
}

var emptyBody = make(map[string]interface{})

// Sent 虚拟币: 我已支付
func (m *Mugglepay) Sent(orderID string) (structs.ServerOrder, error) {
	var sorder structs.ServerOrder
	var err error
	if orderID == "" {
		return sorder, errors.New("order id cannot be null")
	}
	sorder, err = m.GetOrder(orderID)
	if err != nil {
		return sorder, err
	}
	if sorder.Invoice.PayCurrency == "ALIPAY" || sorder.Invoice.PayCurrency == "WECHAT" {
		// 法币不可调用此 API
		return sorder, errors.New("tan 90°")
	}
	err = m.Post(fmt.Sprintf(sentURL, orderID), &sorder, emptyBody)
	return sorder, err
}
