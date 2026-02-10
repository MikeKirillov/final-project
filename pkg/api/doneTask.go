package api

import (
	"errors"
	"log"
	"net/http"
	"strings"
	"time"

	"final-project/pkg/db"
)

func doneTaskHandler(w http.ResponseWriter, req *http.Request) {
	id := req.FormValue("id")
	if len(id) == 0 {
		err := errors.New("id must be specified")
		log.Println(err)
		writeJson(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}

	task, err := db.GetTask(id)
	// returns error if there's no rows by id: 'sql: no rows in result set'
	if err != nil {
		log.Println(err)
		writeJson(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}

	if len(strings.TrimSpace(task.Repeat)) == 0 {
		err = db.DeleteTask(id)
		if err != nil {
			log.Println(err)
			writeJson(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
			return
		}
	} else {
		newData, err := NextDate(time.Now(), task.Date, task.Repeat)
		if err != nil {
			log.Println(err)
			writeJson(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
			return
		}

		task.Date = newData

		err = db.UpdateTask(task)
		if err != nil {
			log.Println(err)
			writeJson(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
			return
		}
	}

	writeJson(w, http.StatusOK, struct{}{})
}
