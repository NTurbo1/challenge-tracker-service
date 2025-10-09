package server

import (
	"fmt"
	"net/http"
	"time"
	"strconv"
)

const serverPort int = 8080

var server = &http.Server{
    Addr:           ":" + strconv.Itoa(serverPort),
    Handler:        &ServerHandler{},
    ReadTimeout:    10 * time.Second,
    WriteTimeout:   10 * time.Second,
    MaxHeaderBytes: 1 << 20,
}

func Start() error {
	fmt.Println("Starting the server...")
    setupHandlers()
	fmt.Println("HTTP Handlers are set!")
    fmt.Println("Server listening on port", serverPort)

	return server.ListenAndServe()
}

func Close() {
	fmt.Println("Closing the server...")
	err := server.Close()
	if err != nil {
		panic(err)
	}
}

func setResponseHeaders(rw http.ResponseWriter) {
    rw.Header().Set("Content-Type", "application/json")
}

func setupHandlers() {
	http.HandleFunc("/login", handleLogin)
}
