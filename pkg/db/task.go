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

const SELECT_BY_LIMIT = `
SELECT * FROM scheduler ORDER BY date LIMIT ?`

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
	var tasks []*Task

	rows, err := db.Query(SELECT_BY_LIMIT, limit)
	if err != nil {
		log.Println(err)
		return tasks, err
	}
	defer rows.Close()

	for rows.Next() {
		task := Task{}

		err := rows.Scan(&task.ID, &task.Date, &task.Title, &task.Comment, &task.Repeat)
		if err != nil {
			log.Println(err)
			return tasks, err
		}

		tasks = append(tasks, &task)
	}

	if err := rows.Err(); err != nil {
		log.Println(err)
		return tasks, nil
	}
	// to avoid responce like {"tasks":null}, it's better to create an empty slice
	if tasks == nil {
		tasks = []*Task{}
	}
	return tasks, nil
}
