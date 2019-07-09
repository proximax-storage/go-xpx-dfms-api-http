package httpclient

import (
	"net/http"
)

var (
	APIPath = "/api/v1"
)

type Client struct {
	url  string
	http http.Client

	Headers http.Header
}

func New(url string) (*Client, error) {
	return NewWithClient(url, http.DefaultClient)
}

func NewWithClient(url string, http *http.Client) (*Client, error) {
	client := &Client{
		url:     url,
		http:    *http,
		Headers: make(map[string][]string),
	}

	return client, nil
}

func (api *Client) NewRequest(command string) RequestBuilder {
	headers := make(map[string]string)
	if api.Headers != nil {
		for k := range api.Headers {
			headers[k] = api.Headers.Get(k)
		}
	}

	return &requestBuilder{
		command: command,
		client:  api,
		headers: headers,
	}
}
