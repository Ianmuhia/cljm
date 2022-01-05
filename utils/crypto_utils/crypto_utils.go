package crypto_utils

import (
	"crypto/md5"
	"encoding/hex"
	"log"

	"golang.org/x/crypto/bcrypt"
)

func Hash(password string) string {
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Println(err)
	}
	converted := string(hashedPass)

	return converted
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func GetMd5(input string) string {
	hash := md5.New()
	defer hash.Reset()
	hash.Write([]byte(input))
	return hex.EncodeToString(hash.Sum(nil))
}
