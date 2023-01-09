package db

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/leegeobuk/financial-ledger/cfg"
)

// MySQL communicates with MySQL server.
type MySQL struct {
	db *sql.DB
}

// NewMySQL returns new MySQL.
func NewMySQL(dsn string) (*MySQL, error) {
	db, err := sql.Open(cfg.Env.DB.Type, dsn)
	if err != nil {
		return nil, fmt.Errorf("new MySQL: %w", err)
	}

	// It is recommended to set MaxLifetime less than 5 minutes
	// to ensure MySQL driver closes the connection rather than
	// MySQL server, OS or middlewares. Check below site for more details.
	// https://github.com/go-sql-driver/mysql
	db.SetConnMaxLifetime(time.Minute * 3)
	// todo: Need to decide how many connections should be open
	// db.SetMaxOpenConns(10)
	// db.SetMaxIdleConns(10)

	return &MySQL{db: db}, nil
}

// Ping verifies the db connection.
func (mysql *MySQL) Ping() error {
	if err := mysql.db.Ping(); err != nil {
		return fmt.Errorf("ping MySQL: %w", err)
	}

	return nil
}

// RetryPing retries pinging db for given repetition on given interval basis.
func (mysql *MySQL) RetryPing(interval time.Duration, reps int) error {
	var err error
	for i := 0; i < reps; i++ {
		<-time.After(interval)
		if err = mysql.Ping(); err == nil {
			return nil
		}
	}

	return err
}

// Close closes the db connection.
func (mysql *MySQL) Close() error {
	if err := mysql.db.Close(); err != nil {
		return fmt.Errorf("close db: %w", err)
	}

	return nil
}
