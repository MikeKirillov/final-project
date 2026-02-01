package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"final-project/pkg/db"
	"net/http"
	"slices"
	"strconv"
	"strings"
	"time"
)

type AddTaskRs struct {
	Id string `json:"id"`
}

type ErrorRs struct {
	Err string `json:"error"`
}

func addTaskHandler(w http.ResponseWriter, req *http.Request) {
	var task db.Task
	var addTaskRs AddTaskRs
	var errorRs ErrorRs

	readJson(w, req, &task)

	if len(task.Title) == 0 {
		err := errors.New("title must not be empty")
		errorRs.Err = err.Error()
		writeJson(w, http.StatusBadRequest, &errorRs)
		return
	}

	if err := checkRepeat(&task); err != nil {
		errorRs.Err = err.Error()
		writeJson(w, http.StatusBadRequest, &errorRs)
		return
	}

	if err := checkDate(&task); err != nil {
		errorRs.Err = err.Error()
		writeJson(w, http.StatusBadRequest, &errorRs)
		return
	}

	id, err := db.AddTask(&task)
	if err != nil {
		errorRs.Err = err.Error()
		writeJson(w, http.StatusBadRequest, &errorRs)
		return
	}

	addTaskRs.Id = strconv.FormatInt(id, 10)
	writeJson(w, http.StatusOK, &addTaskRs)
}

func readJson(w http.ResponseWriter, req *http.Request, task *db.Task) {
	var buf bytes.Buffer

	_, err := buf.ReadFrom(req.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err = json.Unmarshal(buf.Bytes(), &task); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func writeJson(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(status)

	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, "error encoding JSON", http.StatusInternalServerError)
	}
}

func checkDate(task *db.Task) error {
	now := time.Now()
	nowString := now.Format(LAYOUT)
	//if task.Date is an empty string then put time.Now() into
	if len(task.Date) == 0 {
		task.Date = nowString
	}
	// check task.Date for correct format
	t, err := time.Parse(LAYOUT, task.Date)
	if err != nil {
		return err
	}

	if task.Date != nowString && t.Before(now) {
		if len(task.Repeat) != 0 {
			next, _ := NextDate(now, task.Date, task.Repeat)
			task.Date = next
		} else {
			task.Date = nowString
		}
	}
	return nil
}

func checkRepeat(task *db.Task) error {
	splitedRep := strings.Split(task.Repeat, " ")
	repTypes := []string{"y", "d"}

	if len(splitedRep[0]) != 0 && !slices.Contains(repTypes, splitedRep[0]) {
		return errors.New("the 'repeat' value contains an invalid character")
	}
	return nil
}
