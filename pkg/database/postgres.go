package database

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	conf "github.com/whitenight1201/go-devconnector/pkg/config"
	"github.com/whitenight1201/go-devconnector/pkg/models"
)

var DB *gorm.DB
var err error

func DatabaseConnection(cfg conf.Config) {
	fmt.Println("Connecting to Database...")

	conn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		cfg.DbHost,
		cfg.DbPort,
		cfg.DbUser,
		cfg.DbName,
		cfg.DbPassword,
	)

	DB, err = gorm.Open("postgres", conn)

	if err != nil {
		log.Fatal("Error connecting to the database...", err)
	}

	DB.AutoMigrate(models.User{})

	fmt.Println("Database connection successful...")
}
