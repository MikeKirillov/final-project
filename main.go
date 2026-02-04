package main

import (
	"final-project/pkg/db"
	"final-project/pkg/server"
	"log"
	"net/http"
)

func main() {
	err := db.Init("scheduler.db")

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Starting server")

	webDir := "./web"
	http.Handle("/", http.FileServer(http.Dir(webDir)))

	server.Run()

	err = http.ListenAndServe(":7540", nil)
	if err != nil {
		panic(err)
	}

	log.Println("Stop working")
}
