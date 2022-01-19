package models

import "gorm.io/gorm"

type Event struct {
	gorm.Model
	Organizer   *User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	OrganizerID uint  `gorm:"index:,option:CONCURRENTLY"`
	CoverImage  string
	Title       string
	SubTitle    string
	Content     string
	ScheduledOn string
	//Jobs        []*EventJob
}

type VolunteerEventTable struct {
	gorm.Model
	Volunteer   *User
	VolunteerID int
	Event       *Event
	EventID     int
}

type EventJob struct {
	gorm.Model
	Event                 *Event
	EventID               int
	VolunteerEventTable   *VolunteerEventTable
	VolunteerEventTableID int
	Duty                  string
}
