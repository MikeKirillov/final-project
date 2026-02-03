package api

import (
	"errors"
	"log"
	"net/http"
)

func doneTaskHandler(w http.ResponseWriter, req *http.Request) {
	id := req.FormValue("id")
	if len(id) == 0 {
		err := errors.New("id must be specified")
		log.Println(err)
		writeJson(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}

}
