package auth

import (
	"github.com/golang-jwt/jwt/v4"
	"stageflow/api/v1/models"
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
