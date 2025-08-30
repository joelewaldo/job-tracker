package config

import (
	"os"
)

type Config struct {
	DatabaseURL string
	ServerPort  string
	LogLevel    string
}

func Load() *Config {
	cfg := &Config{
		DatabaseURL: getEnv("DATABASE_URL", "postgres://user:pass@localhost:4001/db?sslmode=disable"),
		ServerPort:  getEnv("SERVER_PORT", "3000"),
		LogLevel:    getEnv("LOG_LEVEL", "info"),
	}

	return cfg
}

func getEnv(key string, fallback string) string {
	if value, ok := os.LookupEnv(key); ok && value != "" {
		return value
	}
	return fallback
}
