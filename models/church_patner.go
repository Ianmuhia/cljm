package models

import "gorm.io/gorm"

type ChurchPartner struct {
	gorm.Model
	Name  string
	Image string
	Since string
}
