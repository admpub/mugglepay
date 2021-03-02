package mugglepay

import "github.com/admpub/mugglepay/utils"

// ReqURL HTTP Request URL
func (m *Mugglepay) ReqURL(apiURL string) utils.ReqURL {
	return utils.ReqURL(m.APIURL + apiURL)
}

// Post HTTP POST
func (m *Mugglepay) Post(apiURL string, result interface{}, body interface{}) error {
	_, err := m.ReqURL(apiURL).Post(result, body, m.AppKey)
	return err
}

// Get HTTP GET
func (m *Mugglepay) Get(apiURL string, result interface{}) error {
	_, err := m.ReqURL(apiURL).Get(result, m.AppKey)
	return err
}
