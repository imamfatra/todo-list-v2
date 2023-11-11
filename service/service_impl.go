package service

import (
	"context"
	"database/sql"
	"todo-api/exception"
	"todo-api/helper"
	"todo-api/model"
	"todo-api/repository"

	"github.com/go-playground/validator"
)

type TodoServiceImpl struct {
	Validate       *validator.Validate
	TodoRepository *repository.Queries
	DB             *sql.DB
}

func NewTodoService(todoRepository *repository.Queries, validate *validator.Validate, db *sql.DB) TodoService {
	return &TodoServiceImpl{
		TodoRepository: todoRepository,
		Validate:       validate,
		DB:             db,
	}
}

func (service *TodoServiceImpl) Registrasi(ctx context.Context, request model.RegistrasiRequest) model.RegistrasiResponse {
	err := service.Validate.Struct(request)
	helper.IfError(err)

	arg := repository.CreateAccountParams{
		Email:    request.Email,
		Username: request.Username,
		Password: request.Password,
	}

	tx, err := service.DB.Begin()
	defer helper.CommitOrRollback(tx)
	helper.IfError(err)

	query := service.TodoRepository.WithTx(tx)
	user, err := query.CreateAccount(ctx, arg)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	result := model.RegistrasiResponse{
		Email:    user.Email,
		Username: user.Username,
		Userid:   user.Userid,
	}
	return result
}

func (service *TodoServiceImpl) Login(ctx context.Context, request model.LoginRequest) repository.User {
	err := service.Validate.Struct(request)
	helper.IfError(err)

	tx, err := service.DB.Begin()
	defer helper.CommitOrRollback(tx)
	helper.IfError(err)

	query := service.TodoRepository.WithTx(tx)
	user, err := query.GetAccount(ctx, request.Username)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return user
}

func (service *TodoServiceImpl) GetAllTodo(ctx context.Context, request model.GetAllTodoRequest) model.GetAllTodosResponse {
	err := service.Validate.Struct(request)
	helper.IfError(err)

	userid := request.Userid
	tx, err := service.DB.Begin()
	defer helper.CommitOrRollback(tx)
	helper.IfError(err)

	query := service.TodoRepository.WithTx(tx)
	todos, err := query.GetAllTodos(ctx, userid)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	result := model.GetAllTodosResponse{
		Todos: todos,
		Total: int64(len(todos)),
	}

	return result
}

func (service *TodoServiceImpl) AddTodo(ctx context.Context, request model.AddNewTodoRequest) repository.AddaNewTodoRow {
	err := service.Validate.Struct(request)
	helper.IfError(err)

	arg := repository.AddaNewTodoParams{
		Todo:      request.Todo,
		Complated: request.Complated,
		Userid:    request.Userid,
	}
	tx, err := service.DB.Begin()
	defer helper.CommitOrRollback(tx)
	helper.IfError(err)

	query := service.TodoRepository.WithTx(tx)
	todo, err := query.AddaNewTodo(ctx, arg)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return todo
}

func (service *TodoServiceImpl) GetTodo(ctx context.Context, request model.GetorDeleteTodoRequest) repository.GetSingleaTodosRow {
	err := service.Validate.Struct(request)
	helper.IfError(err)

	arg := repository.GetSingleaTodosParams{
		Userid: request.Userid,
		ID:     request.ID,
	}
	tx, err := service.DB.Begin()
	defer helper.CommitOrRollback(tx)
	helper.IfError(err)

	query := service.TodoRepository.WithTx(tx)
	todo, err := query.GetSingleaTodos(ctx, arg)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return todo
}

func (service *TodoServiceImpl) UpdateStatusTodo(ctx context.Context, request model.UpdateStatusTodoRequest) repository.UpdateStatusComplateRow {
	err := service.Validate.Struct(request)
	helper.IfError(err)

	arg := repository.UpdateStatusComplateParams{
		ID:        request.ID,
		Complated: request.Complated,
		Userid:    request.Userid,
	}
	tx, err := service.DB.Begin()
	defer helper.CommitOrRollback(tx)
	helper.IfError(err)

	query := service.TodoRepository.WithTx(tx)
	todo, err := query.UpdateStatusComplate(ctx, arg)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return todo
}

func (service *TodoServiceImpl) DeleteTodo(ctx context.Context, request model.GetorDeleteTodoRequest) repository.Todo {
	err := service.Validate.Struct(request)
	helper.IfError(err)

	arg := repository.DeleteaTodoParams{
		ID:     request.ID,
		Userid: request.Userid,
	}
	tx, err := service.DB.Begin()
	defer helper.CommitOrRollback(tx)
	helper.IfError(err)

	query := service.TodoRepository.WithTx(tx)
	todo, err := query.DeleteaTodo(ctx, arg)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return todo
}

func (service *TodoServiceImpl) GetRandomTodo(ctx context.Context, request model.GetAllTodoRequest) repository.GetRandomaTodoRow {

	tx, err := service.DB.Begin()
	defer helper.CommitOrRollback(tx)
	helper.IfError(err)

	query := service.TodoRepository.WithTx(tx)
	todo, err := query.GetRandomaTodo(ctx, request.Userid)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return todo
}

func (service *TodoServiceImpl) GetTodoFilter(ctx context.Context, request model.GetTodoFilterRequest) model.GetTodoFilterResponse {
	err := service.Validate.Struct(request)
	helper.IfError(err)

	arg := repository.GetSomeTodosParams{
		Userid: request.Userid,
		Limit:  request.Limit,
		Offset: request.Offset,
	}
	tx, err := service.DB.Begin()
	defer helper.CommitOrRollback(tx)
	helper.IfError(err)

	query := service.TodoRepository.WithTx(tx)
	todos, err := query.GetSomeTodos(ctx, arg)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	result := model.GetTodoFilterResponse{
		Todos: todos,
		Total: int32(len(todos)),
		Skip:  request.Offset,
		Limit: request.Limit,
	}

	return result
}
