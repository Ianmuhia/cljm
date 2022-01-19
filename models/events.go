package models

import "gorm.io/gorm"

type ChurchEvent struct {
	gorm.Model
	Organizer      *User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	OrganizerID    uint  `gorm:"index:,option:CONCURRENTLY"`
	CoverImage     string
	Title          string
	SubTitle       string
	Content        string
	ScheduledOn    string
	ChurchJobsList []ChurchJob `gorm:"foreignKey:Id2;references:ChurchJobId2,constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type VolunteerEventTable struct {
	gorm.Model
	Volunteer   *User
	VolunteerID uint
	DataID      uint
}

type ChurchJob struct {
	//gorm.Model
	Id2         uint `gorm:"autoIncrement;primaryKey;type:int" json:"id"`
	Duty        string
	ChurchEvent ChurchEvent `gorm:"foreignKey:ID;references:ChurchEventID"`
}
