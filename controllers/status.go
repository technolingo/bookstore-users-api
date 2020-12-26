package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Ping checks for endpoint health
func Ping(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}
