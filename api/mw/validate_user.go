package mw

import (
	"github.com/gin-gonic/gin"
	"github.com/leegeobuk/household-ledger/api/resource"
)

// ValidateSignUp validates sign up request.
func ValidateSignUp(c *gin.Context) {
	var req resource.ReqSignUp
	if err := c.ShouldBindJSON(&req); err != nil {
		resource.BadRequest(c, err)
		return
	}

	c.Set("req", req)
	c.Next()
}
