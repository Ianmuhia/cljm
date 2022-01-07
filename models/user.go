package models

type User struct {
	ID         int    `gorm:"PRIMARY_KEY;AUTO_INCREMENT;UNIQUE" json:"id"`
	UserName   string `gorm:"size:255;NOT NULL;UNIQUE" json:"username"`
	FullName   string `gorm:"size:255;NOT NULL;UNIQUE" json:"fullname"`
	Email      string `gorm:"size:100;NOT NULL;UNIQUE" json:"email"`
	Password   string `gorm:"size:100;NOT NULL;" json:"password"`
	IsVerified bool   `gorm:"default:false" json:"isverified" sql:"isverified"`
	CreatedAt  string `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt  string `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}
