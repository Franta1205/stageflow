package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"stageflow/api/v1/dto"
	"stageflow/api/v1/models"
	"stageflow/api/v1/services"
)

type OrganisationController struct {
	OrganisationService *services.OrganisationService
}

func NewOrganisationController(s *services.OrganisationService) *OrganisationController {
	return &OrganisationController{
		OrganisationService: s,
	}
}

func (oc *OrganisationController) Create(c *gin.Context) {
	var organisationRequest dto.OrganisationRequest
	if err := c.ShouldBindJSON(&organisationRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	currentUser, exists := c.Get("currentUser")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found in context"})
		return
	}
	user, ok := currentUser.(*models.User)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid user type in context"})
		return
	}
	err := oc.OrganisationService.Create(&organisationRequest, user.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "org created"})
}

func (oc *OrganisationController) Update(c *gin.Context) {
	var organisationRequest dto.OrganisationRequest

	if err := c.ShouldBindJSON(&organisationRequest); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	if err := oc.OrganisationService.Update(&organisationRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"err": "org updated"})
}
