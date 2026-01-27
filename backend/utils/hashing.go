package utils

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

var cost int = 12

func HashPassword(password string) (string, error) {
	rawBytes, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	if err != nil {
		return "", fmt.Errorf("Error Hashing Password! Error: %s", err)
	}

	return string(rawBytes), nil
}

func ComparePasswordAndHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
