package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type WebinarController struct{}

func NewWebinarController() *WebinarController {
	return &WebinarController{}
}

func (wc *WebinarController) Create(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "webinar created"})
}
