package mw

import (
	"github.com/gin-gonic/gin"
	"github.com/leegeobuk/household-ledger/api/resource"
)

// ValidateAddLedger validates body of add ledger request.
func ValidateAddLedger(c *gin.Context) {
	var req resource.ReqAddLedger
	if err := c.ShouldBindJSON(&req); err != nil {
		resource.BadRequest(c, err)
		return
	}

	c.Set("req", req)
	c.Next()
}
