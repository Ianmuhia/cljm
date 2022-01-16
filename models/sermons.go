package models

import "gorm.io/gorm"

type Sermon struct {
	gorm.Model
	Title    string `json:"title"`
	Url      string `json:"url"`
	DatePub  string `json:"date_pub"`
	Duration string `json:"duration"`
}
