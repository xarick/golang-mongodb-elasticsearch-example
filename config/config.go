package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Application struct {
	GinMode string
	RunPort string

	MongoURL     string
	ElasticURL   string
	ElasticIndex string
}

func LoadConfig() Application {
	cfg := Application{}

	err := godotenv.Load()
	if err != nil {
		log.Fatal(".env file not found")
	}

	cfg.GinMode = os.Getenv("GIN_MODE")
	cfg.RunPort = os.Getenv("RUN_PORT")

	cfg.MongoURL = os.Getenv("MONGO_URL")
	cfg.ElasticURL = os.Getenv("ELASTIC_URL")
	cfg.ElasticIndex = os.Getenv("ELASTIC_INDEX")

	return cfg
}
