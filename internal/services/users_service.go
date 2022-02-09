package services

import (
	"bytes"
	"context"
	"encoding/gob"
	"log"
	redis_db "maranatha_web/internal/datasources/redis"

	"maranatha_web/internal/models"
	"maranatha_web/internal/repository"
	"maranatha_web/internal/utils/crypto_utils"
)

type usersService struct {
	dao repository.DAO
}

type UsersService interface {
	GetUserByEmail(email string) (*models.User, error)
	CreateUser(models.User) error
	UpdateUserStatus(email string) error
	UpdateUserImage(email, imageName string) error
	GetAllUsers() ([]models.User, error)
	VerifyPasswordResetCode(key string) VerificationData
	UpdateUserDetails(id uint, userModel models.User) error
}

func NewUsersService(dao repository.DAO) UsersService {
	return &usersService{dao: dao}
}

func (us *usersService) CreateUser(user models.User) error {
	user.PasswordHash = crypto_utils.Hash(user.PasswordHash)
	if err := us.dao.NewUserQuery().Create(&user); err != nil {
		log.Println(user.PasswordHash)
		return err
	}
	return nil
}

func (us *usersService) GetUserByEmail(email string) (*models.User, error) {
	user, err := us.dao.NewUserQuery().GetUserByEmail(email)
	if err != nil {
		return user, err
	}
	return user, err
}

func (us *usersService) GetAllUsers() ([]models.User, error) {
	users, err := us.dao.NewUserQuery().GetAllUsers()
	if err != nil {
		return users, err
	}
	return users, nil
}

func (us *usersService) UpdateUserImage(email, imageName string) error {
	err := us.dao.NewUserQuery().UpdateUserImage(email, imageName)
	if err != nil {
		log.Println(err)
		return err
	}
	return err
}

func (us *usersService) UpdateUserStatus(email string) error {
	err := us.dao.NewUserQuery().UpdateVerifiedUserStatus(email)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (us *usersService) UpdateUserDetails(id uint, userModel models.User) error {
	err := us.dao.NewUserQuery().UpdateUser(id, userModel)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (us *usersService) VerifyPasswordResetCode(key string) VerificationData {
	var a VerificationData
	data := redis_db.RedisClient.Get(context.TODO(), key)
	cmdb, err := data.Bytes()
	if err != nil {
		log.Println(err)
		return a
	}
	b := bytes.NewReader(cmdb)
	if err := gob.NewDecoder(b).Decode(&a); err != nil {
		log.Println(err)
		return a
	}
	return a
}
