package api

import (
	"database/sql"
	"errors"
	"final-project/pkg/db"
	"log"
	"net/http"
)

func getTaskHandler(w http.ResponseWriter, req *http.Request) {
	id := req.FormValue("id")

	if len(id) == 0 {
		err := errors.New("id must be specified")
		log.Println(err)
		writeJson(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	task, err := db.GetTask(id)
	if err != nil && err == sql.ErrNoRows {
		log.Println(err)
		writeJson(w, http.StatusNotFound, map[string]string{"error": err.Error()})
	}

	writeJson(w, http.StatusOK, task)
}
