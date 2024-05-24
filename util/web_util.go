package util

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type WebUtil struct {
	Client *http.Client
	Config *viper.Viper
}

func NewWebUtil(client *http.Client, config *viper.Viper) *WebUtil {
	return &WebUtil{Client: client, Config: config}
}

func (m *WebUtil) Get(service string, path string, urlValues url.Values) (int, []byte, error) {
	_, code, bytes, err := m.Do("GET", "application/json;charset=UTF-8", service, path, urlValues, nil)
	return code, bytes, err
}

func (m *WebUtil) Post(service string, path string, bodyValue interface{}) (int, []byte, error) {
	_, code, bytes, err := m.Do("POST", "application/json;charset=UTF-8", service, path, nil, bodyValue)
	return code, bytes, err
}

func (m *WebUtil) Put(service string, path string, urlValues url.Values, bodyValue interface{}) (int, []byte, error) {
	_, code, bytes, err := m.Do("PUT", "application/json;charset=UTF-8", service, path, urlValues, bodyValue)
	return code, bytes, err
}

func (m *WebUtil) Delete(service string, path string, urlValues url.Values, bodyValue interface{}) (int, []byte, error) {
	_, code, bytes, err := m.Do("DELETE", "application/json;charset=UTF-8", service, path, urlValues, bodyValue)
	return code, bytes, err
}

func (m *WebUtil) Do(method string, contentType string, service string, path string, urlValues url.Values, bodyValue interface{}) (*http.Response, int, []byte, error) {
	host, err := m.serviceEntry(service)

	if err != nil {
		return nil, 0, nil, err
	}

	bodyValueBytes, err1 := json.Marshal(bodyValue)

	if err1 != nil {
		return nil, 0, nil, err1
	}

	if req, err := http.NewRequest(method, m.url(host, path, urlValues), bytes.NewReader(bodyValueBytes)); err != nil {
		return nil, 0, nil, err
	} else {
		return m.doRequest(req, contentType)
	}
}

func (m *WebUtil) serviceEntry(service string) (string, error) {
	return m.Config.GetString(fmt.Sprintf("services.%s.url", service)), nil
}

func (m *WebUtil) url(baseUrl string, path string, urlValues url.Values) string {
	result, _ := url.JoinPath(baseUrl, path)

	sb := new(strings.Builder)

	sb.WriteString(result)

	if len(urlValues) > 0 {
		sb.WriteString("?")
		sb.WriteString(urlValues.Encode())
	}

	return sb.String()
}

func (m *WebUtil) doRequest(req *http.Request, contentType string) (*http.Response, int, []byte, error) {
	req.Header.Set("Content-Type", contentType)

	if res, err := m.Client.Do(req); err != nil {
		return res, res.StatusCode, nil, err
	} else {
		if resBytes, err := io.ReadAll(res.Body); err != nil {
			return res, res.StatusCode, nil, err
		} else {
			return res, res.StatusCode, resBytes, nil
		}
	}
}
