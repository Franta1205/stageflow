package services

import (
	"fmt"
	"stageflow/api/v1/dto"
)

type OrganisationService struct{}

func NewOrganisationService() *OrganisationService {
	return &OrganisationService{}
}

func (s *OrganisationService) Create(organisationRequset *dto.OrganisationRequest) {
	fmt.Println("stuff", organisationRequset)
}
