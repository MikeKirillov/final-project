package server

import (
	"final-project/pkg/api"
	"log"
	"net/http"
)

func Run() {
	log.Println("Starting server")

	webDir := "./web"
	http.Handle("/", http.FileServer(http.Dir(webDir)))

	api.Init()

	err := http.ListenAndServe(":7540", nil)
	if err != nil {
		panic(err)
	}

	log.Println("Stop working")
}
