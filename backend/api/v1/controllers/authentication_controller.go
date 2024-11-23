package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type AuthController struct{}

func NewAuthenticationController() *AuthController {
	return &AuthController{}
}

func (a *AuthController) CreateUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "user created"})
}
