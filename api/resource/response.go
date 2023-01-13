package resource

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// ResErr returns error response.
type ResErr struct {
	Error string `json:"error"`
}

// OK is a convenience function for 200 response.
func OK(c *gin.Context, data any) {
	c.JSON(http.StatusOK, data)
}

// Created is a convenience function for 201 response.
func Created(c *gin.Context, data any) {
	c.JSON(http.StatusCreated, data)
}

// BadRequest is a convenience function for 400 response.
func BadRequest(c *gin.Context, err error) {
	c.AbortWithStatusJSON(http.StatusBadRequest, ResErr{Error: err.Error()})
}

// NotFound is a convenience function for 404 response.
func NotFound(c *gin.Context, err error) {
	c.AbortWithStatusJSON(http.StatusNotFound, ResErr{Error: err.Error()})
}

// Unauthorized is a convenience function for 401 response.
func Unauthorized(c *gin.Context, err error) {
	c.AbortWithStatusJSON(http.StatusUnauthorized, ResErr{Error: err.Error()})
}

// Conflict is a convenience function for 409 response.
func Conflict(c *gin.Context, err error) {
	c.AbortWithStatusJSON(http.StatusConflict, ResErr{Error: err.Error()})
}

// Error is a convenience function for 500 response.
func Error(c *gin.Context, err error) {
	c.AbortWithStatusJSON(http.StatusInternalServerError, ResErr{Error: err.Error()})
}
