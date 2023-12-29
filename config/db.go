package config

import (
	"fmt"
	"kulina-interview-test/domain/model"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var database *gorm.DB
var e error

func InitDB() *gorm.DB {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=Asia/Jakarta", host, port, username, password, dbName)

	database, e = gorm.Open(postgres.Open(dsn), &gorm.Config{SkipDefaultTransaction: true})

	if err != nil {
		log.Fatal("Failed to connect to database:", e)
	}

	if err != nil {
		panic("Can't connect to database")
	}

	var models = []interface{}{
		&model.User{},
		&model.Supplier{},
		&model.Product{},
		&model.Order{},
		&model.Address{},
		&model.Store{},
	}

	err = database.AutoMigrate(models...)
	if err != nil {
		log.Fatal(err)
	}

	return database
}
