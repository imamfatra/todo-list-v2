package main

import (
	"database/sql"
	"log"
	"net/http"
	"todo-api/app"
	"todo-api/controller"
	"todo-api/helper"
	"todo-api/model"
	"todo-api/service"

	"github.com/go-playground/validator/v10"
	_ "github.com/lib/pq"
)

func main() {

	config, err := model.LoadConfig("./")
	if err != nil {
		log.Fatal("cannot load config: ", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}
	// query = repository.New(db)
	var validate = validator.New()
	serviceTodo := service.NewTodoService(conn, validate)
	controllerTodo := controller.NewTodoController(serviceTodo)
	var hendler http.Handler
	hendler = app.NewRouter(controllerTodo)

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: hendler,
	}
	err = server.ListenAndServe()
	helper.IfError(err)
}
