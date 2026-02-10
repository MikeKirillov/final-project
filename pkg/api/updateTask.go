package api

import (
	"errors"
	"log"
	"net/http"

	"final-project/pkg/db"
)

func updateTaskHandler(w http.ResponseWriter, req *http.Request) {
	var task db.Task
	readJson(w, req, &task)

	if len(task.Title) == 0 {
		err := errors.New("title must not be empty")
		log.Println(err)
		writeJson(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	if err := checkRepeat(&task); err != nil {
		log.Println(err)
		writeJson(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	if err := checkDate(&task); err != nil {
		log.Println(err)
		writeJson(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}

	if err := db.UpdateTask(&task); err != nil {
		log.Println(err)
		writeJson(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	} else {
		writeJson(w, http.StatusOK, struct{}{})
	}
}
