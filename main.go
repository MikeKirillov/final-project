package main

import (
	"final-project/pkg/db"
	"final-project/pkg/server"
	"log"
)

func main() {
	err := db.Init("scheduler.db")

	if err != nil {
		log.Fatal(err)
	}

	server.Start()
}

// Как вариант, можете создать директорию pkg с тремя поддиректориями:
// api — будет содержать файлы с обработчиками API запросов;
// db — файлы с кодом, отвечающие за работу с базой данных;
// server — будет расположен файл с функцией запуска сервера.
// Файлы main.go и go.mod располагаются непосредственно в директории проекта.
