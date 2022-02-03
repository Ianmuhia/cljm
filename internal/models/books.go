package models

import "gorm.io/gorm"

type Books struct {
	gorm.Model
	Title       string
	Synopsis    string
	CreatedBy   *User  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	CreatedByID uint   `gorm:"index:,option:CONCURRENTLY"`
	GenreID     uint   `gorm:"index:,option:CONCURRENTLY"`
	Genre       *Genre `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	File        string
}