package db

import (
	"database/sql"
	"fmt"
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

const SELECT_BY_ID = `
SELECT * FROM scheduler WHERE id = :id`

const UPDATE_ROW = `
UPDATE scheduler SET date = :date, title = :title, comment = :comment, repeat = :repeat WHERE id = :id`

const DELETE_BY_ID = `
DELETE FROM cheduler WHERE id = :id`

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
	// if err != nil, then id returns as '0' because
	// of https://go.dev/ref/spec#The_zero_value
	return id, err
}

func Tasks(limit int) ([]*Task, error) {
	var tasks []*Task

	rows, err := db.Query(SELECT_BY_LIMIT, limit)
	if err != nil {
		return tasks, err
	}
	defer rows.Close()

	for rows.Next() {
		task := Task{}

		err := rows.Scan(&task.ID, &task.Date, &task.Title, &task.Comment, &task.Repeat)
		if err != nil {
			return tasks, err
		}

		tasks = append(tasks, &task)
	}

	if err := rows.Err(); err != nil {
		return tasks, nil
	}
	// to avoid responce like {"tasks":null}, it's better to create an empty slice: {"tasks":[]}
	if tasks == nil {
		tasks = []*Task{}
	}
	return tasks, nil
}

func GetTask(id string) (*Task, error) {
	var task = Task{}

	row := db.QueryRow(SELECT_BY_ID, sql.Named("id", id))
	err := row.Scan(&task.ID, &task.Date, &task.Title, &task.Comment, &task.Repeat)
	if err != nil {
		return nil, err
	}

	return &task, nil
}

func UpdateTask(task *Task) error {
	res, err := db.Exec(UPDATE_ROW,
		sql.Named("id", task.ID),
		sql.Named("date", task.Date),
		sql.Named("title", task.Title),
		sql.Named("comment", task.Comment),
		sql.Named("repeat", task.Repeat))
	if err != nil {
		return err
	}

	count, err := res.RowsAffected() // returns count of updated rows
	if err != nil {
		return err
	}

	if count == 0 {
		return fmt.Errorf("incorrect id for updating task")
	}

	return nil
}

func UpdateTaskDone(id string) error {
	return nil
}

func DeleteTask(id string) error {
	_, err := db.Exec(DELETE_BY_ID, sql.Named("id", id))
	if err != nil {
		return err
	}
	return nil
}
