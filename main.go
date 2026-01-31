package main

import (
	"final-project/pkg/db"
	"final-project/pkg/server"
	"log"
)

func main() {
	err := db.Init("scheduler.db")

	if err != nil {
		log.Fatal(err)
	}

	server.Start()
}
