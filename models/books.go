package models

import "gorm.io/gorm"

type Books struct {
	gorm.Model
	Title    string
	Synopsis string
	Author   string
	GenreID  int
	Genre    Genre
	File     string
}

//Genre hold the Genre model information
type Genre struct {
	gorm.Model
	Name string
}
