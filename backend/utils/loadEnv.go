package utils

import (
	"github.com/charmbracelet/log"
	"github.com/joho/godotenv"
)

// Load the
func LoadENV() {
	err := godotenv.Load()
	if err != nil {
		log.Error("Error Loading .env Files", "err", err)
	}
}
