package resource

// ResAddLedger is a response for adding ledger.
type ResAddLedger struct {
	LedgerID int    `json:"ledger_id"`
	UserID   string `json:"user_id"`
}
