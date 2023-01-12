package db

import "fmt"

// InsertLedger inserts ledger data and returns ledger_id.
func (mysql *MySQL) InsertLedger(userID, desc, date string, income int) (int, error) {
	query := "INSERT INTO ledger (user_id, ledger_desc, income, ledger_date) VALUES(?, ?, ?, ?)"

	result, err := mysql.db.Exec(query, userID, desc, income, date)
	if err != nil {
		return -1, fmt.Errorf("insert ledger: %w", err)
	}

	// mysql driver returns nil error
	ledgerID, _ := result.LastInsertId()

	return int(ledgerID), nil
}
