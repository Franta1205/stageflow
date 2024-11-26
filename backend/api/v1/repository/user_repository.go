package repository

import (
	"github.com/google/uuid"
	"stageflow/api/v1/dto"
	"stageflow/api/v1/models"
	"stageflow/config/initializers"
)

type UserRepository struct{}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (ur *UserRepository) Find(id string) (*models.User, error) {
	var user models.User
	db := initializers.GetDB()
	result := db.Where("id=?", id).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (ur *UserRepository) FindUserByEmail(email string) (*models.User, error) {
	var user models.User
	db := initializers.GetDB()
	result := db.Where("email=?", email).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (ur *UserRepository) CreateUser(signUpRequest *dto.SignUpRequestDTO) (*models.User, error) {
	db := initializers.GetDB()
	user := models.User{
		ID:        uuid.New().String(),
		FirstName: signUpRequest.FirstName,
		LastName:  signUpRequest.LastName,
		Email:     signUpRequest.Email,
		Password:  signUpRequest.Password,
	}
	result := db.Create(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}
