package api

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
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

	c.JSON(http.StatusOK, gin.H{
		"timestamp": timestamp,
		"date":      date,
	})
}
