package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type OrganisationController struct{}

func NewOrganisationController() *OrganisationController {
	return &OrganisationController{}
}

func (oc *OrganisationController) Create(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "org created"})
}
