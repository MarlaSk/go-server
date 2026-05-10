package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"sandbox/internal/config"
	"sandbox/internal/handler"
	"sandbox/internal/service"
	"sandbox/internal/storage"

	_ "github.com/lib/pq"
)

func main() {
	// mux := http.NewServeMux()

	// mux.HandleFunc("/todo/", getToDoById)
	// mux.HandleFunc("/create", createToDo)
	// mux.HandleFunc("/delete/", deleteToDo)
	// mux.HandleFunc("/update/", updateToDo)
	// mux.HandleFunc("/ping", health)

	// fmt.Println("Запускаю сервер на :8080")

	// err := http.ListenAndServe(":8080", mux)
	// if err != nil {
	// 	fmt.Println("Не смог поднять сервер:", err)
	// }
	cfg, err := config.Load("config/config.json")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(cfg)
	fmt.Println(cfg.GetDatabaseDSN())

	db, err := sql.Open("postgres", cfg.GetDatabaseDSN())
	if err != nil {
		log.Fatal("Не удалось подключиться к базе данных:", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal("Не удалось пингануть к базе данных:", err)
	}

	if err := storage.CreateTodoTable(context.Background(), db); err != nil {
		log.Fatal("Не удалось создать таблицу:", err)
	}
	//Слой хранения данных
	repo := storage.NewPostgresRepo(db)

	//Слой бизнес логики
	service := service.NewTodoService(repo)

	//Транспортный слой
	handler := handler.NewTodoHandler(service)

	//Создаем сервер
	mux := http.NewServeMux()

	mux.HandleFunc("/create", handler.CreateToDo)
	mux.HandleFunc("/todo/", handler.GetToDoById)

	http.ListenAndServe(cfg.GetServerPort(), mux)
}
