package db

import (
	"fmt"

	"github.com/leegeobuk/household-ledger/api/model"
)

// FindLedger retrieves ledger data with given userID and date.
func (mysql *MySQL) FindLedger(ledgerID uint) (*model.Ledger, error) {
	query := "SELECT ledger_id, user_id, ledger_desc, income, ledger_date FROM ledger WHERE ledger_id = ?"

	l := model.Ledger{}
	if err := mysql.db.QueryRow(query, ledgerID).Scan(
		&l.LedgerID, &l.UserID, &l.Desc, &l.Income, &l.Date,
	); err != nil {
		return nil, fmt.Errorf("find ledger: %w", err)
	}

	return &l, nil
}

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
