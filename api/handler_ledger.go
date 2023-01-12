package api

import (
	"github.com/gin-gonic/gin"
	"github.com/leegeobuk/household-ledger/api/resource"
)

//	@Tags			Basic
//	@Summary		Ledger
//	@Description	Add ledger
//	@Accept			json
//	@Produce		json
//	@Param			request	body		resource.ReqAddLedger	true	"Add ledger request body"
//	@Success		201		{object}	resource.ResAddLedger
//	@Failure		400		{object}	resource.ResErr
//	@Failure		500		{object}	resource.ResErr
//	@Router			/api/household-ledger/ledger [post]
func (s *Server) AddLedger(c *gin.Context) {
	req := c.MustGet("req").(resource.ReqAddLedger)

	ledgerID, err := s.db.InsertLedger(req.UserID, req.Desc, req.Date, req.Income)
	if err != nil {
		resource.Error(c, err)
		return
	}

	res := resource.ResAddLedger{
		LedgerID: ledgerID,
		UserID:   req.UserID,
	}

	resource.Created(c, res)
}
