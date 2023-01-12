package api

import (
	"database/sql"
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/leegeobuk/household-ledger/api/resource"
)

//	@Tags			Ledger
//	@Summary		Get ledger
//	@Description	Gets ledger with {id} from db.
//	@Produce		json
//	@Param			id	path		string	true	"ledger id"
//	@Success		200	{object}	resource.ResGetLedger
//	@Failure		400	{object}	resource.ResErr
//	@Failure		404	{object}	resource.ResErr
//	@Failure		500	{object}	resource.ResErr
//	@Router			/api/household-ledger/ledger/{id} [get]
func (s *Server) GetLedger(c *gin.Context) {
	reqURI := c.MustGet("reqURI").(resource.ReqGetLedger)

	// query db
	ledger, err := s.db.FindLedger(reqURI.LedgerID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			resource.NotFound(c, err)
			return
		}

		resource.Error(c, err)
		return
	}

	res := resource.ResGetLedger{
		LedgerID: ledger.LedgerID,
		UserID:   ledger.UserID,
		Desc:     ledger.Desc,
		Income:   ledger.Income,
		Date:     ledger.Date,
	}

	resource.OK(c, res)
}

//	@Tags			Ledger
//	@Summary		Add ledger
//	@Description	Adds ledger to db.
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
