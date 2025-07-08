package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	APIPort          string
	WorkerPort       string
	AllowOrigins     string
	GithubToken      string
	OrganizationName string
	RedisHost        string
	RedisPort        string
	RedisPassword    string
	RedisDatabase    string
}

var AppConfig *Config

func Load() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("failed to load .env file")
	}

	AppConfig = &Config{
		APIPort:          os.Getenv("API_PORT"),
		WorkerPort:       os.Getenv("WORKER_PORT"),
		AllowOrigins:     os.Getenv("ALLOW_ORIGINS"),
		GithubToken:      os.Getenv("GITHUB_TOKEN"),
		OrganizationName: os.Getenv("ORGANIZATION_NAME"),
		RedisHost:        os.Getenv("REDIS_HOST"),
		RedisPort:        os.Getenv("REDIS_PORT"),
		RedisPassword:    os.Getenv("REDIS_PASSWORD"),
		RedisDatabase:    os.Getenv("REDIS_DATABASE"),
	}
}
