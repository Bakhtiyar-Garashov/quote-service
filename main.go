package main

import (
	"github.com/Bakhtiyar-Garashov/quote-service/config"
	"github.com/Bakhtiyar-Garashov/quote-service/controllers"
	"github.com/Bakhtiyar-Garashov/quote-service/middlewares"
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
		v1.GET("/users", controllers.GetAllUsers)
		v1.POST("/users", controllers.CreateUser)
		v1.GET("/quotes", controllers.GetAllQuotes)
		v1.POST("/quotes", middlewares.RateLimiterMiddleware(), controllers.CreateQuote)

	}

	// initialize the Db connection
	config.NewPostgresqlDb().DB()
	// fire up the server
	router.Run(":5000")
}
