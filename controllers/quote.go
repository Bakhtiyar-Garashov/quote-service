package controllers

import (
	"github.com/Bakhtiyar-Garashov/quote-service/config"
	"github.com/Bakhtiyar-Garashov/quote-service/dto"
	"github.com/Bakhtiyar-Garashov/quote-service/models"
	"github.com/Bakhtiyar-Garashov/quote-service/repositories"
	"github.com/Bakhtiyar-Garashov/quote-service/utils"
	"github.com/gin-gonic/gin"
)

func CreateQuote(c *gin.Context) {
	newQuote := new(dto.QuoteRequest)
	c.Request.Close = true
	if err := c.ShouldBindJSON(newQuote); err != nil {
		c.JSON(400, gin.H{
			"success": "false",
			"message": err.Error(),
		})
		return
	}

	userRepository := repositories.NewUserRepository(config.NewPostgresqlDb())

	user := userRepository.GetById(newQuote.UserID)

	if user.ID == 0 {
		c.JSON(400, gin.H{
			"success": "false",
			"message": "This user doesn't exist",
		})
		return
	}

	quote := &models.Quote{
		CurrencySource:        newQuote.SourceCurrency,
		CurrencyTarget:        newQuote.TargetCurrency,
		Amount:                newQuote.Amount,
		Fee:                   utils.GenerateRandomFee(),
		EstimatedDeliveryTime: utils.GenerateRandomFutureDate(),
		UserID:                newQuote.UserID,
	}

	quoteRepository := repositories.NewQuoteRepository(config.NewPostgresqlDb())

	createdQuote := quoteRepository.Save(*quote)
	responseQuote := new(dto.QuoteResponse)

	responseQuote.ID = createdQuote.ID
	responseQuote.Fee = createdQuote.Fee
	responseQuote.EstimatedDeliveryTime = createdQuote.EstimatedDeliveryTime

	c.JSON(201, gin.H{
		"success": "true",
		"message": "Quote created",
		"data":    responseQuote,
	})
}

func GetAllQuotes(c *gin.Context) {
	quoteRepository := repositories.NewQuoteRepository(config.NewPostgresqlDb())
	quotes := quoteRepository.GetAll()

	c.JSON(200, gin.H{
		"success": "true",
		"message": "All quotes",
		"data":    quotes,
	})
}
