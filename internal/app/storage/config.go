package storage

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	DBUrl string `toml:"database_url"`
}

func NewConfig() *Config {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal(err)
	}

	return &Config{
		DBUrl: fmt.Sprintf(
			"host=%s dbname=%s user=%s password=%s sslmode=disable",
			os.Getenv("DB_HOST"),
			os.Getenv("DB_NAME"),
			os.Getenv("DB_USER"),
			os.Getenv("DB_PASSWORD"),
		),
	}
}
