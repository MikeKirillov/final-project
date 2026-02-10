package api

import (
	"errors"
	"log"
	"net/http"

	"final-project/pkg/db"
)

func deleteTaskHandler(w http.ResponseWriter, req *http.Request) {
	id := req.FormValue("id")
	if len(id) == 0 {
		err := errors.New("id must be specified")
		log.Println(err)
		writeJson(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}

	_, err := db.GetTask(id)
	// returns error if there's no rows by id: 'sql: no rows in result set'
	if err != nil {
		log.Println(err)
		writeJson(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}

	if err = db.DeleteTask(id); err != nil {
		log.Println(err)
		writeJson(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	} else {
		writeJson(w, http.StatusOK, struct{}{})
	}
}
