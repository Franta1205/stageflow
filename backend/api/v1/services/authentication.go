package services

import (
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
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

	user, err := userRepo.CreateUser(signUpRequest)
	if err != nil {
		return err
	}

	fmt.Println("registering new user", user)
	return nil
}
