package models

//type User struct {
//	bun.BaseModel `bun:"table:users,alias:u"`
//	ID            int64     `bun:"id,pk,autoincrement"`
//	UserName      string    `bun:",notnull,unique"`
//	FullName      string    `bun:",notnull,unique"`
//	Email         string    `bun:",notnull,unique"`
//	ProfileImage  string    `bun:",default:''"`
//	PasswordHash  string    `bun:",notnull"`
//	IsVerified    bool      `bun:",notnull,type:bool,default:'false'"`
//	CreatedAt     time.Time `bun:",nullzero,notnull,default:current_timestamp"`
//	UpdatedAt     time.Time `bun:",nullzero,notnull,default:current_timestamp"`
//}

type User struct {
	ID           int64  `gorm:"PRIMARY_KEY;AUTO_INCREMENT;UNIQUE" `
	UserName     string `gorm:"size:255;NOT NULL;UNIQUE" `
	FullName     string `gorm:"size:255;NOT NULL;UNIQUE" `
	Email        string `gorm:"size:100;NOT NULL;UNIQUE" `
	PasswordHash string `gorm:"size:100;NOT NULL;" `
	ProfileImage string `gorm:"default:''"`
	IsVerified   bool   `gorm:"default:false"  sql:"isverified"`
	CreatedAt    string `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt    string `gorm:"default:CURRENT_TIMESTAMP" `
}
