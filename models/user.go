package models

type User struct {
	ID        int    `gorm:"PRIMARY_KEY;AUTO_INCREMENT;UNIQUE" json:"id"`
	UserName  string `gorm:"size:255;NOT NULL;UNIQUE" json:"username"`
	FullName  string `gorm:"size:255;NOT NULL;UNIQUE" json:"fullname"`
	Email     string `gorm:"size:100;NOT NULL;UNIQUE" json:"email"`
	Password  string `gorm:"size:100;NOT NULL;" json:"password"`
	CreatedAt string `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt string `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}
type createUserRequest struct {
	Username string `json:"username" binding:"required,alphanum"`
	Password string `json:"password" binding:"required,min=6"`
	FullName string `json:"full_name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
}