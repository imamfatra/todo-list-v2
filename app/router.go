package app

import (
	"todo-api/controller"
	"todo-api/exception"

	"github.com/julienschmidt/httprouter"
)

func NewRouter(todoControler controller.TodoController) *httprouter.Router {
	router := httprouter.New()

	// router.GET("/api/todos", todoControllr.GetAllTodo)
	// // router.GET("/api/todos/random", todoControllr.GetRandomTodo)
	// router.POST("/api/todos", todoControllr.AddTodo)
	// router.GET("/api/todos/:id", todoControllr.GetTodo)
	// router.PUT("/api/todos/:id", todoControllr.UpdateStatusTodo)
	// router.DELETE("/api/todos/:id", todoControllr.DeleteTodo)
	// router.GET("/api/todos/:limit/:skip", todoControllr.GetAllTodo)

	router.POST("/api/user/registrasi", todoControler.Registrasi)
	router.POST("/api/user/login", todoControler.Login)

	router.PanicHandler = exception.ErrorHandler

	return router
}
