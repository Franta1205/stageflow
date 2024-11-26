package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
	"stageflow/api/v1/dto"
	"stageflow/api/v1/services"
	"stageflow/config/initializers"
	"time"
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
	user, err := authService.Login(&authInput)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	generateToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  user.ID,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	jwtSecret := initializers.LoadEnvVariable("JWT_SECRET")
	token, err := generateToken.SignedString([]byte(jwtSecret))

	c.JSON(http.StatusOK, gin.H{"token": token})
}

func (a *AuthController) GetUser(c *gin.Context) {
	user, _ := c.Get("currentUser")
	c.JSON(200, gin.H{
		"user": user,
	})
}
