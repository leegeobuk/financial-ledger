package mw

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/leegeobuk/household-ledger/api/resource"
)

// ValidateSignUp validates sign up request.
func ValidateSignUp(c *gin.Context) {
	var req resource.ReqSignUp
	if err := c.ShouldBindJSON(&req); err != nil {
		resource.BadRequest(c, fmt.Errorf("validate sign up: %w", err))
		return
	}

	c.Set("req", req)
	c.Next()
}

// ValidateSignIn validates sign in request.
func ValidateSignIn(c *gin.Context) {
	var req resource.ReqSignIn
	if err := c.ShouldBindJSON(&req); err != nil {
		resource.BadRequest(c, fmt.Errorf("validate sign in: %w", err))
		return
	}

	c.Set("req", req)
	c.Next()
}
