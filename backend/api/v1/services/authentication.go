package services

import (
	"context"
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	"stageflow/api/v1/dto"
	"stageflow/api/v1/repository"
	"stageflow/config/initializers"
	"stageflow/pkg/auth"
)

type AuthService struct {
	UserRepo  *repository.UserRepository
	TokenRepo *repository.TokenRepository
}

func NewAuthService(ur *repository.UserRepository, tr *repository.TokenRepository) *AuthService {
	return &AuthService{
		UserRepo:  ur,
		TokenRepo: tr,
	}
}

func (s *AuthService) Register(signUpRequest *dto.SignUpRequestDTO) error {
	_, err := s.UserRepo.FindUserByEmail(signUpRequest.Email)

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

	user, err := s.UserRepo.CreateUser(signUpRequest)
	if err != nil {
		return err
	}

	fmt.Println("registering new user", user)
	return nil
}

func (s *AuthService) Login(requestDTO *dto.SignInRequestDTO) (*dto.UserResponse, error) {
	user, err := s.UserRepo.FindUserByEmail(requestDTO.Email)
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

	userResponse := dto.NewUserResponse(user, jwt)

	return userResponse, nil
}

func (s *AuthService) LogOut(ctx context.Context, userID string, jwt string) error {
	r := initializers.GetRedisClient()
	tokenRepository := repository.NewTokenRepository(r)
	err := tokenRepository.BlackListJWT(ctx, userID, jwt)
	if err != nil {
		return err
	}
	return nil
}
