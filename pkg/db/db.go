package db

import (
	"database/sql"
	"fmt"
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

func Init(dbFile string) error {
	fmt.Println("DB initialization")

	_, err := os.Stat(dbFile)

	var install bool
	if err != nil {
		fmt.Println("There's no file")
		install = true
	}

	db, err := sql.Open("sqlite", dbFile)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer db.Close()

	if install == true {
		fmt.Println("Creating new schema")

		_, err = db.Exec(schema)
		if err != nil {
			fmt.Println(err)
			return err
		}

		fmt.Println("Table 'scheduler' created successfully (if it did not already exist)")
	}

	return nil
}
