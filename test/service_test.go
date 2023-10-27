package test_test

import (
	"context"
	"fmt"
	"testing"
	"todo-api/model"

	"github.com/stretchr/testify/require"
)

func registrasiSerivce(t *testing.T) model.RegistrasiRequest {
	// delTable(testDB)

	request := model.RegistrasiRequest{
		Email:    "name1@mail.com",
		Username: "mynameisyheone",
		Password: "secret1",
	}

	response, err := serviceTodo.Registrasi(context.Background(), request)
	require.NoError(t, err)
	require.NotEmpty(t, response)

	require.Equal(t, request.Email, response.Email)
	require.Equal(t, request.Username, response.Username)
	require.NotZero(t, response.Userid)
	fmt.Println(response)

	return request
}

func TestRegistrasiService(t *testing.T) {
	delTable(testDB)
	registrasiSerivce(t)
}

func TestLoginService(t *testing.T) {
	delTable(testDB)
	user := registrasiSerivce(t)

	request := model.LoginRequest{
		Username: user.Username,
		Password: user.Password,
	}
	response, err := serviceTodo.Login(context.Background(), request)
	require.NoError(t, err)
	require.NotEmpty(t, response)

	require.Equal(t, request.Username, response.Username)
	require.Equal(t, request.Username, response.Username)
	require.NotZero(t, response.Userid)
	fmt.Println(response)
}
