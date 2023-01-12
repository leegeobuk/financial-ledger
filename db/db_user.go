package db

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/leegeobuk/household-ledger/api/model"
)

// FindUserLogIn retrieves user with given userID.
func (mysql *MySQL) FindUserLogIn(userID string) (*model.UserLogIn, bool, error) {
	query := "SELECT user_id, passwd FROM user_login WHERE user_id = ?"

	u := model.UserLogIn{}
	if err := mysql.db.QueryRow(query, userID).Scan(&u.UserID, &u.Passwd); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, true, nil
		}

		return nil, false, fmt.Errorf("find userID: %w", err)
	}

	return &u, false, nil
}

// InsertUserLogIn inserts user login data to db.
func (mysql *MySQL) InsertUserLogIn(email, passwd string) error {
	query := "INSERT user_login (user_id, passwd) VALUES (?, ?)"

	_, err := mysql.db.Exec(query, email, passwd)
	if err != nil {
		return fmt.Errorf("insert user login: %w", err)
	}

	return nil
}

// InsertUserAccount inserts user account data to db.
func (mysql *MySQL) InsertUserAccount(email string) error {
	query := "INSERT user_account (user_id) VALUES (?)"

	_, err := mysql.db.Exec(query, email)
	if err != nil {
		return fmt.Errorf("insert user account: %w", err)
	}

	return nil
}
