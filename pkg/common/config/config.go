package config

import (
	"os"
)

type Config struct {
	Port       string
	DbLogin    string
	DbPassword string
	DbHost     string
	DbPort     string
	DbName     string
}

func GetConfig() *Config {
	return &Config{
		Port:       getEnv("PORT", ":3000"),
		DbLogin:    getEnv("DB_LOGIN", "postgres"),
		DbPassword: getEnv("DB_PASSWORD", "pswd1234"),
		DbHost:     getEnv("DB_HOST", "localhost"),
		DbPort:     getEnv("DB_PORT", "5432"),
		DbName:     getEnv("DB_NAME", "postgres"),
	}
}

func getEnv(key string, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultValue
}
