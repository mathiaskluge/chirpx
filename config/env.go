package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	Port                   string
	DBPath                 string
	JWTSecret              string
	JWTExpirationInSeconds int
}

var Env = initConfig()

func initConfig() Config {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	return Config{
		Port:                   ":8080",
		DBPath:                 "db/db.json",
		JWTSecret:              getEnv("JWT_SECRET", "no secret..."),
		JWTExpirationInSeconds: getEnvAsInt("JWT_EXP", 3600*24), // 24h default
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); value != "" && ok {
		return value
	}

	return fallback
}

func getEnvAsInt(key string, fallback int) int {
	if value, ok := os.LookupEnv(key); value != "" && ok {
		i, err := strconv.Atoi(value)
		if err != nil {
			return fallback
		}

		return i
	}

	return fallback
}
