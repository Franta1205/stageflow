package models

import "time"

type User struct {
	ID        string    `json:"id" gorm:"type:varchar(36);primary_key"`
	FirstName string    `json:"firstName" gorm:"type:varchar(255);not null"`
	LastName  string    `json:"lastName" gorm:"type:varchar(255);not null"`
	Email     string    `json:"email" gorm:"type:varchar(255);not null;unique"`
	Password  string    `json:"-" gorm:"type:varchar(255);not null"`
	CreatedAt time.Time `json:"createdAt" gorm:"type:timestamp(3);not null;default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"type:timestamp(3);not null"`
}
