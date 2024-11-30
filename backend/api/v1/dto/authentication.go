package dto

import "stageflow/api/v1/models"

type SignUpRequestDTO struct {
	FirstName string `json:"firstName" binding:"required,min=2,max=50"`
	LastName  string `json:"lastName" binding:"required,min=2,max=50"`
	Email     string `json:"email" binding:"required,email"`
	Password  string `json:"password" binding:"required,min=8"`
}

type SignInRequestDTO struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type UserResponse struct {
	FirstName string `json:"firstName" binding:"required,min=2,max=50"`
	LastName  string `json:"lastName" binding:"required,min=2,max=50"`
	Email     string `json:"email" binding:"required,email"`
	JWT       string `json:"jwt" binding:"required"`
}

func NewUserResponse(u *models.User, JWT string) *UserResponse {
	return &UserResponse{
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Email:     u.Email,
		JWT:       JWT,
	}
}
