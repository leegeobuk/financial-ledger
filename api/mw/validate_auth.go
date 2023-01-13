package mw

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/leegeobuk/household-ledger/api/resource"
	"github.com/leegeobuk/household-ledger/cfg"
	"github.com/leegeobuk/household-ledger/util/tokenutil"
)

// ValidateAccessToken validates access token for authorized access.
func ValidateAccessToken(c *gin.Context) {
	var reqHeader resource.ReqAccessToken
	if err := c.ShouldBindHeader(&reqHeader); err != nil {
		resource.Unauthorized(c, fmt.Errorf("no authorization header: %w", err))
		return
	}

	authorization := strings.Split(reqHeader.Authorization, " ")
	scheme, accessToken := authorization[0], authorization[1]
	if scheme != "Bearer" {
		resource.Unauthorized(c, fmt.Errorf("wrong authorization scheme: %s", scheme))
		return
	}

	publicKey := cfg.Env.Token.PublicKey
	if err := tokenutil.ValidateAccessToken(accessToken, publicKey); err != nil {
		resource.Unauthorized(c, fmt.Errorf("token validation failed: %w", err))
		return
	}

	c.Next()
}
