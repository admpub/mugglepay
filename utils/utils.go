package utils

import "net/url"

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
