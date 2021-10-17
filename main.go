package main

import (
	"github.com/Bakhtiyar-Garashov/quote-service/config"
	"github.com/Bakhtiyar-Garashov/quote-service/controllers"
	"github.com/Bakhtiyar-Garashov/quote-service/utils"
	"github.com/gin-gonic/gin"

	cors "github.com/rs/cors/wrapper/gin"
)

func main() {
	router := gin.Default()
	router.Use(cors.Default())

	// append 'api/v1/' to the all routes of v1 group
	v1 := router.Group("/api/v1")
	{
		v1.GET("/healthcheck", utils.HealthCheck)
		v1.POST("/users", controllers.CreateUser)
		v1.POST("/quote", controllers.CreateQuote)
	}

	// initialize the Db connection and run migrations
	config.NewPostgresqlDb().DB()
	router.Run(":5000")
}
