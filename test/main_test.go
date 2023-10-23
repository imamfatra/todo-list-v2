package test_test

import (
	"database/sql"
	"log"
	"os"
	"testing"
	"todo-api/model"
	"todo-api/repository"

	_ "github.com/lib/pq"
)

var testQueries *repository.Queries
var testDB *sql.DB

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

	os.Exit(m.Run())
}
