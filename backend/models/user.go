package models

import "time"

type User struct {
	ID        uint   `json:"id" gorm:"primary_key"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email" gorm:"unique"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
