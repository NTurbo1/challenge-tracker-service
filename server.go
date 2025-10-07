package main

import (
	"fmt"
	"net/http"
	"time"
	"strconv"
	"strings"
	"github.com/nturbo1/challenge-tracker-service/db"
)

const serverPort int = 8080

var Server = &http.Server{
    Addr:           ":" + strconv.Itoa(serverPort),
    Handler:        nil,
    ReadTimeout:    10 * time.Second,
    WriteTimeout:   10 * time.Second,
    MaxHeaderBytes: 1 << 20,
}

func setResponseHeaders(rw http.ResponseWriter) {
	rw.Header().Set("Access-Control-Allow-Origin", "*")
    rw.Header().Set("Content-Type", "application/json")
}

func SetupHandlers() {
	http.HandleFunc("/login", func(rw http.ResponseWriter, req *http.Request) {
		rw.Header().Set("Access-Control-Allow-Origin", "*")

		if strings.Compare(req.Method, "POST") != 0 {
			rw.WriteHeader(400)
			return
		}
		if db.SessionRepository == nil {
			rw.WriteHeader(500)
			return
		}

		cookie := &http.Cookie{
			Name: "session_id",
			Value: "a;sldkfja;lskdfja;sldfk",
			MaxAge: 5342452323,
			HttpOnly: true,
			Secure: true,
		}
		fmt.Println("Just to check cookie string: ", cookie)
	})
}
