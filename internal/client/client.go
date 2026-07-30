package client

import (
	"fmt"
	"net/http"
	"strings"
	"time"
)

type Client struct {
	baseURL    string
	httpClient *http.Client
}

func New(remoteAddress string) *Client {
	addr := strings.TrimSpace(remoteAddress)
	if addr == "" {
		addr = "localhost:8080"
	}

	if !strings.HasPrefix(addr, "http://") && !strings.HasPrefix(addr, "https://") {
		addr = "http://" + addr
	}

	addr = strings.TrimRight(addr, "/")

	return &Client{
		baseURL: addr,
		httpClient: &http.Client{
			Timeout: 5 * time.Minute,
		},
	}
}

func (c *Client) url(path string) string {
	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}
	return c.baseURL + path
}

func (c *Client) BaseURL() string {
	return c.baseURL
}

func errStatus(resp *http.Response, context string) error {
	return fmt.Errorf("%s: server returned %s", context, resp.Status)
}
