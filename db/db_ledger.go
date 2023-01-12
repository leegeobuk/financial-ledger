package db

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/leegeobuk/household-ledger/api/model"
)

// FindLedgers retrieves ledger data with given userID.
func (mysql *MySQL) FindLedgers(userID string) ([]*model.Ledger, error) {
	query := "SELECT ledger_id, user_id, ledger_desc, income, ledger_date FROM ledger WHERE user_id = ?"

	errStr := "find ledgers: %w"
	rows, err := mysql.db.Query(query, userID)
	if err != nil {
		return nil, fmt.Errorf(errStr, err)
	}
	defer rows.Close() //nolint:errcheck

	ledgers := make([]*model.Ledger, 0)
	for rows.Next() {
		l := &model.Ledger{}
		if err = rows.Scan(&l.LedgerID, &l.UserID, &l.Desc, &l.Income, &l.Date); err != nil {
			return nil, fmt.Errorf(errStr, err)
		}
		ledgers = append(ledgers, l)
	}

	if rows.Err() != nil {
		return nil, fmt.Errorf(errStr, rows.Err())
	}

	return ledgers, nil
}

// FindLedger retrieves ledger data with given ledgerID.
func (mysql *MySQL) FindLedger(ledgerID uint) (*model.Ledger, bool, error) {
	query := "SELECT ledger_id, user_id, ledger_desc, income, ledger_date FROM ledger WHERE ledger_id = ?"

	l := model.Ledger{}
	if err := mysql.db.QueryRow(query, ledgerID).Scan(
		&l.LedgerID, &l.UserID, &l.Desc, &l.Income, &l.Date,
	); err != nil {
		err = fmt.Errorf("find ledger: %w", err)
		if errors.Is(err, sql.ErrNoRows) {
			return nil, true, err
		}

		return nil, false, err
	}

	return &l, false, nil
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
