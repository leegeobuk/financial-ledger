package api

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func (a *API) Time(c *gin.Context) {
	now := time.Now()
	timestamp := strconv.FormatInt(now.Unix(), 10)
	date := now.Format("2006-01-02 15:04")

	c.JSON(http.StatusOK, gin.H{
		"timestamp": timestamp,
		"date":      date,
	})
}
