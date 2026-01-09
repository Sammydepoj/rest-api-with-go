package util

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Println("Error hashing password:", err)
		return "", err
	}
	return string(hashedPassword), nil
}

func ComparePassword(storedPassword, providedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(storedPassword), []byte(providedPassword))
	if err != nil {
		return false
	}
	return true
}
