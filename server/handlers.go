package server

import (
	"fmt"
	"net/http"
	"strings"
)

func handleLogin(rw http.ResponseWriter, req *http.Request) {
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	rw.Header().Set("Access-Control-Allow-Headers", "*")

	if strings.Compare(req.Method, "POST") != 0 {
		rw.WriteHeader(400)
		return
	}

	if req.ContentLength <= 0 {
		fmt.Println("Request body content length is ", req.ContentLength)
		fmt.Println(
			"Can't deserialize or parse the content if I don't know the length :( ... Or can I? o.O",
		)
		rw.WriteHeader(400)
		return
	}
}
