package repository

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
	"log"
	"maranatha_web/internal/models"
)

type UserQuery interface {
	GetUserByEmail(email string) (*models.User, error)
	GetUserByID(id uint) (*models.User, error)
	UpdateVerifiedUserStatus(param string) error
	Create(user *models.User) error
	GetAllUsers() (int, []*models.User, error)
	UpdateUserImage(email, imageName string) error
	UpdateUser(id uint, userModel models.User) error
	UpdateUserPassword(id uint, newPasswd string) error
}

type userQuery struct {
	repo postgresDBRepo
}

func (uq *userQuery) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	err := uq.repo.DB.Where("email = ?", email).First(&user).Error
	if err != nil {
		uq.repo.App.ErrorLog.Error("error when trying to get user", zap.Any("error", err))
		return &user, err
	}
	return &user, nil
}

func (uq *userQuery) GetUserByID(id uint) (*models.User, error) {
	var user models.User
	err := uq.repo.DB.Where("id = ?", id).First(&user).Error
	if err != nil {
		uq.repo.App.ErrorLog.Error("error when trying to get user", zap.Any("error", err))
		return &user, err
	}
	return &user, nil
}

func (uq *userQuery) UpdateUser(id uint, userModel models.User) error {
	err := uq.repo.DB.Debug().Where("id = ?", id).Updates(&userModel).Error
	if err != nil || err == gorm.ErrRecordNotFound {
		uq.repo.App.ErrorLog.Error("error when trying to update user ", zap.Any("error", err))
		return err
	}
	return nil
}
func (uq *userQuery) UpdateUserPassword(id uint, newPasswd string) error {
	var userModel models.User
	err := uq.repo.DB.Debug().Model(&userModel).Where("id = ?", id).Update("password_hash", newPasswd).Error
	if err != nil || err == gorm.ErrRecordNotFound {
		uq.repo.App.ErrorLog.Error("error when trying to update user  password", zap.Any("error", err))
		return err
	}
	return nil
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
func (uq *userQuery) GetAllUsers() (int, []*models.User, error) {
	var users []*models.User
	var total int
	err := uq.repo.DB.Debug().Find(&users).Error
	if err != nil {
		uq.repo.App.ErrorLog.Error("error when trying to get users", zap.Any("error", err))
		return total, users, err
	}
	total = len(users)
	return total, users, nil
}
func (uq *userQuery) UpdateUserImage(email, imageName string) error {
	err := uq.repo.DB.Table("users").Debug().Where("email = ?", email).Update("profile_image", imageName).Error
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
