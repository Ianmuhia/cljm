package models

import "gorm.io/gorm"

type Podcast struct {
	gorm.Model
	Author      User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	AuthorID    uint `gorm:"index:,option:CONCURRENTLY"`
	Title       string
	SubTitle    string
	Cast        string
	CoverImage  string
	Description string
	PodcastUrl  string
}
