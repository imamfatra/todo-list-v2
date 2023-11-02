package app

import (
	"todo-api/controller"
	"todo-api/exception"

	"github.com/julienschmidt/httprouter"
)

func NewRouter(todoControler controller.TodoController) *httprouter.Router {
	router := httprouter.New()

	router.GET("/api/todos/:userId", todoControler.GetAllTodo)
	router.GET("/api/todos", todoControler.GetTodoFilter)
	router.GET("/api/todo/random/:userId", todoControler.GetRandomTodo)
	router.POST("/api/todo", todoControler.AddTodo)
	router.GET("/api/todo", todoControler.GetTodo)
	router.PUT("/api/todo", todoControler.UpdateStatusTodo)
	router.DELETE("/api/todo", todoControler.DeleteTodo)

	router.POST("/api/user/registrasi", todoControler.Registrasi)
	router.POST("/api/user/login", todoControler.Login)

	router.PanicHandler = exception.ErrorHandler

	return router
}
