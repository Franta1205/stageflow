package auth

import (
	"context"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"stageflow/api/v1/models"
	"stageflow/api/v1/repository"
	"stageflow/config/initializers"
	"time"
)

func GenerateJWT(u *models.User) (string, error) {
	generateToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  u.ID,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	jwtSecret := initializers.LoadEnvVariable("JWT_SECRET")
	token, err := generateToken.SignedString([]byte(jwtSecret))
	if err != nil {
		return "", err
	}
	return token, nil
}

// IsValid jwt in redis is stored as key and value is userID
func IsValid(ctx context.Context, jwt string) bool {
	tokenRepository := repository.NewTokenRepository()
	token, err := tokenRepository.FindJWT(ctx, jwt)
	fmt.Printf("IsValid check - Token: %v, Error: %v\n", token, err)

	if err != nil {
		fmt.Printf("Error in IsValid: %v\n", err)
		return false
	}

	return token == nil
}
