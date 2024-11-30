package controllers

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"stageflow/api/v1/dto"
	"stageflow/api/v1/models"
	"stageflow/api/v1/services"
	"strings"
	"time"
)

type AuthController struct {
	AuthService *services.AuthService
}

func NewAuthenticationController() *AuthController {
	return &AuthController{
		AuthService: services.NewAuthService(),
	}
}

func (a *AuthController) CreateUser(c *gin.Context) {
	var authInput dto.SignUpRequestDTO

	if err := c.ShouldBindJSON(&authInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := a.AuthService.Register(&authInput); err != nil {
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

	user, err := a.AuthService.Login(&authInput)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": user})
}

func (a *AuthController) LogOut(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()
	bearerToken := c.GetHeader("Authorization")
	token := strings.TrimPrefix(bearerToken, "Bearer ")

	userInterface, _ := c.Get("currentUser")
	user := userInterface.(*models.User)
	err := a.AuthService.LogOut(ctx, token, user.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "user logged out"})
}

func (a *AuthController) GetUser(c *gin.Context) {
	user, _ := c.Get("currentUser")
	c.JSON(200, gin.H{
		"user": user,
	})
}
