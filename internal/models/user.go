package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserName     string `gorm:"size:255;NOT NULL;UNIQUE" `
	FullName     string `gorm:"size:255;NOT NULL;UNIQUE" `
	Email        string `gorm:"size:100;NOT NULL;UNIQUE" `
	PasswordHash string `gorm:"size:100;NOT NULL;" `
	ProfileImage string `gorm:"default:''"`
	IsVerified   bool   `gorm:"default:false"  sql:"isverified"`
}
