package models

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

type Register struct {
	Name     string `json:"name" binding:"required,min=3"`
	Password string `json:"password" binding:"required,min=3"`
	Email    string `json:"email" binding:"required,min=3"`
}

type ResetPassword struct {
	ID              int    `json:"id"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
}

// Login struct
type Login struct {
	Password string `json:"password"`
	Email    string `json:"email"`
}

type CreateReset struct {
	Email string `json:"email"`
}

//HashPassword hashes user password
func HashPassword(user *Register) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		log.Fatal(err)
	}
	user.Password = string(bytes)
}

func CreateHashedPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		log.Fatal(err)
	}
	return string(bytes)
}

//CheckPasswordHash compares hash with password
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
