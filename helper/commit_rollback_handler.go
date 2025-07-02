package helper

import (
	"database/sql"
	"fmt"
)

func CommitOrRollback(tx *sql.Tx) {
	err := recover()
	if err != nil {
		rollBackErr := tx.Rollback()
		if rollBackErr != nil {
			fmt.Printf("transaction rollback failed: %v\n", rollBackErr)
		}
	} else {
		commitErr := tx.Commit()
		if commitErr != nil {
			fmt.Printf("transaction commit failed: %v\n", commitErr)
		} else {
			fmt.Println("transaction committed successfully")
		}
	}
}
