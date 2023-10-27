package test_test

import (
	"context"
	"testing"
	"todo-api/model"
	"todo-api/repository"

	"github.com/stretchr/testify/require"
)

func createAccount(t *testing.T) (repository.User, error) {

	arg := repository.CreateAccountParams{
		Username: model.RandomString(6),
		Email:    model.RandomMail(),
		Password: "secret",
	}

	user, err := testQueries.CreateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, user.Email, arg.Email)
	require.Equal(t, user.Username, arg.Username)
	require.Equal(t, user.Password, arg.Password)
	require.NotZero(t, user.Userid)

	return user, err
}

func TestCreateAccount(t *testing.T) {
	delTable(testDB)
	createAccount(t)
}

func TestGetAccountSuccess(t *testing.T) {
	delTable(testDB)

	user1, err := createAccount(t)
	require.NoError(t, err)
	require.NotEmpty(t, user1)

	arg := repository.GetAccountParams{
		Username: user1.Username,
		Password: user1.Password,
	}

	user2, err := testQueries.GetAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, user1.Email, user2.Email)
	require.Equal(t, user1.Username, user2.Username)
	require.Equal(t, user1.Userid, user2.Userid)

}

func TestGetAccountFailed(t *testing.T) {
	delTable(testDB)

	user1, err := createAccount(t)
	require.NoError(t, err)
	require.NotEmpty(t, user1)

	arg := repository.GetAccountParams{
		Username: user1.Username,
		Password: "Alda",
	}

	user2, err := testQueries.GetAccount(context.Background(), arg)
	require.Error(t, err)
	require.Empty(t, user2)

}
