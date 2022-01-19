package models

import "gorm.io/gorm"

type Prayer struct {
	gorm.Model
	Author   User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	AuthorID uint `gorm:"index:,option:CONCURRENTLY"`
	Content  string
}
