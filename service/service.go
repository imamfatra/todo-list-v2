package service

import (
	"context"
	"todo-api/model"
	"todo-api/repository"
)

type TodoService interface {
	Registrasi(ctx context.Context, request model.RegistrasiRequest) model.RegistrasiResponse
	Login(ctx context.Context, request model.LoginRequest) repository.User

	GetAllTodo(ctx context.Context, request model.GetAllTodoRequest) model.GetAllTodosResponse
	AddTodo(ctx context.Context, request model.AddNewTodoRequest) repository.AddaNewTodoRow
	GetTodo(ctx context.Context, request model.GetorDeleteTodoRequest) repository.GetSingleaTodosRow
	UpdateStatusTodo(ctx context.Context, request model.UpdateStatusTodoRequest) repository.UpdateStatusComplateRow
	DeleteTodo(ctx context.Context, request model.GetorDeleteTodoRequest) repository.Todo
	GetRandomTodo(ctx context.Context) repository.GetRandomaTodoRow
	GetTodoFilter(ctx context.Context, request model.GetTodoFilterRequest) model.GetTodoFilterResponse
}
