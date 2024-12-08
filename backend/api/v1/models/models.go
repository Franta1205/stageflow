package models

import "time"

type User struct {
	ID            string          `json:"id" gorm:"type:varchar(36);primary_key"`
	FirstName     string          `json:"firstName" gorm:"type:varchar(255);not null"`
	LastName      string          `json:"lastName" gorm:"type:varchar(255);not null"`
	Email         string          `json:"email" gorm:"type:varchar(255);not null;unique"`
	Password      string          `json:"-" gorm:"type:varchar(255);not null"`
	Organisations []*Organisation `gorm:"many2many:user_organisations"`
	CreatedAt     time.Time       `json:"createdAt" gorm:"type:timestamp(3);not null;default:CURRENT_TIMESTAMP"`
	UpdatedAt     time.Time       `json:"updatedAt" gorm:"type:timestamp(3);not null"`
}

type Organisation struct {
	ID        string     `json:"id" gorm:"type:varchar(36);primary_key"`
	Name      string     `json:"name" gorm:"type:varchar(255);not null"`
	Users     []*User    `gorm:"many2many:user_organisations"`
	Webinars  []*Webinar `json:"webinars" gorm:"foreignKey:OrganisationID"`
	CreatedAt time.Time  `json:"createdAt" gorm:"type:timestamp(3);not null;default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time  `json:"updatedAt" gorm:"type:timestamp(3);not null"`
}

type UserOrganisation struct {
	UserID         string    `json:"userId" gorm:"type:varchar(36);primary_key"`
	OrganisationID string    `json:"organisationId" gorm:"type:varchar(36);primary_key"`
	Role           string    `json:"role" gorm:"type:varchar(50);not null;default:'owner'"`
	CreatedAt      time.Time `json:"createdAt" gorm:"type:timestamp(3);not null;default:CURRENT_TIMESTAMP"`
	UpdatedAt      time.Time `json:"updatedAt" gorm:"type:timestamp(3);not null"`

	User         User         `json:"user" gorm:"foreignKey:UserID"`
	Organisation Organisation `json:"organization" gorm:"foreignKey:OrganizationID"`
}

type Webinar struct {
	ID             string          `json:"id" gorm:"type:varchar(36);primary_key"`
	Title          string          `json:"title" gorm:"type:varchar(255);not null"`
	ScheduledAt    string          `json:"scheduledAt" gorm:"type:timestamp(3);not null"`
	OrganisationID string          `json:"organisationId" gorm:"type:varchar(36);not null"`
	Organisation   *Organisation   `json:"organisation" gorm:"foreignKey:OrganisationID"`
	Participants   []*Participants `json:"participants" gorm:"foreignKey:ParticipantID"`
	CreatedAt      time.Time       `json:"createdAt" gorm:"type:timestamp(3);not null;default:CURRENT_TIMESTAMP"`
	UpdatedAt      time.Time       `json:"updatedAt" gorm:"type:timestamp(3);not null"`
}

type Participants struct {
	ID        string    `json:"id" gorm:"type:varchar(36);primary_key"`
	Email     string    `json:"email" gorm:"type:varchar(255);not null;"`
	FirstName string    `json:"firstName" gorm:"type:varchar(255);not null"`
	LastName  string    `json:"lastName" gorm:"type:varchar(255);not null"`
	WebinarID string    `json:"WebinarId" gorm:"type:varchar(36);not null"`
	Webinar   *Webinar  `json:"webinar" gorm:"foreignKey:WebinarID"`
	CreatedAt time.Time `json:"createdAt" gorm:"type:timestamp(3);not null;default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"type:timestamp(3);not null"`
}
