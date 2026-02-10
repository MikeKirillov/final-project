package db

import (
	"database/sql"
	"log"
	"os"

	_ "modernc.org/sqlite"
)

const schema = `
CREATE TABLE scheduler (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    date CHAR(8) NOT NULL DEFAULT "",
	title VARCHAR(128) NOT NULL DEFAULT "",
	comment TEXT NULL,
	repeat VARCHAR(128) NOT NULL DEFAULT ""
);
CREATE INDEX scheduler_date ON scheduler (date);`

var db *sql.DB

func Init(dbFile string) error {
	log.Println("DB initialization")

	_, err := os.Stat(dbFile)

	var install bool
	if err != nil {
		log.Println("There's no file")
		install = true
	}

	db, err = sql.Open("sqlite", dbFile)
	if err != nil {
		log.Println(err)
		defer db.Close()
		return err
	}

	if install == true {
		log.Println("Creating new schema")

		_, err = db.Exec(schema)
		if err != nil {
			log.Println(err)
			defer db.Close()
			return err
		}

		log.Println("Table 'scheduler' created successfully (if it did not already exist)")
	}

	return nil
}

func Close() {
	defer db.Close()
}
