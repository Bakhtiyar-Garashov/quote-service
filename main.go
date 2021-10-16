package main

import (
	"github.com/Bakhtiyar-Garashov/quote-service/utils"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	// append api/v1/ to the all routes
	v1 := router.Group("/api/v1")

	v1.GET("/healthcheck", utils.HealthCheck)
	router.Run(":5000")
}
