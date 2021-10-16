package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var HealthCheck = func(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "OK",
		"success": true,
	})
}
