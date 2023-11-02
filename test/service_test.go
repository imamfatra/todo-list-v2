package test_test

import (
	"context"
	"fmt"
	"testing"
	"todo-api/model"
	"todo-api/repository"

	"github.com/stretchr/testify/require"
)

func registrasiSerivce(t *testing.T) repository.CreateAccountParams {
	// delTable(testDB)

	request := repository.CreateAccountParams{
		Email:    "name1@mail.com",
		Username: "mynameisyheone",
		Password: "secret1",
	}

	response := serviceTodo.Registrasi(context.Background(), request)
	// require.NoError(t, err)
	require.NotEmpty(t, response)

	require.Equal(t, request.Email, response.Email)
	require.Equal(t, request.Username, response.Username)
	require.NotZero(t, response.Userid)
	fmt.Println(response)

	return request
}

func TestRegistrasiServiceSuccess(t *testing.T) {
	delTable(testDB)
	registrasiSerivce(t)
}

func TestLoginServiceSuccess(t *testing.T) {
	delTable(testDB)
	user := registrasiSerivce(t)

	request := model.LoginRequest{
		Username: user.Username,
		Password: user.Password,
	}
	response := serviceTodo.Login(context.Background(), request)
	// require.NoError(t, err)
	require.NotEmpty(t, response)

	require.Equal(t, request.Username, response.Username)
	require.Equal(t, request.Username, response.Username)
	require.NotZero(t, response.Userid)
	fmt.Println(response)
}

func TestUpdateTodo(t *testing.T) {
	request := model.UpdateStatusTodoRequest{
		ID:        367,
		Complated: false,
		Userid:    54,
	}

	response := serviceTodo.UpdateStatusTodo(context.Background(), request)
	require.Equal(t, response.Complated, request.Complated)
	require.Equal(t, response.Userid, request.Userid)
	require.Equal(t, response.ID, request.ID)

}

// func TestLoginFailed(t *testing.T) {
// 	reques := model.LoginRequest{
// 		Username: "Anonim1",
// 		Password: "123abcd",
// 	}
// 	response := serviceTodo.Login(context.Background(), reques)
// 	require.Error(t, response)
// }
