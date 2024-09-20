package main

import "net/http"

type Nginx struct {
	application        *Application
	maxAllowdedRequest int
	rateLimiter        map[string]int
}

func newNginxServer() *Nginx {
	return &Nginx{
		application:        &Application{},
		maxAllowdedRequest: 2,
		rateLimiter:        map[string]int{},
	}
}

func (n *Nginx) handleRequest(url, method string) (int, string) {
	if allowed := n.checkRateLimit(url); !allowed {
		return http.StatusForbidden, "Not allowed"
	}

	return n.application.handleRequest(url, method)
}

func (n *Nginx) checkRateLimit(url string) bool {
	if n.rateLimiter[url] == 0 {
		n.rateLimiter[url] = 1
	}

	if n.rateLimiter[url] > n.maxAllowdedRequest {
		return false
	}

	n.rateLimiter[url] = n.rateLimiter[url] + 1
	return true
}
