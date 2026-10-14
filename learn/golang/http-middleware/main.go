package main

import (
	"fmt"
	"net/http"
	"slices"
	"time"
)

var (
	retryAbleMethod    = []string{http.MethodGet, http.MethodHead, http.MethodOptions, http.MethodPut, http.MethodDelete}
	nonRetryAbleMethod = []string{http.MethodPost, http.MethodPatch}
	idempotencyKey     = "abc-123"
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
	currentReq := req.Clone(req.Context())
	for key, value := range t.Headers {
		currentReq.Header.Set(key, value)
	}

	return next.RoundTrip(currentReq)
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

func isMethodNonRetryable(header string) bool {
	return slices.Contains(nonRetryAbleMethod, header)
}

func (t *RetryTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	next := t.Next
	if next == nil {
		next = http.DefaultTransport
	}
	backoffTime := 100 * time.Millisecond
	exponential := 2
	maxBackoffTime := 400 * time.Millisecond

	for attempt := 0; attempt <= t.MaxRetries; attempt++ {
		currentReq := req.Clone(req.Context())

		if attempt > 0 && req.Body != nil { // second try and the body is drained
			if req.GetBody == nil {
				return nil, fmt.Errorf("request body cannot be replayed")
			}
			body, err := req.GetBody()
			if err != nil {
				return nil, err
			}

			currentReq.Body = body
		}

		if attempt > 0 {
			fmt.Printf("Waiting %v before retry attempt %d...\n", backoffTime, attempt)
			timer := time.NewTimer(backoffTime)
			select {
			case <-timer.C:
			case <-currentReq.Context().Done():
				fmt.Println("context cancelled")
				timer.Stop()
				return nil, currentReq.Context().Err()
			}
			backoffTime = min(backoffTime*time.Duration(exponential), maxBackoffTime)
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

		idemKey := currentReq.Header.Get("Idempotency-Key")

		if isMethodNonRetryable(currentReq.Method) && idemKey != "" {
			return nil, fmt.Errorf("non retryable method header")
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

	req, _ := http.NewRequest("POST", "https://examdsfsdfple.com", nil)
	// req.Header.Add("Idempotency-Key", "abc-123")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer resp.Body.Close()

	fmt.Println(req.Header)

	fmt.Println(resp)
}
