package handler

import (
	"github.com/gin-gonic/gin"
)

//	@Tags			Basic
//	@Summary		Ping
//	@Description	Ping
//	@Produce		json
//	@Success		200	{string}	pong
//	@Failure		500	{object}	resource.ResErr
//	@Router			/api/ledger/ping [get]
func Ping(c *gin.Context) {
	OK(c, "pong")
}
