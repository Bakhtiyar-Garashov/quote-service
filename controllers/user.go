package controllers

import (
	"fmt"
	"reflect"

	"github.com/Bakhtiyar-Garashov/quote-service/config"
	"github.com/Bakhtiyar-Garashov/quote-service/dto"
	"github.com/Bakhtiyar-Garashov/quote-service/models"
	"github.com/Bakhtiyar-Garashov/quote-service/repositories"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	userData := new(dto.UserCreate)

	if err := c.ShouldBindJSON(userData); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	fmt.Println(reflect.TypeOf(userData))

	user := &models.User{
		Name:  userData.Name,
		Email: userData.Email,
	}

	userRepository := repositories.NewUserRepository(config.NewPostgresqlDb())

	res := userRepository.Save(*user)

	fmt.Println(res)

	c.JSON(200, gin.H{
		"message": "pong",
	})
}
