package models

import "gorm.io/gorm"

type Events struct {
	gorm.Model
	Author      User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	AuthorID    uint `gorm:"index:,option:CONCURRENTLY"`
	CoverImage  string
	Title       string
	SubTitle    string
	Content     string
	ScheduledOn string
	VolunteerId int
	Volunteer   Volunteers
}

type Volunteers struct {
	Volunteer  User
	Assignment string
}
