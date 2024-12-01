package services

import (
	"stageflow/api/v1/dto"
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

func (s *OrganisationService) Create(organisationRequest *dto.OrganisationRequest) error {
	_, err := s.OrganisationRepo.Create(organisationRequest)
	if err != nil {
		return err
	}
	return nil
}
