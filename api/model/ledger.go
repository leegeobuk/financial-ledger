package model

import "time"

// Ledger represents ledger data.
type Ledger struct {
	LedgerID    string
	Description string
	Income      int
	Date        time.Time
	UserID      string
}
