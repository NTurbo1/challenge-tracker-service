package server

import (
	"net/http"
	"regexp"

	"github.com/nturbo1/challenge-tracker-service/log"
)

var authorizedPathsPattterns = []string{
	"^/user/\\w*",
}

type AuthHandler struct {
	next http.Handler
}

func (ah *AuthHandler) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	path := req.URL.Path
	log.Debug("req path: %s", path)
	if isAuthorizedPath(path) {
		log.Debug("Path %s is AUTHORIZED!")
	}
	log.Error("IMPLEMENT AUTH HANDLER!!!")

	ah.next.ServeHTTP(rw, req)
}

func isAuthorizedPath(path string) bool {
	for _, pattern := range authorizedPathsPattterns {
		matched, _ := regexp.MatchString(pattern, path) // Ignoring the error, guess it's fine???
		if matched {
			return true
		}
	}

	return false
}
