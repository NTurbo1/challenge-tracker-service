package main

import (
    "fmt"
	"github.com/nturbo1/challenge-tracker-service/db"
)

func main() {
	err := db.InitDb()
	defer db.Flush()
	defer db.Close()

	if err != nil {
		panic(err)
	}
    SetupHandlers()
    fmt.Println("Server listening on port", serverPort)
    Server.ListenAndServe()
}
