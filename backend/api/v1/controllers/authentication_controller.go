package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"stageflow/api/v1/dto"
	"stageflow/api/v1/services"
)

type AuthController struct{}

func NewAuthenticationController() *AuthController {
	return &AuthController{}
}

func (a *AuthController) CreateUser(c *gin.Context) {
	var authInput dto.SignUpRequestDTO

	if err := c.ShouldBindJSON(&authInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	authService := services.NewAuthService()
	if err := authService.Register(&authInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "user created"})
}

func (a *AuthController) Login(c *gin.Context) {
	var authInput dto.SignInRequestDTO

	if err := c.ShouldBindJSON(&authInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	authService := services.NewAuthService()
	if err := authService.Login(&authInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "successfully logged in"})
}
