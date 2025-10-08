package main

import (
	"github.com/nturbo1/challenge-tracker-service/db"
	"github.com/nturbo1/challenge-tracker-service/server"
)

func main() {
	err := db.InitDb()
	defer db.Flush()
	defer db.Close()

	if err != nil {
		panic(err)
	}

	err = server.Start()
	if err != nil {
		panic(err)
	}
	defer server.Close()
}
