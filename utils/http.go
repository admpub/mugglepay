package utils

import (
	"sync"
	"time"

	"github.com/admpub/resty/v2"
)

var (
	DefaultTimeout = 10 * time.Second
	restyG         *resty.Client
	once           sync.Once
)

func initResty() {
	restyG = resty.New().SetTimeout(DefaultTimeout)
}

func Resty() *resty.Client {
	once.Do(initResty)
	return restyG
}

func Request(key string, contentType ...string) *resty.Request {
	r := Resty().R()
	if len(contentType) > 0 {
		switch contentType[0] {
		case `json`:
			r.SetHeader("Content-Type", "application/json")
		}
	}
	r.SetHeader("Accept", "application/json")
	r.SetHeader("token", key)
	return r
}

type ReqURL string

func (r ReqURL) Get(result interface{}, key string) (*resty.Response, error) {
	client := Request(key)
	client.SetResult(result)
	return client.Get(r.String())
}

func (r ReqURL) String() string {
	return string(r)
}

func (r ReqURL) Post(result interface{}, body interface{}, key string) (*resty.Response, error) {
	client := Request(key, `json`)
	client.SetBody(body)
	client.SetResult(result)
	return client.Post(r.String())
}
