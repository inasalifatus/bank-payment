package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var PORT int

func LoadEnv(key string) string {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error loading .env file")
	}
	return os.Getenv(key)
}

func InitPort() {
	PORT, _ = strconv.Atoi(LoadEnv("PORT"))

}
