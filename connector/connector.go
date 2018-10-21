package connector

import (
	"errors"
	"net/http"
)

type ApiAccessPoint struct {
	Key    string `json:"Key"`
	Secret string `json:"Secret"`
}

type HttpEngine struct {
	credentials *ApiAccessPoint
	HttpClient  *http.Client
}

func New(credentials *ApiAccessPoint) *HttpEngine {
	engine := new(HttpEngine)
	engine.credentials = credentials

	return engine
}

func (e *HttpEngine) SendRequest(request *http.Request) (*http.Response, error) {
	return nil, errors.New("SendRequest() not implemented")
}
