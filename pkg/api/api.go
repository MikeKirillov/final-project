package api

import (
	"net/http"
)

func Init() {
	http.HandleFunc("/api/nextdate", nextDayHandler)
	http.HandleFunc("/api/task", taskHandler)
	http.HandleFunc("/api/tasks", tasksHandler)
	http.HandleFunc("/api/task/done", doneTaskHandler)
}

func taskHandler(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodPost:
		addTaskHandler(w, req)
	case http.MethodGet:
		getTaskHandler(w, req)
	case http.MethodPut:
		updateTaskHandler(w, req)
	case http.MethodDelete:
		deleteTaskHandler(w, req)
	}
}
