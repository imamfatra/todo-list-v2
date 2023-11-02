package main

import (
	"net/http"
	"todo-api/app"
	"todo-api/controller"
	"todo-api/helper"
	"todo-api/middleware"
	"todo-api/service"

	"github.com/go-playground/validator/v10"
	_ "github.com/lib/pq"
)

func main() {
	db := app.NewDB()

	var validate = validator.New()
	serviceTodo := service.NewTodoService(db, validate)
	controllerTodo := controller.NewTodoController(serviceTodo)
	// var hendler http.Handler
	hendler := app.NewRouter(controllerTodo)

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: middleware.Auth(hendler),
	}
	err := server.ListenAndServe()
	helper.IfError(err)
}
