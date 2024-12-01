package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"stageflow/api/v1/dto"
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
	err := oc.OrganisationService.Create(&organisationRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "org created"})
}
