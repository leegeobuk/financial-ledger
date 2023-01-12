package resource

// ReqAddLedger is a request for adding ledger.
type ReqAddLedger struct {
	UserID string `json:"user_id" binding:"required"`
	Desc   string `json:"description" binding:"required"`
	Income int    `json:"income" binding:"required"`
	Date   string `json:"date" binding:"required,datetime=2006-01-02"`
}
