package conf

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DbHost     string
	DbPort     string
	DbName     string
	DbUser     string
	DbPassword string
	JwtSecret  string
	Env        string
}

func getEnvVariable(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}
	return os.Getenv(key)
}

func NewConfig() Config {
	var (
		dbHost     = getEnvVariable("DB_HOST")
		dbPort     = getEnvVariable("DB_PORT")
		dbUser     = getEnvVariable("DB_USER")
		dbName     = getEnvVariable("DB_NAME")
		dbPassword = getEnvVariable("DB_PASSWORD")
		jwtSecret  = getEnvVariable("JWT_SECRET")
	)

	return Config{
		DbHost:     dbHost,
		DbPort:     dbPort,
		DbName:     dbName,
		DbUser:     dbUser,
		DbPassword: dbPassword,
		JwtSecret:  jwtSecret,
	}
}
