package mw

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/leegeobuk/household-ledger/api/resource"
)

// ValidateGetLedgers validates request of get ledgers.
func ValidateGetLedgers(c *gin.Context) {
	var reqURI resource.ReqGetLedgers
	if err := c.ShouldBindUri(&reqURI); err != nil {
		resource.BadRequest(c, err)
		return
	}

	c.Set("reqURI", reqURI)
	c.Next()
}

// ValidateGetLedger validates request of get ledger request.
func ValidateGetLedger(c *gin.Context) {
	var reqURI resource.ReqGetLedger
	if err := c.ShouldBindUri(&reqURI); err != nil {
		resource.BadRequest(c, fmt.Errorf("validation failed: %w", err))
		return
	}

	c.Set("reqURI", reqURI)
	c.Next()
}

// ValidateAddLedger validates body of add ledger request.
func ValidateAddLedger(c *gin.Context) {
	var req resource.ReqAddLedger
	if err := c.ShouldBindJSON(&req); err != nil {
		resource.BadRequest(c, fmt.Errorf("validation failed: %w", err))
		return
	}

	c.Set("req", req)
	c.Next()
}
