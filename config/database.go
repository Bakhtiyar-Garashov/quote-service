package config

import (
	"fmt"
	"log"
	"os"

	"github.com/Bakhtiyar-Garashov/quote-service/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
)

type PostgresqlDb interface {
	DB() *gorm.DB
}

type postgresqlDb struct {
	db *gorm.DB
}

func NewPostgresqlDb() PostgresqlDb {
	var err error
	var conn *gorm.DB

	e := godotenv.Load()
	if e != nil {
		fmt.Print(e)
	}

	username := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")

	connectionString := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", dbHost, dbPort, username, dbName, password)

	conn, err = gorm.Open("postgres", connectionString)

	if err != nil {
		log.Println(fmt.Sprintf("Error connecting to database: %s", err))
	}

	conn.AutoMigrate(&models.User{}, &models.Quote{})

	return &postgresqlDb{
		db: conn,
	}

}

func (p *postgresqlDb) DB() *gorm.DB {
	return p.db
}
