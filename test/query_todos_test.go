package test_test

import (
	"context"
	"database/sql"
	"fmt"
	"testing"
	"todo-api/repository"

	"github.com/stretchr/testify/require"
)

func TestAddNewTodo(t *testing.T) {
	delTable(testDB)

	addOneTodo(t)
}

func addOneTodo(t *testing.T) (int32, int32) {
	delTable(testDB)

	user, err := createAccount(t)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	arg := repository.AddaNewTodoParams{
		Todo:      "learning golang",
		Complated: false,
		Userid:    user.Userid,
	}

	todo, err := testQueries.AddaNewTodo(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, todo)

	require.NotZero(t, todo.ID)
	require.Equal(t, todo.Todo, arg.Todo)
	require.Equal(t, todo.Complated, arg.Complated)
	require.Equal(t, todo.Userid, user.Userid)

	return user.Userid, todo.ID
}

func addSomeTodo(t *testing.T) repository.User {
	user, err := createAccount(t)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	arg := repository.AddaNewTodoParams{
		Todo:      "learn golang",
		Complated: false,
		Userid:    user.Userid,
	}

	for i := 0; i < 15; i++ {
		todo, err := testQueries.AddaNewTodo(context.Background(), arg)
		require.NoError(t, err)
		require.NotEmpty(t, todo)
	}

	return user
}

func getIdTodo(db *sql.DB, userid int32) int32 {
	var id int32
	const query = "SELECT id FROM todos WHERE userid = $1 LIMIT 1"

	err := db.QueryRow(query, userid).Scan(&id)
	if err != nil {
		fmt.Println("error: ", err)
		return -1
	}

	return id
}

// func TestGetIdTodo(t *testing.T) {
// 	id, err := getIdTodo(testDB, 16)
// 	if err != nil {
// 		panic(err)
// 	}

// 	fmt.Println(id)
// }

func TestCountAllTodo(t *testing.T) {
	delTable(testDB)
	user := addSomeTodo(t)

	totalTodo, err := testQueries.CountAllTodos(context.Background(), user.Userid)
	require.NoError(t, err)
	require.NotZero(t, totalTodo)
	require.Equal(t, totalTodo, int64(15))
}

func TestGetSingleTodo(t *testing.T) {
	delTable(testDB)
	userid, id := addOneTodo(t)

	arg := repository.GetSingleaTodosParams{
		Userid: userid,
		ID:     id,
	}

	todo, err := testQueries.GetSingleaTodos(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, todo)
	require.False(t, todo.Complated)
	require.Equal(t, todo.Userid, userid)
	require.Equal(t, todo.ID, id)
}

func TestDeleteTodo(t *testing.T) {
	delTable(testDB)
	userid, id := addOneTodo(t)

	arg := repository.DeleteaTodoParams{
		ID:     id,
		Userid: userid,
	}

	todo, err := testQueries.DeleteaTodo(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, todo)

	require.Equal(t, todo.ID, id)
	require.Equal(t, todo.Userid, userid)
	require.True(t, todo.Isdelete)
	require.NotZero(t, todo.Deletedon)
}

func TestGetRandomTodoNoRow(t *testing.T) {

	todo, err := testQueries.GetRandomaTodo(context.Background())
	require.ErrorContains(t, err, fmt.Sprint(sql.ErrNoRows))
	fmt.Println(todo.ID)
}

func TestGetRandomSuccess(t *testing.T) {
	delTable(testDB)
	addSomeTodo(t)

	todo, err := testQueries.GetRandomaTodo(context.Background())
	require.NoError(t, err)
	require.NotEmpty(t, todo)
	fmt.Println(todo.ID)
}

func TestGetSomeTodo(t *testing.T) {
	delTable(testDB)
	user := addSomeTodo(t)

	arg := repository.GetSomeTodosParams{
		Userid: user.Userid,
		Limit:  5,
		Offset: 5,
	}

	todos, err := testQueries.GetSomeTodos(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, todos)

	require.Equal(t, len(todos), 5)

	for _, todo := range todos {
		require.NotEmpty(t, todo)
		require.Equal(t, todo.Userid, user.Userid)
	}

}

func TestUpdateStatusComplate(t *testing.T) {
	delTable(testDB)
	userid, id := addOneTodo(t)

	arg := repository.UpdateStatusComplateParams{
		ID:        id,
		Complated: true,
		Userid:    userid,
	}

	todo, err := testQueries.UpdateStatusComplate(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, todo)

	require.True(t, todo.Complated)
	require.Equal(t, todo.ID, id)
	require.Equal(t, todo.Userid, userid)
}
