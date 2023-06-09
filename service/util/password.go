package util

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// hashPassword returs the bcrypt hash of the password
func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("failed to hash password, %v", err)
	}
	return string(hashedPassword), nil
}

// CheckPassword checkd if the provide passwors is correct
func CheckPassword(password string, hashedPassowrd string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassowrd), []byte(password))
}
