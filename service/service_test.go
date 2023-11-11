package service_test

import (
	"context"
	"database/sql"
	"log"
	"testing"
	"todo-api/model"
	"todo-api/repository"
	"todo-api/service"

	"github.com/go-playground/validator"
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

	a := true
	arg := model.AddNewTodoRequest{
		Todo:      model.RandomString(100),
		Complated: a,
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
		Complated: true,
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
	suite.Equal(todo.Complated, true)
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

func TestTodoService(t *testing.T) {
	suite.Run(t, new(todoServiceSuite))
}
