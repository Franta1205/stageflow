package models

import "time"

type Organisation struct {
	ID        string    `json:"id" gorm:"type:varchar(36);primary_key"`
	Name      string    `json:"name" gorm:"type:varchar(255);not null"`
	Users     []*User   `gorm:"many2many:user_organisations"`
	CreatedAt time.Time `json:"createdAt" gorm:"type:timestamp(3);not null;default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"type:timestamp(3);not null"`
}
