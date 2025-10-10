package server

import (
	"net/http"
	"time"
	"strconv"

	"github.com/nturbo1/challenge-tracker-service/log"
)

const serverPort int = 8080

var server = &http.Server{
    Addr:           ":" + strconv.Itoa(serverPort),
    Handler:        nil,
    ReadTimeout:    10 * time.Second,
    WriteTimeout:   10 * time.Second,
    MaxHeaderBytes: 1 << 20,
}

func Start() error {
	log.Info("Starting the server...")
    setupHandlers()
	server.Handler = &CorsHandler{&AuthHandler{globalServeMux}}
	log.Info("HTTP Handlers are set!")
    log.Info("Server listening on port %d", serverPort)

	return server.ListenAndServe()
}

func Close() {
	log.Info("Closing the server...")
	err := server.Close()
	if err != nil {
		panic(err)
	}
}
