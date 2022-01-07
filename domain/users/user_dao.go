package users

import (
	"log"

	postgresql_db "maranatha_web/datasources/postgresql"
	"maranatha_web/logger"
	"maranatha_web/models"
	"maranatha_web/utils/errors"

	"gorm.io/gorm"
)

func GetUserByEmail(param string) (*errors.RestErr, *models.User) {
	var user models.User
	err := postgresql_db.Client.Debug().Where("email = ?", param).First(&user).Scan(&user).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return errors.NewBadRequestError("User does not exit"), &user
	}
	return nil, &user
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

func Update() *errors.RestErr {

	return nil

}

func Search(status string) ([]models.User, *errors.RestErr) {

	//
	results := make([]models.User, 0)

	return results, nil
}
