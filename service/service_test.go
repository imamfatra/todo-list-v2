package service_test

import (
	"context"
	"database/sql"
	"log"
	"testing"
	"todo-api/model"
	"todo-api/repository"
	"todo-api/service"

	"github.com/go-playground/validator/v10"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/suite"
)

type todoServiceSuite struct {
	suite.Suite
	repository      *repository.Queries
	service         service.TodoService
	cleanUpDatabase model.TruncateTableExecutor
}

func (suite *todoServiceSuite) SetupTest() {
	validate := validator.New()
	config, err := model.LoadConfig("../.")
	if err != nil {
		log.Fatal("cannot load config: ", err)
	}

	db, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}

	repository := repository.NewTodoRepository()
	service := service.NewTodoService(repository, validate, db)

	suite.repository = repository
	suite.service = service
	suite.cleanUpDatabase = model.InitTruncateTableExecutor(db)
}

func (suite *todoServiceSuite) TearDownTest() {
	defer suite.cleanUpDatabase.TruncateTable([]string{"users", "todos"})
}

func (suite *todoServiceSuite) registrasi_positive() model.RegistrasiRequest {
	arg := model.RegistrasiRequest{
		Email:    model.RandomMail(),
		Username: model.RandomString(10) + "12345",
		Password: "password111",
	}

	user := suite.service.Registrasi(context.Background(), arg)
	suite.Equal(arg.Email, user.Email)
	suite.Equal(arg.Username, user.Username)
	suite.NotZero(user.Userid)

	return arg
}

func (suite *todoServiceSuite) TestRegistrasi_Positive() {
	suite.registrasi_positive()
}

func (suite *todoServiceSuite) TestLogin_Positive() {
	user := suite.registrasi_positive()

	arg := model.LoginRequest{
		Username: user.Username,
		Password: user.Password,
	}

	result := suite.service.Login(context.Background(), arg)
	suite.Equal(result.Username, arg.Username)
	suite.Equal(result.Password, arg.Password)
	suite.Equal(result.Email, user.Email)
	suite.NotZero(result.Userid)
}

func (suite *todoServiceSuite) createOneTodo() repository.AddaNewTodoRow {
	user := suite.registrasi_positive()
	account := suite.service.Login(context.Background(), model.LoginRequest{
		Username: user.Username,
		Password: user.Password,
	})

	arg := model.AddNewTodoRequest{
		Todo:      model.RandomString(100),
		Complated: false,
		Userid:    account.Userid,
	}

	todo := suite.service.AddTodo(context.Background(), arg)
	suite.Equal(todo.Todo, arg.Todo)
	suite.Equal(todo.Complated, arg.Complated)
	suite.Equal(todo.Userid, account.Userid)
	suite.NotZero(todo.ID)

	return todo
}

func (suite *todoServiceSuite) createManyTodos() model.AddNewTodoRequest {
	user := suite.registrasi_positive()
	account := suite.service.Login(context.Background(), model.LoginRequest{
		Username: user.Username,
		Password: user.Password,
	})

	arg := model.AddNewTodoRequest{
		Todo:      model.RandomString(1443),
		Complated: false,
		Userid:    account.Userid,
	}

	for i := 0; i < 100; i++ {
		todo := suite.service.AddTodo(context.Background(), arg)
		suite.NotZero(todo.ID)
	}
	return arg
}

func (suite *todoServiceSuite) TestAddTodo() {
	todo := suite.createOneTodo()
	suite.Equal(todo.Complated, false)
	suite.NotZero(todo.ID)
	suite.NotZero(todo.Userid)
}

func (suite *todoServiceSuite) TestGetAllTodo_Positive() {
	user := suite.createManyTodos()
	arg := model.GetAllTodoRequest{
		Userid: user.Userid,
	}

	todos := suite.service.GetAllTodo(context.Background(), arg)
	todo := todos.Todos[0]
	suite.Equal(len(todos.Todos), 100)
	suite.Equal(todos.Total, int64(100))
	suite.Equal(todo.Userid, arg.Userid)
	suite.NotZero(todo.ID)
}

func (suite *todoServiceSuite) TestGetTodo_Positive() {
	user := suite.createOneTodo()

	arg := model.GetorDeleteTodoRequest{
		Userid: user.Userid,
		ID:     user.ID,
	}
	todo := suite.service.GetTodo(context.Background(), arg)
	suite.Equal(todo.ID, user.ID)
	suite.Equal(todo.Userid, user.Userid)
	suite.Equal(todo.Todo, user.Todo)
	suite.False(todo.Complated)
}

func (suite *todoServiceSuite) TestUpdateStatusTodo_Positive() {
	user := suite.createOneTodo()

	arg := model.UpdateStatusTodoRequest{
		ID:        user.ID,
		Complated: true,
		Userid:    user.Userid,
	}
	todo := suite.service.UpdateStatusTodo(context.Background(), arg)
	suite.Equal(todo.ID, user.ID)
	suite.Equal(todo.Userid, user.Userid)
	suite.Equal(todo.Todo, user.Todo)
	suite.NotEqual(todo.Complated, user.Complated)
}

func (suite *todoServiceSuite) TestDeleteTodo_Positive() {
	user := suite.createOneTodo()

	arg := model.GetorDeleteTodoRequest{
		Userid: user.Userid,
		ID:     user.ID,
	}
	todo := suite.service.DeleteTodo(context.Background(), arg)
	suite.Equal(todo.ID, user.ID)
	suite.Equal(todo.Userid, user.Userid)
	suite.Equal(todo.Todo, user.Todo)
	suite.NotZero(todo.Deletedon)
	suite.True(todo.Isdelete)
}

func (suite *todoServiceSuite) TestGetTodoRandom_Positive() {
	users := suite.createManyTodos()

	arg := model.GetAllTodoRequest{
		Userid: users.Userid,
	}
	todo := suite.service.GetRandomTodo(context.Background(), arg)
	suite.NotZero(todo.ID)
	suite.False(todo.Complated)
	suite.Equal(todo.Userid, users.Userid)
}

func (suite *todoServiceSuite) TestGetTodoFilter_Positive() {
	users := suite.createManyTodos()

	arg := model.GetTodoFilterRequest{
		Userid: users.Userid,
		Limit:  int32(20),
		Offset: int32(10),
	}
	todos := suite.service.GetTodoFilter(context.Background(), arg)
	suite.Equal(todos.Total, int32(20))
	suite.Equal(todos.Limit, arg.Limit)
	suite.Equal(todos.Skip, arg.Offset)
	suite.Equal(todos.Todos[1].Userid, users.Userid)
	suite.Equal(todos.Todos[0].ID, int32(11))

}

func TestTodoService(t *testing.T) {
	suite.Run(t, new(todoServiceSuite))
}
