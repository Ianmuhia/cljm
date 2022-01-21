package models

import (
	"time"

	"gorm.io/gorm"
)

type ChurchEvent struct {
	gorm.Model
	Organizer   *User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	OrganizerID uint  `gorm:"index:,option:CONCURRENTLY"`
	CoverImage  string
	Title       string
	SubTitle    string
	Content     string
	ScheduledOn time.Time
	ChurchJobs  []*ChurchJob
}

type VolunteerEventTable struct {
	gorm.Model
	Volunteer   *User
	VolunteerID uint
	DataID      uint
}

type ChurchJob struct {
	gorm.Model
	Duty          string
	ChurchEventID uint
}
