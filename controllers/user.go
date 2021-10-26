package controllers

import (
	"fmt"

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

	user := models.User{
		Name:  userData.Name,
		Email: userData.Email,
	}

	userRepository := repositories.NewUserRepository(config.NewPostgresqlDb())

	userByEmail := userRepository.GetByEmail(userData.Email)

	if userByEmail.ID != 0 {
		c.JSON(400, gin.H{
			"success": "false",
			"message": fmt.Sprintf("User with email %s already exists", userData.Email),
		})
		return
	}

	userRepository.Save(user)

	c.JSON(201, gin.H{
		"success": "true",
		"message": "User created",
	})
}

func GetAllUsers(c *gin.Context) {
	userRepository := repositories.NewUserRepository(config.NewPostgresqlDb())
	users := userRepository.GetAll()

	c.JSON(200, gin.H{
		"success": "true",
		"message": "All users",
		"data":    users,
	})
}
