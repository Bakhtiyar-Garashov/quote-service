package controllers

import (
	"github.com/Bakhtiyar-Garashov/quote-service/config"
	"github.com/Bakhtiyar-Garashov/quote-service/dto"
	"github.com/Bakhtiyar-Garashov/quote-service/models"
	"github.com/Bakhtiyar-Garashov/quote-service/repositories"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	userData := new(dto.UserCreate)

	if err := c.ShouldBindJSON(userData); err != nil {
		c.JSON(400, gin.H{
			"success": "false",
			"message": err.Error(),
		})
		return
	}

	user := &models.User{
		Name:  userData.Name,
		Email: userData.Email,
	}

	userRepository := repositories.NewUserRepository(config.NewPostgresqlDb())

	userRepository.Save(*user)

	c.JSON(201, gin.H{
		"success": "true",
		"message": "User created successfully",
	})
}
