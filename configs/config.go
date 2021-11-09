package configs

import (
	"log"
	"os"
	"projectstructuring/models"

	"github.com/alexsasharegan/dotenv"
)

func SetENV() models.Config {
	err := dotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	host := os.Getenv("POSTGRES_HOST")
	username := os.Getenv("POSTGRES_USERNAME")
	password := os.Getenv("POSTGRES_PASSWORD")
	port := os.Getenv("POSTGRES_PORT")
	database := os.Getenv("POSTGRES_DATABASE")

	return models.Config{
		Host:     host,
		Username: username,
		Password: password,
		Port:     port,
		Database: database,
	}
}
