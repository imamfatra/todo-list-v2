package helper

import (
	"context"
	"database/sql"
	"fmt"
	"todo-api/repository"
)

func ExecTx(ctx context.Context, db *sql.DB, fn func(*repository.Queries) error) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}

	query := repository.New(tx)
	err = fn(query)
	if err != nil {
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			return fmt.Errorf("tx error: %v, rollback error: %v", err, rollbackErr)
		}
		return err
	}
	return tx.Commit()
}
