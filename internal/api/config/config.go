package config

import (
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	APIPort          string
	WorkerPort       string
	OrganizationName string
}

var config *Config

func Init() error {
	err := godotenv.Load()

	config = &Config{
		APIPort:          os.Getenv("API_PORT"),
		WorkerPort:       os.Getenv("WORKER_PORT"),
		OrganizationName: os.Getenv("ORGANIZATION_NAME"),
	}

	return err
}

func GetConfig() *Config {
	return config
}
