package repository

import (
	"go.uber.org/zap"
	"log"
	"maranatha_web/internal/models"
)

type UserQuery interface {
	GetUserByEmail(email string) (*models.User, error)
	UpdateVerifiedUserStatus(param string) error
	Create(user *models.User) error
	GetAllUsers() ([]models.User, error)
	UpdateUserImage(email, imageName string) error
}

type userQuery struct {
	repo postgresDBRepo
}

func (uq *userQuery) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	err := uq.repo.DB.Where("email = ?", email).First(&user).Error
	if err != nil {
		return &user, err
	}
	return &user, nil
}
func (uq *userQuery) UpdateVerifiedUserStatus(param string) error {
	err := uq.repo.DB.Table("users").Debug().Where("email = ?", param).Update("is_verified", true).Error
	if err != nil {
		log.Panicln(err)
	}
	return nil
}
func (uq *userQuery) Create(user *models.User) error {
	err := uq.repo.DB.Debug().Create(&user).Error
	if err != nil {
		uq.repo.App.ErrorLog.Error("error when trying to save user", zap.Any("error", err))
		return err
	}
	return nil
}
func (uq *userQuery) GetAllUsers() ([]models.User, error) {
	var users []models.User
	err := uq.repo.DB.Debug().Find(&users).Error
	if err != nil {
		uq.repo.App.ErrorLog.Error("error when trying to save user", zap.Any("error", err))
		return users, err
	}
	log.Println(users)
	return users, nil
}
func (uq *userQuery) UpdateUserImage(email, imageName string) error {
	err := uq.repo.DB.Table("users").Debug().Where("email = ?", email).Update("profile_image", imageName).Error
	if err != nil {
		log.Panicln(err)

	}
	return nil

}
