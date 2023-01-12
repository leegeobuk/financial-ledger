package api

import (
	"github.com/gin-gonic/gin"
	"github.com/leegeobuk/household-ledger/api/resource"
)

//	@Tags			Basic
//	@Summary		Ping
//	@Description	Ping
//	@Produce		json
//	@Success		200	{string}	pong
//	@Failure		500	{object}	resource.ResErr
//	@Router			/api/household-ledger/ping [get]
func (s *Server) Ping(c *gin.Context) {
	resource.OK(c, "pong")
}
