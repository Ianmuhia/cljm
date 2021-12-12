package models

import (
	"time"
)

//User holds the user model information
type Users struct {
	ID         int       `json:"id"`
	UserName   string    `json:"username"`
	Email      string    `json:"email"`
	Password   string    `json:"-"`
	Dp         string    `json:"dp"`
	AcessLevel int       `json:"acess_level"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

//News hold the news model information
type News struct {
	User       Users        `json:"user"`
	CoverImage string       `json:"cover_image"`
	Images     []NewsImages `json:"images"`
	Title      string       `json:"title"`
	SubTitle   string       `json:"sub_title"`
	Content    string       `json:"content"`
	CreatedAt  time.Time    `json:"created_at"`
	UpdatedAt  time.Time    `json:"updated_at"`
}

//NewsImages hold the images for the News model
type NewsImages struct {
	Image     string    `json:"image"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

//Blogs hold the blog model information
type Blogs struct {
	User       Users        `json:"user"`
	CoverImage string       `json:"cover_image"`
	Images     []BlogImages `json:"images"`
	Title      string       `json:"title"`
	SubTitle   string       `json:"sub_title"`
	Content    string       `json:"content"`
	CreatedAt  time.Time    `json:"created_at"`
	UpdatedAt  time.Time    `json:"updated_at"`
}

//BlogImages hold the images for the Blog model
type BlogImages struct {
	Image     string    `json:"image"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

//Books hold the Books model information
type Books struct {
	Title     string    `json:"title"`
	Synopsis  string    `json:"synopsis"`
	Author    string    `json:"author"`
	Gener     Geners    `json:"gener"`
	File      string    `json:"file"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

//Geners hold the Geners model information
type Geners struct {
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
	Image     string    `json:"image"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

//Sermon holds information about hte sermon model
type Sermon struct {
	Title     string    `json:"title"`
	SubTitle  string    `json:"sub_title"`
	Content   string    `json:"content"`
	VideoData string    `json:"video"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

