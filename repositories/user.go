package repositories

import (
	"github.com/Bakhtiyar-Garashov/quote-service/config"
	"github.com/Bakhtiyar-Garashov/quote-service/models"
)

type UserRepositoryInterface interface {
	GetAll() []models.User
	GetById(id uint) models.User
	Save(user models.User) models.User
	Delete(user models.User)
}

type userRepository struct {
	DB config.PostgresqlDb
}

func NewUserRepository(DB config.PostgresqlDb) UserRepositoryInterface {
	return &userRepository{
		DB: DB,
	}
}

func (u *userRepository) GetAll() []models.User {
	var users []models.User
	u.DB.DB().Table("users").Preload("Quotes").Find(&users)

	return users
}

func (u *userRepository) GetById(id uint) models.User {
	var user models.User
	u.DB.DB().First(&user, id)

	return user
}

func (u *userRepository) Save(user models.User) models.User {
	u.DB.DB().Save(&user)

	return user
}

func (u *userRepository) Delete(user models.User) {
	u.DB.DB().Delete(&user)
}
