package server

import (
	"net/http"
)

type CorsHandler struct {
	next http.Handler
}

func (ch *CorsHandler) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	rw.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173") // TODO: Adjust this on production environment!
	rw.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
	rw.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	rw.Header().Set("Access-Control-Allow-Credentials", "true")

	if req.Method == http.MethodOptions {
		rw.WriteHeader(http.StatusNoContent)
		return
	}

	ch.next.ServeHTTP(rw, req)
}
