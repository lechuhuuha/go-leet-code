package main

import (
	"fmt"
	"net/http"
	"time"
)

type HeaderTransport struct {
	Next    http.RoundTripper
	Headers map[string]string
}

func (t *HeaderTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	next := t.Next
	if next == nil {
		next = http.DefaultTransport
	}

	for key, value := range t.Headers {
		req.Header.Set(key, value)
	}

	return next.RoundTrip(req)
}

type LoggingTransport struct {
	Next http.RoundTripper
}

func (t *LoggingTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	next := t.Next
	if next == nil {
		next = http.DefaultTransport
	}
	start := time.Now()
	defer func() {
		fmt.Printf("%s %s took %dms\n", req.Method, req.URL, time.Since(start).Milliseconds())
	}()
	return next.RoundTrip(req)
}

type RetryTransport struct {
	Next       http.RoundTripper
	MaxRetries int
}

func (t *RetryTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	next := t.Next
	if next == nil {
		next = http.DefaultTransport
	}
	for attempt := 0; attempt <= t.MaxRetries; attempt++ {
		currentReq := req

		if attempt > 0 && req.Body != nil { // second try and the body is drained
			if req.GetBody == nil {
				return nil, fmt.Errorf("request body cannot be replayed")
			}
			body, err := req.GetBody()
			if err != nil {
				return nil, err
			}

			currentReq = req.Clone(req.Context())
			currentReq.Body = body
		}

		resp, err := next.RoundTrip(currentReq)
		if err != nil {
			if attempt == t.MaxRetries {
				return nil, err
			}

			continue
		}

		if resp.StatusCode < 500 {
			return resp, nil
		}

		if attempt == t.MaxRetries {
			return resp, nil
		}

		resp.Body.Close()
	}
	panic("unreachable")
}

func main() {
	client := &http.Client{
		Transport: &RetryTransport{
			MaxRetries: 3,
			Next: &LoggingTransport{
				Next: &HeaderTransport{
					Next: http.DefaultTransport,
					Headers: map[string]string{
						"X-Market": "VN",
						"X-App":    "learning-client",
					},
				},
			},
		},
	}

	req, _ := http.NewRequest("GET", "https://examdsfsdfple.com", nil)

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()
	fmt.Println(req.Header)

	fmt.Println(resp)
}
