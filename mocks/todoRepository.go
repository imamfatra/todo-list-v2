package mocks

import (
	"context"
	"todo-api/repository"

	"github.com/stretchr/testify/mock"
)

type TodoRepository struct {
	mock.Mock
}

func (_m *TodoRepository) AddaNewTodo(ctx context.Context, arg repository.AddaNewTodoParams) (repository.AddaNewTodoRow, error) {
	argument := _m.Called(ctx, arg)

	var result0 repository.AddaNewTodoRow
	if rf, ok := argument.Get(0).(func(context.Context, repository.AddaNewTodoParams) repository.AddaNewTodoRow); ok {
		result0 = rf(ctx, arg)
	} else {
		if argument.Get(0) != nil {
			result0 = argument.Get(0).(repository.AddaNewTodoRow)
		}
	}

	var result1 error
	if rf, ok := argument.Get(1).(func(context.Context, repository.AddaNewTodoParams) error); ok {
		result1 = rf(ctx, arg)
	} else {
		if argument.Get(1) != nil {
			result1 = argument.Get(1).(error)
		}
	}

	return result0, result1
}

func (_m *TodoRepository) CountAllTodos(ctx context.Context, userid int32) (int64, error) {
	argument := _m.Called(ctx, userid)

	var result0 int64
	if rf, ok := argument.Get(0).(func(context.Context, int32) int64); ok {
		result0 = rf(ctx, userid)
	} else {
		if argument.Get(0) != nil {
			result0 = argument.Get(0).(int64)
		}
	}

	var result1 error
	if rf, ok := argument.Get(1).(func(context.Context, int32) error); ok {
		result1 = rf(ctx, userid)
	} else {
		if argument.Get(1) != nil {
			result1 = argument.Get(1).(error)
		}
	}

	return result0, result1
}

func (_m *TodoRepository) CreateAccount(ctx context.Context, arg repository.CreateAccountParams) (repository.User, error) {
	argument := _m.Called(ctx, arg)

	var result0 repository.User
	if rf, ok := argument.Get(0).(func(context.Context, repository.CreateAccountParams) repository.User); ok {
		result0 = rf(ctx, arg)
	} else {
		if argument.Get(0) != nil {
			result0 = argument.Get(0).(repository.User)
		}
	}

	var result1 error
	if rf, ok := argument.Get(1).(func(context.Context, repository.CreateAccountParams) error); ok {
		result1 = rf(ctx, arg)
	} else {
		if argument.Get(1) != nil {
			result1 = argument.Get(1).(error)
		}
	}

	return result0, result1
}

func (_m *TodoRepository) DeleteaTodo(ctx context.Context, arg repository.DeleteaTodoParams) (repository.Todo, error) {
	argument := _m.Called(ctx, arg)

	var result0 repository.Todo
	if rf, ok := argument.Get(0).(func(context.Context, repository.DeleteaTodoParams) repository.Todo); ok {
		result0 = rf(ctx, arg)
	} else {
		if argument.Get(0) != nil {
			result0 = argument.Get(0).(repository.Todo)
		}
	}

	var result1 error
	if rf, ok := argument.Get(1).(func(context.Context, repository.DeleteaTodoParams) error); ok {
		result1 = rf(ctx, arg)
	} else {
		if argument.Get(1) != nil {
			result1 = argument.Get(1).(error)
		}
	}

	return result0, result1
}

func (_m *TodoRepository) GetAccount(ctx context.Context, username string) (repository.User, error) {
	argument := _m.Called(ctx, username)

	var result0 repository.User
	if rf, ok := argument.Get(0).(func(context.Context, string) repository.User); ok {
		result0 = rf(ctx, username)
	} else {
		if argument.Get(0) != nil {
			result0 = argument.Get(0).(repository.User)
		}
	}

	var result1 error
	if rf, ok := argument.Get(1).(func(context.Context, string) error); ok {
		result1 = rf(ctx, username)
	} else {
		if argument.Get(1) != nil {
			result1 = argument.Get(1).(error)
		}
	}

	return result0, result1
}

func (_m *TodoRepository) GetAllTodos(ctx context.Context, userid int32) ([]repository.GetAllTodosRow, error) {
	argument := _m.Called(ctx, userid)

	var result0 []repository.GetAllTodosRow
	if rf, ok := argument.Get(0).(func(context.Context, int32) []repository.GetAllTodosRow); ok {
		result0 = rf(ctx, userid)
	} else {
		if argument.Get(0) != nil {
			result0 = argument.Get(0).([]repository.GetAllTodosRow)
		}
	}

	var result1 error
	if rf, ok := argument.Get(1).(func(context.Context, int32) error); ok {
		result1 = rf(ctx, userid)
	} else {
		if argument.Get(1) != nil {
			result1 = argument.Get(1).(error)
		}
	}

	return result0, result1
}

func (_m *TodoRepository) GetRandomaTodo(ctx context.Context, userid int32) (repository.GetRandomaTodoRow, error) {
	argument := _m.Called(ctx, userid)

	var result0 repository.GetRandomaTodoRow
	if rf, ok := argument.Get(0).(func(context.Context, int32) repository.GetRandomaTodoRow); ok {
		result0 = rf(ctx, userid)
	} else {
		if argument.Get(0) != nil {
			result0 = argument.Get(0).(repository.GetRandomaTodoRow)
		}
	}

	var result1 error
	if rf, ok := argument.Get(1).(func(context.Context, int32) error); ok {
		result1 = rf(ctx, userid)
	} else {
		if argument.Get(1) != nil {
			result1 = argument.Get(1).(error)
		}
	}

	return result0, result1
}

func (_m *TodoRepository) GetSingleaTodos(ctx context.Context, arg repository.GetSingleaTodosParams) (repository.GetSingleaTodosRow, error) {
	argument := _m.Called(ctx, arg)

	var result0 repository.GetSingleaTodosRow
	if rf, ok := argument.Get(0).(func(context.Context, repository.GetSingleaTodosParams) repository.GetSingleaTodosRow); ok {
		result0 = rf(ctx, arg)
	} else {
		if argument.Get(0) != nil {
			result0 = argument.Get(0).(repository.GetSingleaTodosRow)
		}
	}

	var result1 error
	if rf, ok := argument.Get(1).(func(context.Context, repository.GetSingleaTodosParams) error); ok {
		result1 = rf(ctx, arg)
	} else {
		if argument.Get(1) != nil {
			result1 = argument.Get(1).(error)
		}
	}

	return result0, result1
}

func (_m *TodoRepository) GetSomeTodos(ctx context.Context, arg repository.GetSomeTodosParams) ([]repository.GetSomeTodosRow, error) {
	argument := _m.Called(ctx, arg)

	var result0 []repository.GetSomeTodosRow
	if rf, ok := argument.Get(0).(func(context.Context, repository.GetSomeTodosParams) []repository.GetSomeTodosRow); ok {
		result0 = rf(ctx, arg)
	} else {
		if argument.Get(0) != nil {
			result0 = argument.Get(0).([]repository.GetSomeTodosRow)
		}
	}

	var result1 error
	if rf, ok := argument.Get(1).(func(context.Context, repository.GetSomeTodosParams) error); ok {
		result1 = rf(ctx, arg)
	} else {
		if argument.Get(1) != nil {
			result1 = argument.Get(1).(error)
		}
	}

	return result0, result1
}

func (_m *TodoRepository) UpdateStatusComplate(ctx context.Context, arg repository.UpdateStatusComplateParams) (repository.UpdateStatusComplateRow, error) {
	argument := _m.Called(ctx, arg)

	var result0 repository.UpdateStatusComplateRow
	if rf, ok := argument.Get(0).(func(context.Context, repository.UpdateStatusComplateParams) repository.UpdateStatusComplateRow); ok {
		result0 = rf(ctx, arg)
	} else {
		if argument.Get(0) != nil {
			result0 = argument.Get(0).(repository.UpdateStatusComplateRow)
		}
	}

	var result1 error
	if rf, ok := argument.Get(1).(func(context.Context, repository.UpdateStatusComplateParams) error); ok {
		result1 = rf(ctx, arg)
	} else {
		if argument.Get(1) != nil {
			result1 = argument.Get(1).(error)
		}
	}

	return result0, result1
}
