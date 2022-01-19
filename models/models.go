package models

import (
	"mime/multipart"
	"time"
)

//Blogs hold the blog model information
type Blogs struct {
	User       User                  `json:"user"`
	CoverImage *multipart.FileHeader `json:"cover_image"`
	Images     []BlogImages          `json:"images"`
	Title      string                `json:"title"`
	SubTitle   string                `json:"sub_title"`
	Content    string                `json:"content"`
	CreatedAt  time.Time             `json:"created_at"`
	UpdatedAt  time.Time             `json:"updated_at"`
}

//BlogImages hold the images for the Blog model
type BlogImages struct {
	Image     *multipart.FileHeader `json:"image"`
	CreatedAt time.Time             `json:"created_at"`
	UpdatedAt time.Time             `json:"updated_at"`
}

//Books hold the Books model information

//Gallery hold the Gallery model information
type Gallery struct {
	Image     []GalleryImages `json:"images"`
	Caption   string          `caption:"caption"`
	CreatedAt time.Time       `json:"created_at"`
	UpdatedAt time.Time       `json:"updated_at"`
}

//GalleryImages hold the images for the Gallery model
type GalleryImages struct {
	Image     *multipart.FileHeader `json:"image"`
	CreatedAt time.Time             `json:"created_at"`
	UpdatedAt time.Time             `json:"updated_at"`
}

//Sermon holds information about the sermon model
