package helper

import (
	"os"

	"golang.org/x/crypto/bcrypt"
	"github.com/joho/godotenv"

	"log"
)

func Getenv(key, fallback string) string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	valueEnv := os.Getenv(key)
	
	if len(valueEnv) > 0 {
		return valueEnv
	}
	return fallback
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
