package server

import (
	"net/http"
)

var globalServeMux = http.NewServeMux()

func setupHandlers() {
	globalServeMux.HandleFunc("POST /login", handleLogin)
}
