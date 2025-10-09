package server

import (
	"net/http"

	"github.com/nturbo1/challenge-tracker-service/log"
)

type ServerHandler struct {}
func (sh *ServerHandler) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	log.HttpRequest(req)
	// TODO: Configure the CORS properly later!
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	rw.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	rw.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if req.Method == http.MethodOptions {
		rw.WriteHeader(http.StatusNoContent)
	}
}
