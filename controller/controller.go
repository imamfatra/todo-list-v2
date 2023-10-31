package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type TodoController interface {
	Registrasi(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	Login(writer http.ResponseWriter, request *http.Request, params httprouter.Params)

	GetAllTodo(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	AddTodo(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	GetTodo(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	UpdateStatusTodo(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	DeleteTodo(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	GetRandomTodo(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	GetTodoFilter(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}
