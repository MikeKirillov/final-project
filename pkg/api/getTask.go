package api

import (
	"database/sql"
	"errors"
	"log"
	"net/http"

	"final-project/pkg/db"
)

func getTaskHandler(w http.ResponseWriter, req *http.Request) {
	id := req.FormValue("id")

	if len(id) == 0 {
		err := errors.New("id must be specified")
		log.Println(err)
		writeJson(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}

	if task, err := db.GetTask(id); err != nil && err == sql.ErrNoRows {
		err := errors.New("task is not found")
		log.Println(err)
		writeJson(w, http.StatusNotFound, map[string]string{"error": err.Error()})
		return
	} else {
		writeJson(w, http.StatusOK, task)
	}
}
