package resource

// ReqGetLedgers binds path variable for getting ledgers.
type ReqGetLedgers struct {
	UserID string `uri:"user_id" binding:"required"`
}

// ReqGetLedger binds path variable for getting a ledger.
type ReqGetLedger struct {
	LedgerID uint `uri:"ledger_id" binding:"min=1"`
}

// ReqAddLedger binds request body for adding ledger.
type ReqAddLedger struct {
	UserID string `json:"user_id" binding:"required"`
	Desc   string `json:"description" binding:"required"`
	Income int    `json:"income" binding:"required"`
	Date   string `json:"date" binding:"required,datetime=2006-01-02"`
}
