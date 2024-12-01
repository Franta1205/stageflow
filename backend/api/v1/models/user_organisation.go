package models

import "time"

type UserOrganisation struct {
	UserID         string    `json:"userId" gorm:"type:varchar(36);primary_key"`
	OrganizationID string    `json:"organizationId" gorm:"type:varchar(36);primary_key"`
	Role           string    `json:"role" gorm:"type:varchar(50);not null;default:'owner'"`
	CreatedAt      time.Time `json:"createdAt" gorm:"type:timestamp(3);not null;default:CURRENT_TIMESTAMP"`
	UpdatedAt      time.Time `json:"updatedAt" gorm:"type:timestamp(3);not null"`

	User         User         `json:"user" gorm:"foreignkey:UserID"`
	Organisation Organisation `json:"organization" gorm:"foreignkey:OrganizationID"`
}
