package database

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
	"github.com/whitenight1201/go-devconnector/pkg/models"
)

var DB *gorm.DB
var err error

func getEnvVariable(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}
	return os.Getenv(key)
}

func DatabaseConnection() {
	fmt.Println("Connecting to Database...")

	var (
		host     = getEnvVariable("DB_HOST")
		port     = getEnvVariable("DB_PORT")
		dbUser   = getEnvVariable("DB_USER")
		dbName   = getEnvVariable("DB_NAME")
		password = getEnvVariable("DB_PASSWORD")
	)

	conn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		host,
		port,
		dbUser,
		dbName,
		password,
	)

	DB, err = gorm.Open("postgres", conn)

	if err != nil {
		log.Fatal("Error connecting to the database...", err)
	}

	DB.AutoMigrate(models.User{})

	fmt.Println("Database connection successful...")
}
