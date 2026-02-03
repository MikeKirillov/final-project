package api

import (
	"errors"
	"final-project/pkg/db"
	"log"
	"net/http"
)

func deleteTaskHandler(w http.ResponseWriter, req *http.Request) {
	id := req.FormValue("id")
	if len(id) == 0 {
		err := errors.New("id must be specified")
		log.Println(err)
		writeJson(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}

	err := db.DeleteTask(id)
	if err != nil {
		log.Println(err)
		writeJson(w, http.StatusNotFound, map[string]string{"error": err.Error()})
	}

	writeJson(w, http.StatusOK, db.Task{})
}
