package main

import (
	"log"
	"os"

	"final-project/pkg/db"
	"final-project/pkg/server"
)

const TODO_PORT = "TODO_PORT"

func main() {
	// allow to run like: TODO_PORT=<port_number> go run .
	port := os.Getenv(TODO_PORT)

	log.Println("port is: " + port)

	err := db.Init("scheduler.db")

	if err != nil {
		log.Fatal(err)
	}

	server.Run()
	db.Close()
}
