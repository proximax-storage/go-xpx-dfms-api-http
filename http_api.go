package apihttp

import (
	"net/http"

	api "github.com/proximax-storage/go-xpx-dfms-api"
)

var (
	APIPath = "/api/v1"
)

type apiHttp struct {
	address string
	token   api.AccessToken
	http    http.Client

	Headers http.Header
}

func newHTTP(address string, token api.AccessToken, http *http.Client) *apiHttp {
	return &apiHttp{
		address: address,
		token:   token,
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

	rb := &requestBuilder{
		command: command,
		client:  c,
		headers: headers,
	}

	if c.token != nil {
		return rb.Option("token", string(c.token))
	}

	return rb
}
