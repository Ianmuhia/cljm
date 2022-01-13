package services

import (
	"log"

	"maranatha_web/domain/users"
	"maranatha_web/models"
	"maranatha_web/utils/crypto_utils"
	"maranatha_web/utils/errors"
)

var (
	UsersService usersServiceInterface = &usersService{}
)

type usersService struct{}

type usersServiceInterface interface {
	GetUserByEmail(email string) (*models.User, error)
	CreateUser(models.User) (*models.User, *errors.RestErr)
	UpdateUserStatus(email string) *errors.RestErr
	UpdateUserImage(email, imageName string) error
}

func (s *usersService) CreateUser(user models.User) (*models.User, *errors.RestErr) {
	user.PasswordHash = crypto_utils.Hash(user.PasswordHash)
	if err := users.Create(&user); err != nil {
		log.Println(user.PasswordHash)
		return nil, err
	}
	return &user, nil
}

func (s *usersService) GetUserByEmail(email string) (*models.User, error) {
	user, err := users.GetUserByEmail(email)
	if err != nil {
		return user, err
	}
	return user, err
}

func (s *usersService) GetAllUsers() error {
	err := users.GetAllUsers()
	if err != nil {
		return err
	}
	return nil
}

func (s *usersService) UpdateUserImage(email, imageName string) error {
	err := users.UpdateUserImage(email, imageName)
	if err != nil {
		log.Println(err)
		return err
	}
	return err
}

func (s *usersService) UpdateUserStatus(email string) *errors.RestErr {
	err := users.UpdateVerifiedUserStatus(email)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
