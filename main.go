package main

import (
    "fmt"
	"github.com/nturbo1/challenge-tracker-service/db"
)

func main() {
	err := db.InitDb()
	if err != nil {
		panic(err)
	}
    SetupHandlers()
    fmt.Println("Server listening on port", serverPort)
    Server.ListenAndServe()
}
