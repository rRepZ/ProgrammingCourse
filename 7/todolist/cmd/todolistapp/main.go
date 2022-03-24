package main

import (
	"database/sql"
	"log"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

const (
	connStr = "user=postgres password=mypass dbname=productdb sslmode=disable"
)

// реализовать регистрацию пользователя

// реализовать создание, просмотр, обновление и удаление задач у пользователя

func main() {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	userRepo := NewUserRepo(db)

	userHandler := NewUserHandler(userRepo)

	r := mux.NewRouter()
	r.HandleFunc("/users/create", UserHandler.Create)

}
