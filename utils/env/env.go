package env

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

const (
	local      = "local"
	staging    = "staging"
	production = "production"
)

func MustGet(key string) string {
	val := os.Getenv(key)
	if val == "" && key != "PORT" {
		panic("Env key missing " + key)
	}
	return val
}

func CheckDotEnv() {
	err := godotenv.Load()
	if err != nil && os.Getenv("ENV") == local {
		log.Println("Error loading .env file")
	}
}
