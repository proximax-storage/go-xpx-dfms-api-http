package apihttp

import (
	"net/http"
)

var (
	APIPath = "/api/v1"
)

type apiHttp struct {
	address string
	http    http.Client

	Headers http.Header
}

func newHTTP(address string, http *http.Client) *apiHttp {
	return &apiHttp{
		address: address,
		http:    *http,
		Headers: make(map[string][]string),
	}
}

// NewRequest constructs new builder for requests
func (c *apiHttp) NewRequest(command string) RequestBuilder {
	headers := make(map[string]string)
	if c.Headers != nil {
		for k := range c.Headers {
			headers[k] = c.Headers.Get(k)
		}
	}

	return &requestBuilder{
		command: command,
		client:  c,
		headers: headers,
	}
}
