package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/leegeobuk/financial-ledger/api/resource"
)

// OK is a convenience function for 200 response.
func OK(c *gin.Context, data any) {
	c.JSON(http.StatusOK, data)
}

// NotFound is a convenience function for 404 response.
func NotFound(c *gin.Context, err error) {
	c.JSON(http.StatusNotFound, resource.ResErr{Error: err.Error()})
}

// Error is a convenience function for 500 response.
func Error(c *gin.Context, err error) {
	c.JSON(http.StatusInternalServerError, resource.ResErr{Error: err.Error()})
}
