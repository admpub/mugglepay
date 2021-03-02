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

func Request(contentType string, key string) *resty.Request {
	r := Resty().R()
	if contentType == `json` {
		r.SetHeader("Content-Type", "application/json")
	}
	r.SetHeader("Accept", "application/json")
	r.SetHeader("token", key)
	return r
}

func HTTPGet(fullURL string, result interface{}, key string) (*resty.Response, error) {
	client := Request(``, key)
	client.SetResult(result)
	return client.Get(fullURL)
}

func HTTPPost(fullURL string, result interface{}, body interface{}, key string) (*resty.Response, error) {
	client := Request(`json`, key)
	client.SetBody(body)
	client.SetResult(result)
	return client.Post(fullURL)
}
