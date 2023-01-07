package api

import (
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/leegeobuk/financial-ledger/api/resource"
)

//	@Tags			Basic
//	@Summary		Time
//	@Description	Time
//	@Produce		json
//	@Success		200	{object}	resource.ResGetTime
//	@Failure		500	{object}	resource.ResErr
//	@Router			/api/ledger/time [get]
func (a *API) Time(c *gin.Context) {
	now := time.Now()
	timestamp := strconv.FormatInt(now.Unix(), 10)
	date := now.Format("2006-01-02 15:04")

	OK(c, resource.ResGetTime{
		Timestamp: timestamp,
		Date:      date,
	})
}
