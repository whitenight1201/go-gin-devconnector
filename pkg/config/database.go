package config

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
	"github.com/whitenight1201/go-devconnector/pkg/entity"
	"github.com/whitenight1201/go-devconnector/pkg/exception"
)

// creating new connection database
func DatabaseConnection() *gorm.DB {
	fmt.Println("Connecting to Database...")

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("host=%s user=%s password=%s port=%s dbname=%s sslmode=disable", dbHost, dbUsername, dbPassword, dbPort, dbName)

	db, err := gorm.Open("postgres", dsn)
	exception.PanicIfNeeded(err)

	db.AutoMigrate(&entity.User{})

	fmt.Println("Database connected successfully!")

	return db
}

// closing database connection
func CloseDatabaseConnection(db *gorm.DB) {
	db.Close()
}
