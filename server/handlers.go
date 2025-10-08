package server

import (
	"net/http"
	"strings"

	"github.com/nturbo1/challenge-tracker-service/db"
)

func handleLogin(rw http.ResponseWriter, req *http.Request) {
	rw.Header().Set("Access-Control-Allow-Origin", "*")

	if strings.Compare(req.Method, "POST") != 0 {
		rw.WriteHeader(400)
		return
	}
	if db.SessionRepository == nil {
		rw.WriteHeader(500)
		return
	}
}
