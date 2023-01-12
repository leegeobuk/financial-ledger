package resource

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
