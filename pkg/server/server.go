package server

import (
	"log"
	"net/http"

	"final-project/pkg/api"
)

const port = ":7540"

func Run() {
	log.Printf("Starting server on port %v\n", port)

	webDir := "./web"
	http.Handle("/", http.FileServer(http.Dir(webDir)))

	api.Init()

	err := http.ListenAndServe(port, nil)
	if err != nil {
		panic(err)
	}
}
