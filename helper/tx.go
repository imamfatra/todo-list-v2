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
