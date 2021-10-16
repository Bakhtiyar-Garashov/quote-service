package config

import (
	"fmt"
	"os"

	"github.com/Bakhtiyar-Garashov/quote-service/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
)

type PostgresqlDb interface {
	GetDB() *gorm.DB
}

type postgresqlDb struct {
	db *gorm.DB
}

var db *gorm.DB

func init() {
	openDBConnection()
}

func GetDB() *gorm.DB {
	// Make sure we are connected to a database
	if err := db.DB().Ping(); err != nil {
		openDBConnection()
	}
	return db
}

func openDBConnection() interface{} {
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

	conn, err := gorm.Open("postgres", connectionString)
	if err != nil {
		fmt.Println("Error connecting to database: ", err)
	}

	db = conn
	// TODO: migrate database
	db.Debug().AutoMigrate(&models.User{}, &models.Quote{})
	return nil
}
