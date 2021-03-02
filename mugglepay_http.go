package mugglepay

import "github.com/admpub/mugglepay/utils"

// HTTPPost HTTP POST
func (m *Mugglepay) HTTPPost(apiURL string, result interface{}, body interface{}) error {
	_, err := utils.HTTPPost(m.APIURL+apiURL, result, body, m.AppKey)
	return err
}

// HTTPGet HTTP GET
func (m *Mugglepay) HTTPGet(apiURL string, result interface{}) error {
	_, err := utils.HTTPGet(m.APIURL+apiURL, result, m.AppKey)
	return err
}
