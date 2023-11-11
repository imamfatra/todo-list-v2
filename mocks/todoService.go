package mocks

import (
	"context"
	"todo-api/model"
	"todo-api/repository"

	"github.com/stretchr/testify/mock"
)

type TodoServiceMock struct {
	mock.Mock
}

func (_m *TodoServiceMock) Registrasi(ctx context.Context, request repository.CreateAccountParams) model.RegistrasiResponse {
	argument := _m.Called(ctx, request)

	var result0 model.RegistrasiResponse
	if rf, ok := argument.Get(0).(func(context.Context, repository.CreateAccountParams) model.RegistrasiResponse); ok {
		result0 = rf(ctx, request)
	} else {
		if argument.Get(0) != nil {
			result0 = argument.Get(0).(model.RegistrasiResponse)
		}
	}

	return result0
}

func (_m *TodoServiceMock) Login(ctx context.Context, request model.LoginRequest) repository.User {
	argument := _m.Called(ctx, request)

	var result0 repository.User
	if rf, ok := argument.Get(0).(func(context.Context, model.LoginRequest) repository.User); ok {
		result0 = rf(ctx, request)
	} else {
		if argument.Get(0) != nil {
			result0 = argument.Get(0).(repository.User)
		}
	}

	return result0
}

func (_m *TodoServiceMock) GetAllTodo(ctx context.Context, request model.GetAllTodoRequest) model.GetAllTodosResponse {
	argument := _m.Called(ctx, request)

	var result0 model.GetAllTodosResponse
	if rf, ok := argument.Get(0).(func(context.Context, model.GetAllTodoRequest) model.GetAllTodosResponse); ok {
		result0 = rf(ctx, request)
	} else {
		if argument.Get(0) != nil {
			result0 = argument.Get(0).(model.GetAllTodosResponse)
		}
	}

	return result0
}

func (_m *TodoServiceMock) AddTodo(ctx context.Context, request model.AddNewTodoRequest) repository.AddaNewTodoRow {
	argument := _m.Called(ctx, request)

	var result0 repository.AddaNewTodoRow
	if rf, ok := argument.Get(0).(func(context.Context, model.AddNewTodoRequest) repository.AddaNewTodoRow); ok {
		result0 = rf(ctx, request)
	} else {
		if argument.Get(0) != nil {
			result0 = argument.Get(0).(repository.AddaNewTodoRow)
		}
	}

	return result0
}

func (_m *TodoServiceMock) GetTodo(ctx context.Context, request model.GetorDeleteTodoRequest) repository.GetSingleaTodosRow {
	argument := _m.Called(ctx, request)

	var result0 repository.GetSingleaTodosRow
	if rf, ok := argument.Get(0).(func(context.Context, model.GetorDeleteTodoRequest) repository.GetSingleaTodosRow); ok {
		result0 = rf(ctx, request)
	} else {
		if argument.Get(0) != nil {
			result0 = argument.Get(0).(repository.GetSingleaTodosRow)
		}
	}

	return result0
}

func (_m *TodoServiceMock) UpdateStatusTodo(ctx context.Context, request model.UpdateStatusTodoRequest) repository.UpdateStatusComplateRow {
	argument := _m.Called(ctx, request)

	var result0 repository.UpdateStatusComplateRow
	if rf, ok := argument.Get(0).(func(context.Context, model.UpdateStatusTodoRequest) repository.UpdateStatusComplateRow); ok {
		result0 = rf(ctx, request)
	} else {
		if argument.Get(0) != nil {
			result0 = argument.Get(0).(repository.UpdateStatusComplateRow)
		}
	}

	return result0
}

func (_m *TodoServiceMock) DeleteTodo(ctx context.Context, request model.GetorDeleteTodoRequest) repository.Todo {
	argument := _m.Called(ctx, request)

	var result0 repository.Todo
	if rf, ok := argument.Get(0).(func(context.Context, model.GetorDeleteTodoRequest) repository.Todo); ok {
		result0 = rf(ctx, request)
	} else {
		if argument.Get(0) != nil {
			result0 = argument.Get(0).(repository.Todo)
		}
	}

	return result0
}

func (_m *TodoServiceMock) GetRandomTodo(ctx context.Context, request model.GetAllTodoRequest) repository.GetRandomaTodoRow {
	argument := _m.Called(ctx, request)

	var result0 repository.GetRandomaTodoRow
	if rf, ok := argument.Get(0).(func(context.Context, model.GetAllTodoRequest) repository.GetRandomaTodoRow); ok {
		result0 = rf(ctx, request)
	} else {
		if argument.Get(0) != nil {
			result0 = argument.Get(0).(repository.GetRandomaTodoRow)
		}
	}

	return result0
}

func (_m *TodoServiceMock) GetTodoFilter(ctx context.Context, request model.GetTodoFilterRequest) model.GetTodoFilterResponse {
	argument := _m.Called(ctx, request)

	var result0 model.GetTodoFilterResponse
	if rf, ok := argument.Get(0).(func(context.Context, model.GetTodoFilterRequest) model.GetTodoFilterResponse); ok {
		result0 = rf(ctx, request)
	} else {
		if argument.Get(0) != nil {
			result0 = argument.Get(0).(model.GetTodoFilterResponse)
		}
	}

	return result0
}
