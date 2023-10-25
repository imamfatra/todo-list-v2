package service

import (
	"context"
	"todo-api/model"
	"todo-api/repository"
)

type TodoService interface {
	registrasi(ctx context.Context, request model.RegistrasiRequest) (model.RegistrasiResponse, error)
	login(ctx context.Context, request model.LoginRequest) (model.LoginResponse, error)

	GetAllTodo(ctx context.Context, request model.GetAllTodoRequest) (model.GetAllTodosResponse, error)
	AddTodo(ctx context.Context, request model.AddNewTodoRequest) (repository.AddaNewTodoRow, error)
	GetTodo(ctx context.Context, request model.GetorDeleteTodoRequest) (repository.GetSingleaTodosRow, error)
	UpdateStatusTodo(ctx context.Context, request model.UpdateStatusTodoRequest) (repository.UpdateStatusComplateRow, error)
	DeleteTodo(ctx context.Context, request model.GetorDeleteTodoRequest) (repository.Todo, error)
	GetRandomTodo(ctx context.Context) (repository.GetRandomaTodoRow, error)
	GetTodoFilter(ctx context.Context, request model.GetTodoFilterRequest) (model.GetTodoFilterResponse, error)
}
