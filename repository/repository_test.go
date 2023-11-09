package repository_test

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"testing"
	"todo-api/model"
	"todo-api/repository"

	_ "github.com/lib/pq"
	"github.com/stretchr/testify/suite"
)

type todoRepositorySuite struct {
	suite.Suite
	repository      repository.Queries
	cleanUpDatabase model.TruncateTableExecutor
}

func (suite *todoRepositorySuite) SetupSuite() {
	config, err := model.LoadConfig("../.")
	if err != nil {
		log.Fatal("cannot load config: ", err)
	}

	db, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}
	repository := repository.New(db)

	suite.repository = *repository
	suite.cleanUpDatabase = model.InitTruncateTableExecutor(db)
}

func (suite *todoRepositorySuite) TearDownTest() {
	defer suite.cleanUpDatabase.TruncateTable([]string{"users", "todos"})
}

func (suite *todoRepositorySuite) createAccount_positive() (repository.User, error) {
	arg := repository.CreateAccountParams{
		Email:    model.RandomMail(),
		Username: model.RandomString(10),
		Password: "secret19",
	}

	user, err := suite.repository.CreateAccount(context.Background(), arg)
	suite.Nil(err, "No error when create account")
	suite.NotEmpty(user)

	suite.Equal(user.Email, arg.Email)
	suite.Equal(user.Username, arg.Username)
	suite.Equal(user.Password, arg.Password)
	suite.NotZero(user.Userid)

	return user, err
}

func (suite *todoRepositorySuite) TestCreateAccout_positive() {
	user, err := suite.createAccount_positive()
	suite.Nil(err)
	suite.NotEmpty(user)
}

func (suite *todoRepositorySuite) TestCreateAccount_EmptyField_Positive() {

	user, err := suite.repository.CreateAccount(context.Background(), repository.CreateAccountParams{})
	suite.Nil(err)
	suite.NotEmpty(user)
}

func (suite *todoRepositorySuite) TestGetAccout_Positive() {
	user1, err := suite.createAccount_positive()
	suite.NoError(err)
	suite.NotEmpty(user1)

	user2, err := suite.repository.GetAccount(context.Background(), user1.Username)
	suite.NoError(err)
	suite.NotEmpty(user2)

	suite.Equal(user1.Email, user2.Email)
	suite.Equal(user1.Username, user2.Username)
	suite.Equal(user1.Userid, user2.Userid)
}

func (suite *todoRepositorySuite) TestGetAccout_Negative() {
	user1, err := suite.createAccount_positive()
	suite.NoError(err)
	suite.NotEmpty(user1)

	user2, err := suite.repository.GetAccount(context.Background(), "abasuf")
	suite.Error(err)
	suite.Empty(user2)
}

func (suite *todoRepositorySuite) addOneTodo() (int32, int32) {
	user, err := suite.createAccount_positive()
	suite.NoError(err)
	suite.NotEmpty(user)

	arg := repository.AddaNewTodoParams{
		Todo:      model.RandomString(35),
		Complated: false,
		Userid:    user.Userid,
	}

	todo, err := suite.repository.AddaNewTodo(context.Background(), arg)
	suite.NoError(err)
	suite.NotEmpty(todo)

	suite.NotZero(todo.ID)
	suite.Equal(todo.Todo, arg.Todo)
	suite.Equal(todo.Complated, arg.Complated)
	suite.Equal(todo.Userid, user.Userid)

	return user.Userid, todo.ID
}

func (suite *todoRepositorySuite) addManyTodo() int32 {
	user, err := suite.createAccount_positive()
	suite.NoError(err)
	suite.NotEmpty(user)

	arg := repository.AddaNewTodoParams{
		Todo:      model.RandomString(35),
		Complated: false,
		Userid:    user.Userid,
	}

	for i := 0; i < 100; i++ {
		todo, err := suite.repository.AddaNewTodo(context.Background(), arg)
		suite.NoError(err)
		suite.NotEmpty(todo)
	}

	return user.Userid
}

func (suite *todoRepositorySuite) TestAddNewTodo_Positive() {
	userid, id := suite.addOneTodo()
	suite.NotZero(userid)
	suite.NotZero(id)
	suite.Equal(userid, int32(1))
	suite.Equal(id, int32(1))
}

func (suite *todoRepositorySuite) TestCoutAlltodo_Positive() {
	userId := suite.addManyTodo()

	totalTodo, err := suite.repository.CountAllTodos(context.Background(), userId)
	suite.NoError(err)
	suite.NotZero(totalTodo)
	suite.Equal(totalTodo, int64(100))
}

func (suite *todoRepositorySuite) TestCoutAlltodo_Negative() {
	userId := suite.addManyTodo()

	totalTodo, err := suite.repository.CountAllTodos(context.Background(), userId)
	suite.NoError(err)
	suite.NotZero(totalTodo)
	suite.NotEqual(totalTodo, int64(10))
}

func (suite *todoRepositorySuite) TestGetOneTodo_Positive() {
	userid, id := suite.addOneTodo()

	arg := repository.GetSingleaTodosParams{
		Userid: userid,
		ID:     id,
	}

	todo, err := suite.repository.GetSingleaTodos(context.Background(), arg)
	suite.NoError(err)
	suite.NotEmpty(todo)
	suite.False(todo.Complated)
	suite.Equal(todo.Userid, userid)
	suite.Equal(todo.ID, id)
}

func (suite *todoRepositorySuite) TestGetOneTodo_Negative() {
	userid, id := suite.addOneTodo()

	arg := repository.GetSingleaTodosParams{
		Userid: userid,
		ID:     id,
	}

	todo, err := suite.repository.GetSingleaTodos(context.Background(), arg)
	suite.NoError(err)
	suite.NotEmpty(todo)
	suite.False(todo.Complated)
	suite.Equal(todo.Userid, userid)
	suite.Equal(todo.ID, id)
}

func (suite *todoRepositorySuite) TestDelete_Positive() {
	userId, id := suite.addOneTodo()

	arg := repository.DeleteaTodoParams{
		ID:     id,
		Userid: userId,
	}

	todo, err := suite.repository.DeleteaTodo(context.Background(), arg)
	suite.NoError(err)
	suite.NotEmpty(todo)

	suite.Equal(todo.ID, id)
	suite.Equal(todo.Userid, userId)
	suite.True(todo.Isdelete)
	suite.NotZero(todo.Deletedon)

}

func (suite *todoRepositorySuite) TestDelete_Negative() {
	userId, _ := suite.addOneTodo()

	arg := repository.DeleteaTodoParams{
		ID:     int32(1000),
		Userid: userId,
	}

	todo, err := suite.repository.DeleteaTodo(context.Background(), arg)
	suite.Error(err)
	suite.Empty(todo)
}

func (suite *todoRepositorySuite) TestGetRandomTodo_Positive() {
	userId := suite.addManyTodo()

	todo, err := suite.repository.GetRandomaTodo(context.Background(), userId)
	suite.NoError(err)
	suite.NotEmpty(todo)
	suite.NotZero(todo.ID)
}

func (suite *todoRepositorySuite) TestRandomTodoNotRow_Negative() {
	todo, err := suite.repository.GetRandomaTodo(context.Background(), 12123)
	suite.ErrorContains(err, fmt.Sprint(sql.ErrNoRows))
	suite.Empty(todo)
}

func (suite *todoRepositorySuite) TestGetSomeTodo_Positive() {
	userId := suite.addManyTodo()

	arg := repository.GetSomeTodosParams{
		Userid: userId,
		Limit:  55,
		Offset: 10,
	}

	todos, err := suite.repository.GetSomeTodos(context.Background(), arg)
	suite.NoError(err)
	suite.NotEmpty(todos)

	suite.Equal(len(todos), 55)

	for _, todo := range todos {
		suite.NotEmpty(todo)
		suite.Equal(todo.Userid, userId)
	}

}

func (suite *todoRepositorySuite) TestUpdateStatusTodo_Positive() {
	userId, id := suite.addOneTodo()

	arg := repository.UpdateStatusComplateParams{
		ID:        id,
		Complated: true,
		Userid:    userId,
	}

	todo, err := suite.repository.UpdateStatusComplate(context.Background(), arg)
	suite.NoError(err)
	suite.NotEmpty(todo)

	suite.True(todo.Complated)
	suite.Equal(todo.ID, id)
	suite.Equal(todo.Userid, userId)
}

func TestTodoRepository(t *testing.T) {
	suite.Run(t, new(todoRepositorySuite))
}
