package services

import (
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	"stageflow/api/v1/dto"
	"stageflow/api/v1/models"
	"stageflow/api/v1/repository"
	"stageflow/pkg/auth"
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

	hashedPassword, err := auth.HashPassword(signUpRequest.Password)
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

func (s *AuthService) Login(requestDTO *dto.SignInRequestDTO) (*models.User, error) {
	userRepo := repository.NewUserRepository()

	user, err := userRepo.FindUserByEmail(requestDTO.Email)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, errors.New("user does not exist")
	}

	if err := auth.ComparePassword(user.Password, requestDTO.Password); err != nil {
		return nil, err
	}

	jwt, err := auth.GenerateJWT(user)
	if err != nil {
		return nil, err
	}

	tokenRepository := repository.NewTokenRepository()
	err = tokenRepository.SetUserJWT(user.ID, jwt)
	if err != nil {
		return nil, err
	}

	fmt.Println("this is jwt", jwt)

	return user, nil
}
