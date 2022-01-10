package users

import (
	"context"
	"log"

	postgresql_db "maranatha_web/datasources/postgresql"
	"maranatha_web/models"
	"maranatha_web/utils/errors"
)

var ctx = context.Background()

func GetUserByEmail(email string) (*errors.RestErr, *models.User) {
	var user models.User
	//user := new(models.User)
	if err := postgresql_db.Client.NewSelect().
		Model(user).
		Where("email = ?", email).
		Scan(ctx); err != nil {
		return errors.NewBadRequestError("User does not exit"), nil
	}
	//m := new(models.User)
	////err := postgresql_db.Client.Debug().Where("email = ?", param).First(&user).Scan(&user).Error
	//err, _ := postgresql_db.Client.NewSelect().Model(()(nil)).Count(ctx)
	//if err != nil && err != gorm.ErrRecordNotFound {
	//	return errors.NewBadRequestError("User does not exit"), &user
	//}
	return nil, &user
}

func UpdateVerifiedUserStatus(param string) *errors.RestErr {
	//err := postgresql_db.Client.Table("users").Debug().Where("email = ?", param).Update("is_verified", true).Error
	//if err != nil {
	//	log.Panicln(err)
	//
	//}
	return nil
}

func Create(user *models.User) *errors.RestErr {
	//err := postgresql_db.Client.Debug().Create(&user).Error
	//
	if data, err := postgresql_db.Client.NewInsert().Model(user).Exec(ctx); err != nil {
		log.Println(err)
		log.Println(data)
		return errors.NewBadRequestError("Could not insert user")
	}
	//if err != nil {
	//	logger.Error("error when trying to save user", err)
	//
	//	return errors.NewInternalServerError("database error")
	//}
	return nil
}
