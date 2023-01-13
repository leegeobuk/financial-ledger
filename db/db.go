package db

import "time"

// DB is an interface for database servers.
type DB interface {
	Ping() error
	RetryPing(interval time.Duration, reps int) error
	Close() error
}
