package test_test

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"testing"
	"todo-api/app"
	"todo-api/controller"
	"todo-api/model"
	"todo-api/repository"
	"todo-api/service"

	"github.com/go-playground/validator/v10"
	_ "github.com/lib/pq"
)

var testQueries *repository.Queries
var testDB *sql.DB
var serviceTodo service.TodoService
var hendler http.Handler

func delTable(db *sql.DB) {
	db.Exec("DELETE FROM todos")
	db.Exec("DELETE FROM users")
}

func TestMain(m *testing.M) {
	config, err := model.LoadConfig("../")
	if err != nil {
		log.Fatal("cannot load config: ", err)
	}

	testDB, err = sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}

	testQueries = repository.New(testDB)

	var validate = validator.New()
	serviceTodo = service.NewTodoService(testDB, validate)
	controllerTodo := controller.NewTodoController(serviceTodo)
	hendler = app.NewRouter(controllerTodo)

	os.Exit(m.Run())
}
