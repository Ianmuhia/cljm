package models

import (
	"mime/multipart"
	"time"
)

// Users  holds the user model information


//News hold the news model information
type News struct {
	User       User                `json:"user"`
	CoverImage *multipart.FileHeader `json:"cover_image"`
	Images     []NewsImages          `json:"images"`
	Title      string                `json:"title"`
	SubTitle   string                `json:"sub_title"`
	Content    string                `json:"content"`
	CreatedAt  time.Time             `json:"created_at"`
	UpdatedAt  time.Time             `json:"updated_at"`
}

//NewsImages hold the images for the News model
type NewsImages struct {
	Image     *multipart.FileHeader `json:"image"`
	CreatedAt time.Time             `json:"created_at"`
	UpdatedAt time.Time             `json:"updated_at"`
}

//Blogs hold the blog model information
type Blogs struct {
	User       User                `json:"user"`
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
type Books struct {
	Title     string                `json:"title"`
	Synopsis  string                `json:"synopsis"`
	Author    string                `json:"author"`
	Genre     Genre                 `json:"genre"`
	File      *multipart.FileHeader `json:"file"`
	CreatedAt time.Time             `json:"created_at"`
	UpdatedAt time.Time             `json:"updated_at"`
}

//Genre hold the Genre model information
type Genre struct {
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

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
type Sermon struct {
	Title     string    `json:"title"`
	SubTitle  string    `json:"sub_title"`
	Content   string    `json:"content"`
	VideoData string    `json:"video"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
