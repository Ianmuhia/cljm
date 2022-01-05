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
	// UpdateUser(isPartial bool, user users.User) (*users.User, *errors.RestErr)
	// DeleteUser(int64) *errors.RestErr
	// Search(string) (users.Users, *errors.RestErr)
	// LoginUser(users.LoginRequest) (*models.User, *errors.RestErr)
}

// func (s *usersService) GetUser(userId int64) (*users.User, *errors.RestErr) {
// 	result := &users.User{
// 		Id: userId,
// 	}
// 	if err := result.Get(); err != nil {
// 		return nil, err
// 	}
// 	return result, nil
// }

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

// func (s *usersService) UpdateUser(isPartial bool, user users.User) (*users.User, *errors.RestErr) {
// 	currentUser := &users.User{Id: user.Id}

// 	if err := currentUser.Get(); err != nil {
// 		return nil, err
// 	}
// 	if err := user.Validate(); err != nil {
// 		return nil, err
// 	}
// 	if isPartial {
// 		if user.FirstName != "" {
// 			currentUser.FirstName = user.FirstName
// 		}
// 		if user.LastName != "" {
// 			currentUser.LastName = user.LastName
// 		}
// 		if user.Email != "" {
// 			currentUser.Email = user.Email
// 		}
// 		if user.FirstName != "" {
// 			currentUser.FirstName = user.FirstName
// 		}
// 	} else {
// 		currentUser.FirstName = user.FirstName
// 		currentUser.LastName = user.LastName
// 		currentUser.Email = user.Email
// 	}
// 	if err := currentUser.Update(); err != nil {
// 		return nil, err
// 	}

// 	return currentUser, nil
// }

// func (s *usersService) DeleteUser(userId int64) *errors.RestErr {
// 	user := &users.User{
// 		Id: userId,
// 	}

// 	return user.Delete()
// }

// func (s *usersService) Search(status string) (users.Users, *errors.RestErr) {
// 	dao := &users.User{}
// 	return dao.Search(status)

// }

func (s *usersService) LoginUser(request users.LoginRequest) (*models.User, *errors.RestErr) {
	// dao := &users.User{
	// 	Email:    request.Email,
	// 	Password: crypto_utils.GetMd5(request.Password),
	// }
	// if err := dao.FindByEmailAndPassword(); err != nil {
	// 	return nil, err
	// }
	return &models.User{}, nil
}
