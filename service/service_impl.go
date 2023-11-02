package service

import (
	"context"
	"database/sql"
	"todo-api/exception"
	"todo-api/helper"
	"todo-api/model"
	"todo-api/repository"

	"github.com/go-playground/validator/v10"
)

type TodoServiceImpl struct {
	// Query    *repository.Queries
	DB       *sql.DB
	Validate *validator.Validate
}

func NewTodoService(db *sql.DB, validate *validator.Validate) TodoService {
	return &TodoServiceImpl{
		DB:       db,
		Validate: validate,
	}
}

func (service *TodoServiceImpl) Registrasi(ctx context.Context, request repository.CreateAccountParams) model.RegistrasiResponse {
	err := service.Validate.Struct(request)
	helper.IfError(err)

	var user repository.User
	err = helper.ExecTx(ctx, service.DB, func(q *repository.Queries) error {
		arg := repository.CreateAccountParams{
			Email:    request.Email,
			Username: request.Username,
			Password: request.Password,
		}

		user, err = q.CreateAccount(ctx, arg)
		if err != nil {
			return err
		}
		return nil
	})
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

	var user repository.User
	err = helper.ExecTx(ctx, service.DB, func(q *repository.Queries) error {
		// arg := repository.GetAccountParams{
		// 	Username: request.Username,
		// 	Password: request.Password,
		// }

		user, err = q.GetAccount(ctx, request.Username)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return user
}

func (service *TodoServiceImpl) GetAllTodo(ctx context.Context, request model.GetAllTodoRequest) model.GetAllTodosResponse {
	err := service.Validate.Struct(request)
	helper.IfError(err)

	var todos []repository.GetAllTodosRow
	err = helper.ExecTx(ctx, service.DB, func(q *repository.Queries) error {
		userid := request.Userid
		todos, err = q.GetAllTodos(ctx, userid)
		if err != nil {
			return err
		}
		return nil
	})
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

	var user repository.AddaNewTodoRow
	err = helper.ExecTx(ctx, service.DB, func(q *repository.Queries) error {
		arg := repository.AddaNewTodoParams{
			Todo:      request.Todo,
			Complated: request.Complated,
			Userid:    request.Userid,
		}
		user, err = q.AddaNewTodo(ctx, arg)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return user
}

func (service *TodoServiceImpl) GetTodo(ctx context.Context, request model.GetorDeleteTodoRequest) repository.GetSingleaTodosRow {
	err := service.Validate.Struct(request)
	helper.IfError(err)

	var user repository.GetSingleaTodosRow
	err = helper.ExecTx(ctx, service.DB, func(q *repository.Queries) error {
		arg := repository.GetSingleaTodosParams{
			Userid: request.Userid,
			ID:     request.ID,
		}
		user, err = q.GetSingleaTodos(ctx, arg)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return user
}

func (service *TodoServiceImpl) UpdateStatusTodo(ctx context.Context, request model.UpdateStatusTodoRequest) repository.UpdateStatusComplateRow {
	err := service.Validate.Struct(request)
	helper.IfError(err)

	var user repository.UpdateStatusComplateRow
	err = helper.ExecTx(ctx, service.DB, func(q *repository.Queries) error {
		arg := repository.UpdateStatusComplateParams{
			ID:        request.ID,
			Complated: request.Complated,
			Userid:    request.Userid,
		}
		user, err = q.UpdateStatusComplate(ctx, arg)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	return user
}

func (service *TodoServiceImpl) DeleteTodo(ctx context.Context, request model.GetorDeleteTodoRequest) repository.Todo {
	err := service.Validate.Struct(request)
	helper.IfError(err)

	var user repository.Todo
	err = helper.ExecTx(ctx, service.DB, func(q *repository.Queries) error {
		arg := repository.DeleteaTodoParams{
			ID:     request.ID,
			Userid: request.Userid,
		}
		user, err = q.DeleteaTodo(ctx, arg)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return user
}

func (service *TodoServiceImpl) GetRandomTodo(ctx context.Context, request model.GetAllTodoRequest) repository.GetRandomaTodoRow {
	var user repository.GetRandomaTodoRow
	err := helper.ExecTx(ctx, service.DB, func(q *repository.Queries) error {
		var err error
		user, err = q.GetRandomaTodo(ctx, request.Userid)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return user
}

func (service *TodoServiceImpl) GetTodoFilter(ctx context.Context, request model.GetTodoFilterRequest) model.GetTodoFilterResponse {
	err := service.Validate.Struct(request)
	helper.IfError(err)

	var user []repository.GetSomeTodosRow
	err = helper.ExecTx(ctx, service.DB, func(q *repository.Queries) error {
		arg := repository.GetSomeTodosParams{
			Userid: request.Userid,
			Limit:  request.Limit,
			Offset: request.Offset,
		}
		user, err = q.GetSomeTodos(ctx, arg)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	result := model.GetTodoFilterResponse{
		Todos: user,
		Total: int32(len(user)),
		Skip:  request.Offset,
		Limit: request.Limit,
	}

	return result
}
