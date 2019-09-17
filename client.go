package client

import (
	"net/http"
)

var (
	APIPath = "/api/v1"
)

// Client is an entry point for communication with DFMS applications
type Client struct {
	address string
	http    http.Client

	Headers http.Header
}

// New creates new Client from given address and default http.Client
func New(address string) *Client {
	return NewWithClient(address, http.DefaultClient)
}

// New creates new Client from address and custom http.Client
func NewWithClient(address string, http *http.Client) *Client {
	return &Client{
		address: address,
		http:    *http,
		Headers: make(map[string][]string),
	}
}

// NewRequest constructs new builder for requests
func (c *Client) NewRequest(command string) RequestBuilder {
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

func (c *Client) DriveAPI() *DriveAPI {
	return (*DriveAPI)(c)
}

func (c *Client) ContractAPI() *ContractAPI {
	return (*ContractAPI)(c)
}
