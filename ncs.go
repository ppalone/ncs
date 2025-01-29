package ncs

import (
	"net/http"
)

// constants
const (
	baseURL string = "https://ncs.io"
)

// NCS Client
type Client struct {
	httpClient *http.Client
}

// NewClient returns a new NCS client
func NewClient(httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = &http.Client{}
	}

	return &Client{httpClient}
}
