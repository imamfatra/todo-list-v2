package main

import (
	"net/http"
	"todo-api/app"
	"todo-api/controller"
	"todo-api/helper"
	"todo-api/middleware"
	"todo-api/repository"
	"todo-api/service"

	"github.com/go-playground/validator/v10"
	_ "github.com/lib/pq"
)

func main() {
	db := app.NewDB()
	validate := validator.New()

	serviceTodo := service.NewTodoService(&repository.Queries{}, validate, db)
	controllerTodo := controller.NewTodoController(serviceTodo)
	hendler := app.NewRouter(controllerTodo)

	server := http.Server{
		Addr:    "0.0.0.0:3000",
		Handler: middleware.Auth(hendler),
	}
	err := server.ListenAndServe()
	helper.IfError(err)
}
