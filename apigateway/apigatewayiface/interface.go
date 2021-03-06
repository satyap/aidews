package apigatewayiface

import (
	"net/http"
	"net/url"
)

// Service is an interface for making signed requests to API Gateway.
type Service interface {
	Do(*http.Request) (*http.Response, error)
	Delete(*http.Request) (*http.Response, error)
	Get(string, url.Values) (*http.Response, error)
	Put(string, interface{}) (*http.Response, error)
	Post(string, interface{}) (*http.Response, error)
	SetHeaders(map[string]string)
}
