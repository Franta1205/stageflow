package auth

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(passwordHash), nil
}

func ComparePassword(userPassword string, requestPassword string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(userPassword), []byte(requestPassword)); err != nil {
		return errors.New("invalid password")
	}
	return nil
}
