package server

import (
	"net/http"

	"github.com/nturbo1/challenge-tracker-service/log"
)

type AuthHandler struct {
	next http.Handler
}

func (ah *AuthHandler) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	log.Error("IMPLEMENT AUTH HANDLER!!!")

	ah.next.ServeHTTP(rw, req)
}
