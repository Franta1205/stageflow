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

func (or *OrganisationRepository) Create(organisationRequest *dto.OrganisationRequest, userID string) (*models.Organisation, error) {
	tx := or.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	organisation := models.Organisation{
		ID:   uuid.New().String(),
		Name: organisationRequest.Name,
	}

	if err := tx.Create(&organisation).Error; err != nil {
		tx.Rollback()
		return nil, err
	}
	userOrg := models.UserOrganisation{
		UserID:         userID,
		OrganisationID: organisation.ID,
		Role:           "owner",
	}

	if err := tx.Create(&userOrg).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		return nil, err
	}

	return &organisation, nil
}
