package models

import (
	"gorm.io/gorm"
)

type News struct {
	gorm.Model
	Author     User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	AuthorID   uint `gorm:"index:,option:CONCURRENTLY"`
	CoverImage string
	Title      string
	SubTitle   string
	Content    string
}

type NewsImages struct {
	ID        int
	Image     string `gorm:"image"`
	CreatedAt string `gorm:"created_at"`
	UpdatedAt string `gorm:"updated_at"`
}
