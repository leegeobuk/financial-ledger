package db

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	_migrate "github.com/golang-migrate/migrate/v4"
	_mysql "github.com/golang-migrate/migrate/v4/database/mysql"
	"github.com/leegeobuk/financial-ledger/cfg"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
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

// Migrate migrates db tables if any change is detected.
func (mysql *MySQL) Migrate(path string) error {
	driver, err := _mysql.WithInstance(mysql.db, &_mysql.Config{})
	if err != nil {
		return fmt.Errorf("WithInstance: %w", err)
	}

	m, err := _migrate.NewWithDatabaseInstance(
		fmt.Sprintf("file://%s", path),
		cfg.Env.DB.Type,
		driver,
	)
	if err != nil {
		return fmt.Errorf("NewWithDatabaseInstance: %w", err)
	}

	if err = m.Up(); err != nil {
		if !errors.Is(err, _migrate.ErrNoChange) {
			return fmt.Errorf("up: %w", err)
		}
	}

	return nil
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
