package users

import (
	"log"

	postgresql_db "maranatha_web/datasources/postgresql"
	"maranatha_web/logger"
	"maranatha_web/models"
	"maranatha_web/utils/errors"
)

func GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	err := postgresql_db.Client.Where("email = ?", email).First(&user).Error
	if err != nil {
		return &user, err
	}
	log.Println(&user)
	return &user, nil
}
func UpdateVerifiedUserStatus(param string) *errors.RestErr {
	err := postgresql_db.Client.Table("users").Debug().Where("email = ?", param).Update("is_verified", true).Error
	if err != nil {
		log.Panicln(err)
	}
	return nil
}
func Create(user *models.User) *errors.RestErr {
	err := postgresql_db.Client.Debug().Create(&user).Error
	if err != nil {
		logger.Error("error when trying to save user", err)
		return errors.NewInternalServerError("database error")
	}
	return nil
}
func GetAllUsers() ([]models.User, error) {
	var users []models.User
	err := postgresql_db.Client.Debug().Find(&users).Error
	if err != nil {
		logger.Error("error when trying to save user", err)
		return users, err
	}
	log.Println(users)
	return users, nil
}
func UpdateUserImage(email, imageName string) error {
	err := postgresql_db.Client.Table("users").Debug().Where("email = ?", email).Update("profile_image", imageName).Error
	if err != nil {
		log.Panicln(err)

	}
	return nil

}
