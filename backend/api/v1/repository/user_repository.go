package repository

import (
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"stageflow/api/v1/dto"
	"stageflow/api/v1/models"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		DB: db,
	}
}

func (ur *UserRepository) Find(id string) (*models.User, error) {
	var user models.User
	result := ur.DB.Where("id=?", id).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (ur *UserRepository) FindUserByEmail(email string) (*models.User, error) {
	var user models.User
	result := ur.DB.Where("email=?", email).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (ur *UserRepository) CreateUser(signUpRequest *dto.SignUpRequestDTO) (*models.User, error) {
	user := models.User{
		ID:        uuid.New().String(),
		FirstName: signUpRequest.FirstName,
		LastName:  signUpRequest.LastName,
		Email:     signUpRequest.Email,
		Password:  signUpRequest.Password,
	}
	result := ur.DB.Create(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}
