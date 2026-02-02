package server

import (
	"log"
	"net/http"
)

func Start() {
	log.Println("Starting server")

	webDir := "./web"
	http.Handle("/", http.FileServer(http.Dir(webDir)))

	err := http.ListenAndServe(":7540", nil)
	if err != nil {
		panic(err)
	}

	log.Println("Stop working")
}
