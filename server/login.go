package server

import (
	"fmt"
	"net/http"

	"github.com/nturbo1/challenge-tracker-service/log"
)

func handleLogin(rw http.ResponseWriter, req *http.Request) {
	log.HttpRequest(req)
	log.Error("IMPLEMENT HANDLE LOGIN!!!")
	if req.ContentLength <= 0 {
		fmt.Println("Request body content length is ", req.ContentLength)
		fmt.Println(
			"Can't deserialize or parse the content if I don't know the length :( ..." +
			"Or can I? o.O",
		)
		rw.WriteHeader(400)
		return
	}
}
