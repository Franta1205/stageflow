package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"stageflow/api/v1/dto"
)

type OrganisationController struct{}

func NewOrganisationController() *OrganisationController {
	return &OrganisationController{}
}

func (oc *OrganisationController) Create(c *gin.Context) {
	var organisationRequest dto.OrganisationRequest
	if err := c.ShouldBindJSON(&organisationRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "org created"})
}
