package service

import (
	"context"
	"database/sql"
	"todo-api/model"
	"todo-api/repository"

	"github.com/go-playground/validator/v10"
)

type TodoServiceImpl struct {
	Query    repository.Queries
	DB       *sql.DB
	Validate *validator.Validate
}

func NewTodoService(query repository.Queries, db *sql.DB, validate *validator.Validate) TodoService {
	return &TodoServiceImpl{
		Query:    query,
		DB:       db,
		Validate: validate,
	}
}

func (repository *TodoServiceImpl) registrasi(ctx context.Context, request model.RegistrasiRequest) (model.RegistrasiResponse, error) {
	panic("not implemented") // TODO: Implement
}

func (repository *TodoServiceImpl) login(ctx context.Context, request model.LoginRequest) (model.LoginResponse, error) {
	panic("not implemented") // TODO: Implement
}

func (repository *TodoServiceImpl) GetAllTodo(ctx context.Context, userid model.GetAllTodoRequest) (model.GetAllTodosResponse, error) {
	panic("not implemented") // TODO: Implement
}

func (repository *TodoServiceImpl) AddTodo(ctx context.Context, request model.AddNewTodoRequest) (repository.AddaNewTodoRow, error) {
	panic("not implemented") // TODO: Implement
}

func (repository *TodoServiceImpl) GetTodo(ctx context.Context, request model.GetorDeleteTodoRequest) (repository.GetSingleaTodosRow, error) {
	panic("not implemented") // TODO: Implement
}

func (repository *TodoServiceImpl) UpdateStatusTodo(ctx context.Context, request model.UpdateStatusTodoRequest) (repository.UpdateStatusComplateRow, error) {
	panic("not implemented") // TODO: Implement
}

func (repository *TodoServiceImpl) DeleteTodo(ctx context.Context, request model.GetorDeleteTodoRequest) (repository.Todo, error) {
	panic("not implemented") // TODO: Implement
}

func (repository *TodoServiceImpl) GetRandomTodo(ctx context.Context) (repository.GetRandomaTodoRow, error) {
	panic("not implemented") // TODO: Implement
}

func (repository *TodoServiceImpl) GetTodoFilter(ctx context.Context, request model.GetTodoFilterRequest) (model.GetTodoFilterResponse, error) {
	panic("not implemented") // TODO: Implement
}
