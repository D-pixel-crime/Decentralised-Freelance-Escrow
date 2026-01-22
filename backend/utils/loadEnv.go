package utils

import (
	"log"

	"github.com/joho/godotenv"
)

// Load the
func LoadENV() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error Loading .env Files")
	}
}
