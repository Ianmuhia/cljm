package users

import (
	"context"
	"log"

	postgresql_db "maranatha_web/datasources/postgresql"
	"maranatha_web/models"
	"maranatha_web/utils/errors"
)

var ctx = context.Background()

func GetUserByEmail(email string) (*models.User, error) {
	//var user models.User
	user := new(models.User)
	if err := postgresql_db.Client.NewSelect().
		Model(user).
		Where("email = ?", email).
		Scan(ctx); err != nil {
		log.Println(err)
		return user, nil
	}
	log.Println(&user)
	return user, nil
}

func UpdateVerifiedUserStatus(param string) *errors.RestErr {

	var model *models.User

	if data, err := postgresql_db.Client.NewUpdate().Model(model).Where("email = ?", param).Set("is_verified = ?", true).Exec(ctx); err != nil {
		log.Println(err)
		log.Println(data)
		return errors.NewBadRequestError("Could not update user status")
	}
	return nil

}

func Create(user *models.User) *errors.RestErr {
	if data, err := postgresql_db.Client.NewInsert().Model(user).Exec(ctx); err != nil {
		log.Println(err)
		log.Println(data)
		return errors.NewBadRequestError("Could not insert user")
	}
	return nil
}
