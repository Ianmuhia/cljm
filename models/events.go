package models

import "gorm.io/gorm"

type ChurchEvent struct {
	gorm.Model
	Organizer   *User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	OrganizerID uint  `gorm:"index:,option:CONCURRENTLY"`
	CoverImage  string
	Title       string
	SubTitle    string
	Content     string
	ScheduledOn string
	ChurchJobs  []ChurchJob `gorm:"foreignKey:Id;"`
}

type VolunteerEventTable struct {
	gorm.Model
	Volunteer   *User
	VolunteerID uint
	DataID      uint
}

type ChurchJob struct {
	//gorm.Model
	Id            uint `gorm:"autoIncrement;primaryKey;type:int" json:"id"`
	Duty          string
	ChurchEventID int
	ChurchEvent   ChurchEvent `gorm:"foreignKey:ID;references:ChurchEventID"`
}
