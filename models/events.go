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

type VolunteerChurchJob struct {
	gorm.Model
	Volunteer   *User      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	VolunteerID uint       `gorm:"index:,option:CONCURRENTLY"`
	ChurchJob   *ChurchJob `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	ChurchJobID uint       `gorm:"index:,option:CONCURRENTLY"`
}

type ChurchJob struct {
	gorm.Model
	Duty          string
	ChurchEventID uint
}
