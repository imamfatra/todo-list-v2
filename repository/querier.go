// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.21.0

package repository

import (
	"context"
)

type Querier interface {
	AddaNewTodo(ctx context.Context, arg AddaNewTodoParams) (AddaNewTodoRow, error)
	CountAllTodos(ctx context.Context, userid int32) (int64, error)
	CreateAccount(ctx context.Context, arg CreateAccountParams) (User, error)
	DeleteaTodo(ctx context.Context, arg DeleteaTodoParams) (Todo, error)
	GetAccount(ctx context.Context, username string) (User, error)
	GetAllTodos(ctx context.Context, userid int32) ([]GetAllTodosRow, error)
	GetRandomaTodo(ctx context.Context, userid int32) (GetRandomaTodoRow, error)
	GetSingleaTodos(ctx context.Context, arg GetSingleaTodosParams) (GetSingleaTodosRow, error)
	GetSomeTodos(ctx context.Context, arg GetSomeTodosParams) ([]GetSomeTodosRow, error)
	UpdateStatusComplate(ctx context.Context, arg UpdateStatusComplateParams) (UpdateStatusComplateRow, error)
}

var _ Querier = (*Queries)(nil)
