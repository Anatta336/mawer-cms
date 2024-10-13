package config

import (
	"os"
)

type Config struct {
	ServerAddress string
}

func LoadConfig() (*Config, error) {
	return &Config{
		ServerAddress: getEnv("SERVER_ADDRESS", "netdev.discovering-mawer.co.uk:8089"),
	}, nil
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
