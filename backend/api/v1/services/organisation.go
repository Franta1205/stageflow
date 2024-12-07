package services

import (
	"stageflow/api/v1/dto"
	"stageflow/api/v1/models"
	"stageflow/api/v1/repository"
)

type OrganisationService struct {
	OrganisationRepo *repository.OrganisationRepository
}

func NewOrganisationService(r *repository.OrganisationRepository) *OrganisationService {
	return &OrganisationService{
		OrganisationRepo: r,
	}
}

func (s *OrganisationService) Create(organisationRequest *dto.OrganisationRequest, userID string) (*models.Organisation, error) {
	organisation, err := s.OrganisationRepo.Create(organisationRequest, userID)
	if err != nil {
		return nil, err
	}
	return organisation, nil
}

func (s *OrganisationService) Update(organisationRequest *dto.OrganisationRequest) error {
	return nil
}
