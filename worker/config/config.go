package config

import (
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	WorkerPort  string
	GithubToken string
}

var config *Config

func Init() error {
	err := godotenv.Load()

	config = &Config{
		WorkerPort:  os.Getenv("WORKER_PORT"),
		GithubToken: os.Getenv("GITHUB_TOKEN"),
	}
	
	return err
}

func GetConfig() *Config {
	return config
}
