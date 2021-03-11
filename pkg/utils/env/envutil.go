package env

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Get is
func Get(key string) string {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	value := os.Getenv(key)

	return value
}
