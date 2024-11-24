package services

import (
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"stageflow/api/v1/dto"
	"stageflow/api/v1/repository"
)

type AuthService struct{}

func NewAuthService() *AuthService {
	return &AuthService{}
}

func (s *AuthService) Register(signUpRequest *dto.SignUpRequestDTO) error {
	userRepo := repository.NewUserRepository()
	_, err := userRepo.FindUserByEmail(signUpRequest.Email)

	if err == nil {
		return errors.New("user already exists")
	}

	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	hashedPassword, err := hashPassword(signUpRequest.Password)
	if err != nil {
		return nil
	}

	signUpRequest.Password = hashedPassword

	user, err := userRepo.CreateUser(signUpRequest)
	if err != nil {
		return err
	}

	fmt.Println("registering new user", user)
	return nil
}

func hashPassword(password string) (string, error) {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(passwordHash), nil
}
