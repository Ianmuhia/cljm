package models

import (
	"time"

	"github.com/uptrace/bun"
)

type User struct {
	bun.BaseModel `bun:"table:users,alias:u"`
	ID            int64     `bun:"id,pk,autoincrement"`
	UserName      string    `bun:",notnull,unique"`
	FullName      string    `bun:",notnull,unique"`
	Email         string    `bun:",notnull,unique"`
	ProfileImage  string    `bun:",default:''"`
	PasswordHash  string    `bun:",notnull"`
	IsVerified    bool      `bun:",notnull,type:bool,default:'false'"`
	CreatedAt     time.Time `bun:",nullzero,notnull,default:current_timestamp"`
	UpdatedAt     time.Time `bun:",nullzero,notnull,default:current_timestamp"`
}
