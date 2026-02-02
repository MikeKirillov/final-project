package db

import (
	"database/sql"
	"log"
)

type Task struct {
	ID      string `json:"id"`
	Date    string `json:"date"`
	Title   string `json:"title"`
	Comment string `json:"comment"`
	Repeat  string `json:"repeat"`
}

const INSERT = `
INSERT INTO scheduler (date, title, comment, repeat)
	VALUES (:date, :title, :comment, :repeat)`

func AddTask(task *Task) (int64, error) {
	var id int64

	res, err := db.Exec(INSERT,
		sql.Named("date", task.Date),
		sql.Named("title", task.Title),
		sql.Named("comment", task.Comment),
		sql.Named("repeat", task.Repeat))
	if err == nil {
		id, err = res.LastInsertId()
	}
	if err != nil {
		log.Println(err)
	}
	// if err != nil, then id returns as '0' because
	// of https://go.dev/ref/spec#The_zero_value
	return id, err
}

func Tasks(limit int) ([]*Task, error) {
	// rows, err := db.Query("")
	// if err != nil {
	// 	log.Println(err)
	// 	return nil, err
	// }

	return nil, nil
}
