package config

import (
	"os"
	"strconv"
)

type Config struct{
	DBDiver string
	DBHost string
	DBPort int
	DBUser string
	DBPassword string
	DBName string
	DBSSLMode string
}


func LoadConfig() *Config{
	return &Config{
		DBDiver : getEnv("DB_DRIVER", "postgres"),
		DBHost : getEnv("DB_HOST", "127.0.0.1"),
		DBPort: getEnvInt("DB_PORT", 5432),
		DBUser: getEnv("DB_USER", "postgres"),
		DBPassword: getEnv("DB_PASSWORD", "root"),
		DBName: getEnv("DB_NAME", "attempt2"),
		DBSSLMode: getEnv("DB_SSLMODE", "disable"),
	}
}

func getEnv(key, defaultValue string) string {
	if value, exist := os.LookupEnv(key); exist{
		return value
	}
	return defaultValue
}

func getEnvInt(key string, defaultValue int) int{
	if value, exist := os.LookupEnv(key); exist{
		valInt, err := strconv.Atoi(value)
		if err == nil {
			return valInt
		}
	}
	return defaultValue
}