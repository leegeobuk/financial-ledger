package resource

// ReqGetLedger represents request for getting ledger.
type ReqGetLedger struct {
	LedgerID uint `uri:"id" binding:"min=1"`
}

// ReqAddLedger represents request body for adding ledger.
type ReqAddLedger struct {
	UserID string `json:"user_id" binding:"required"`
	Desc   string `json:"description" binding:"required"`
	Income int    `json:"income" binding:"required"`
	Date   string `json:"date" binding:"required,datetime=2006-01-02"`
}
