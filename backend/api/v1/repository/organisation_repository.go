package repository

import (
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"stageflow/api/v1/dto"
	"stageflow/api/v1/models"
)

type OrganisationRepository struct {
	DB *gorm.DB
}

func NewOrganisationRepository(db *gorm.DB) *OrganisationRepository {
	return &OrganisationRepository{
		DB: db,
	}
}

func (or *OrganisationRepository) Create(organisationRequest *dto.OrganisationRequest) (*models.Organisation, error) {
	organisation := models.Organisation{
		ID:   uuid.New().String(),
		Name: organisationRequest.Name,
	}
	result := or.DB.Create(&organisation)
	if result.Error != nil {
		return nil, result.Error
	}
	return &organisation, nil
}
