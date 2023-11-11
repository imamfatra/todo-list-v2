package helper

import (
	"database/sql"
)

func CommitOrRollback(tx *sql.Tx) {
	err := recover()
	if err != nil {
		errRollback := tx.Rollback()
		IfError(errRollback)
		panic(err)
	} else {
		errCommit := tx.Commit()
		IfError(errCommit)
	}
}

// func ExecTx(ctx context.Context, db *sql.DB, fn func(repository.Querier) error) error {
// 	tx, err := db.Begin()
// 	if err != nil {
// 		return err
// 	}

// 	query := repository.New(tx)
// 	err = fn(query)
// 	if err != nil {
// 		if rollbackErr := tx.Rollback(); rollbackErr != nil {
// 			return fmt.Errorf("tx error: %v, rollback error: %v", err, rollbackErr)
// 		}
// 		return err
// 	}
// 	return tx.Commit()
// }
