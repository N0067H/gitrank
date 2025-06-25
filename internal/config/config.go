package config

import (
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	APIPort          string
	WorkerPort       string
	GithubToken      string
	OrganizationName string
	RedisHost        string
	RedisPort        string
	RedisPassword    string
	RedisDatabase    string
}

var config *Config

func Init() error {
	err := godotenv.Load()

	config = &Config{
		APIPort:          os.Getenv("API_PORT"),
		WorkerPort:       os.Getenv("WORKER_PORT"),
		GithubToken:      os.Getenv("GITHUB_TOKEN"),
		OrganizationName: os.Getenv("ORGANIZATION_NAME"),
		RedisHost:        os.Getenv("REDIS_HOST"),
		RedisPort:        os.Getenv("REDIS_PORT"),
		RedisPassword:    os.Getenv("REDIS_PASSWORD"),
		RedisDatabase:    os.Getenv("REDIS_DATABASE"),
	}

	return err
}

func GetConfig() *Config {
	return config
}
