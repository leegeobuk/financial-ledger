package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (a *API) Echo(c *gin.Context) {
	c.JSON(http.StatusOK, "echo")
}
