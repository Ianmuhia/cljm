package services

import (
	"fmt"
	"log"

	"maranatha_web/domain/users"
	"maranatha_web/models"
	"maranatha_web/utils/crypto_utils"
	"maranatha_web/utils/date_utils"
	"maranatha_web/utils/errors"
)

var (
	UsersService usersServiceInterface = &usersService{}
)

type usersService struct{}

type usersServiceInterface interface {
	GetUserByEmail(email string) (*models.User, *errors.RestErr)
	CreateUser(models.User) (*models.User, *errors.RestErr)
	UpdateUserStatus(email string) *errors.RestErr
}

func (s *usersService) CreateUser(user models.User) (*models.User, *errors.RestErr) {

	user.CreatedAt = date_utils.GetNowDBFormat()

	user.UpdatedAt = date_utils.GetNowDBFormat()

	user.Password = crypto_utils.Hash(user.Password)

	log.Println(user.Password)

	if err := users.Create(&user); err != nil {
		log.Println(user.Password)
		return nil, err
	}
	return &user, nil
}

func (s *usersService) GetUserByEmail(email string) (*models.User, *errors.RestErr) {
	err, user := users.GetUserByEmail(email)

	if err != nil {
		fmt.Println(err)
		return user, nil
	}
	log.Println(*user)
	return user, nil
}

func (s *usersService) UpdateUserStatus(email string) *errors.RestErr {
	err := users.UpdateVerifiedUserStatus(email)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
