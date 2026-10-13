package main

import (
	"fmt"
	"net/http"
	"time"
)

const (
	DefaultRetries = 5
	DefaultTimeout = 10 * time.Minute
)

type Client struct {
	URL        string
	Headers    map[string]string
	Retries    int
	Debug      bool
	httpClient *http.Client
}

type Option func(*Client)

func NewClient(baseURL string, options ...Option) *Client {
	client := &Client{
		httpClient: &http.Client{
			Timeout: DefaultTimeout,
		},
		URL:     baseURL,
		Retries: DefaultRetries,
		Headers: map[string]string{},
	}
	for _, option := range options {
		option(client)
	}
	return client
}

func WithHeader(key, value string) Option {
	return func(c *Client) {
		if key != "" && value != "" {
			c.Headers[key] = value
		}
	}
}

func WithTimeout(timeout time.Duration) Option {
	return func(c *Client) {
		if timeout > 0 {
			c.httpClient.Timeout = timeout
		}
	}
}

func WithRetries(retries int) Option {
	return func(c *Client) {
		if retries > 0 {
			c.Retries = retries
		}
	}
}

func WithDebug() Option {
	return func(c *Client) {
		c.Debug = true
	}
}

func main() {
	c1 := NewClient("https://api.test.com")

	c2 := NewClient(
		"https://api.test.com",
		WithTimeout(10*time.Minute),
		WithDebug(),
		WithRetries(5),
		WithHeader("Authorization", "Bearer abc"),
		WithHeader("X-Market", "VN"),
	)

	fmt.Printf("%+v\n", c1)
	fmt.Printf("%+v\n", c2)

}
