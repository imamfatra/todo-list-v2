package model

import (
	"database/sql"
	"fmt"
)

type TruncateTableExecutor struct {
	DB *sql.DB
}

func InitTruncateTableExecutor(db *sql.DB) TruncateTableExecutor {
	return TruncateTableExecutor{DB: db}
}

func (truncate *TruncateTableExecutor) TruncateTable(tableName []string) {

	tx, err := truncate.DB.Begin()
	if err != nil {
		panic(err)
	}
	for _, table := range tableName {
		query := fmt.Sprintf("TRUNCATE TABLE %s RESTART IDENTITY CASCADE;", table)
		_, err = tx.Exec(query)
		if err != nil {
			tx.Rollback()
			panic(err)
		}
	}
	tx.Commit()

}
