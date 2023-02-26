package config

import (
	"os"
)

type Config struct {
	Port  string
	DbUrl string
}

func GetConfig() *Config {
	return &Config{
		Port:  getEnv("PORT", ":3000"),
		DbUrl: getEnv("DB_URL", ""),
	}
}

func getEnv(key string, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultValue
}
