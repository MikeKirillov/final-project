package api

import (
	"final-project/pkg/db"
	"log"
	"net/http"
)

type TasksResp struct {
	Tasks []*db.Task `json:"tasks"`
}

func tasksHandler(w http.ResponseWriter, req *http.Request) {
	if tasks, err := db.Tasks(50); err != nil {
		log.Println(err)
		writeJson(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	} else {
		writeJson(w, http.StatusOK, TasksResp{Tasks: tasks})
	}
}
