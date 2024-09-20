package main

import "net/http"

type Application struct{}

func (a *Application) handleRequest(url, method string) (int, string) {
	if url == "/app/status" && method == "GET" {
		return http.StatusOK, "OK"
	}

	if url == "/create/user" && method == "POST" {
		return http.StatusCreated, "User created"
	}

	return http.StatusNotFound, "Not ok"
}
