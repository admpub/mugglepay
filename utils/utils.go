package utils

import (
	"net/url"

	"github.com/webx-top/com"
)

func URLQueryValueGetter(longURL string) func(string) string {
	if u, err := url.Parse(longURL); err == nil {
		if p, err := url.ParseQuery(u.RawQuery); err == nil {
			return func(key string) string {
				var res string
				if val, ok := p[key]; ok {
					res = val[0]
				}
				return res
			}
		}
	}
	return func(_ string) string {
		return ``
	}
}

// LegalCurrencies 法币支付
var LegalCurrencies = []string{`ALIPAY`, `WECHAT`}

// IsCryptoCurrency 是否是加密货币支付
func IsCryptoCurrency(payCurrency string) bool {
	return !com.InSlices(payCurrency, LegalCurrencies)
}
