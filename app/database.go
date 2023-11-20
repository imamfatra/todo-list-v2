package app

import (
	"database/sql"
	"log"
	"time"
	"todo-api/model"
)

func NewDB() *sql.DB {
	config, err := model.LoadConfig("./")
	if err != nil {
		log.Fatal("cannot load config: ", err)
	}

	db, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(60 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)

	log.Println("Success connect to database")

	return db
}
