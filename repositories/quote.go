package repositories

import (
	"github.com/Bakhtiyar-Garashov/quote-service/config"
	"github.com/Bakhtiyar-Garashov/quote-service/models"
)

type QuoteRepositoryInterface interface {
	GetAll() []models.Quote
	GetById(id uint) models.Quote
	Save(quote models.Quote) models.Quote
	Delete(quote models.Quote)
}

type quoteRepository struct {
	DB config.PostgresqlDb
}

func NewProductRepostiory(DB config.PostgresqlDb) QuoteRepositoryInterface {
	return &quoteRepository{
		DB: DB,
	}
}

func (q *quoteRepository) GetAll() []models.Quote {
	var quotes []models.Quote
	q.DB.GetDB().Find(&quotes)

	return quotes
}

func (q *quoteRepository) GetById(id uint) models.Quote {
	var quote models.Quote
	q.DB.GetDB().First(&quote, id)

	return quote
}

func (q *quoteRepository) Save(quote models.Quote) models.Quote {
	q.DB.GetDB().Save(&quote)

	return quote
}

func (q *quoteRepository) Delete(quote models.Quote) {
	q.DB.GetDB().Delete(&quote)
}
