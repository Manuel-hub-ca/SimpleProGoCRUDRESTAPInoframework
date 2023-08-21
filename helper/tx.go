package helper

import "database/sql"

func CommitOrRollback(tx *sql.Tx){
	err := recover()

	if err != nil{
		errRolBack := tx.Rollback()
		PanicIfError(errRolBack)
		panic(err)
	} else {
		errCommit := tx.Commit()
		PanicIfError(errCommit)
	}
}