package configs

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	Port        string
	DatabaseURL string
	Database    string
}

func LoadConfig() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	dbPass := os.Getenv("MONGO_PASS")
	dbUser := os.Getenv("MONGO_USER")
	dbPort := os.Getenv("MONGO_PORT")
	dbUrl := fmt.Sprintf(
		"mongodb://%s:%s@localhost:%s", dbUser, dbPass, dbPort)

	cfg := &Config{
		Port:        os.Getenv("SERVER_PORT"),
		Database:    os.Getenv("DATABASE"),
		DatabaseURL: dbUrl,
	}

	return cfg, nil
}
