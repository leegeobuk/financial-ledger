package resource

// ResGetLedgers is a response for getting ledgers.
type ResGetLedgers struct {
	Ledgers []ResGetLedger `json:"ledgers"`
}

// ResGetLedger is a response for getting a ledger.
type ResGetLedger struct {
	LedgerID string `json:"ledger_id"`
	UserID   string `json:"user_id"`
	Desc     string `json:"description"`
	Income   int    `json:"income"`
	Date     string `json:"date"`
}

// ResAddLedger is a response for adding ledger.
type ResAddLedger struct {
	LedgerID int    `json:"ledger_id"`
	UserID   string `json:"user_id"`
}
