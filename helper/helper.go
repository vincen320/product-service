package helper

import "database/sql"

func CommitOrRollBack(tx *sql.Tx) {
	err := recover()
	if err != nil {
		errRollback := tx.Rollback()
		if errRollback != nil {
			panic("Rollback Error : " + errRollback.Error()) // 500 Internal Server Error
		}
		panic(err) //Tergantung nanti panicnya apa
	} else {
		errCommit := tx.Commit()
		if errCommit != nil {
			panic("Commit Error : " + errCommit.Error()) //500 Internal Server Error
		}
	}
}
