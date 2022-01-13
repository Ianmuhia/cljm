package news

import (
	"context"
	"log"

	postgresql_db "maranatha_web/datasources/postgresql"
	"maranatha_web/logger"
	"maranatha_web/models"
	"maranatha_web/utils/errors"
)

var ctx = context.Background()

func CreateNewsPost(news *models.News) *errors.RestErr {
	err := postgresql_db.Client.Debug().Create(&news).Error
	if err != nil {
		logger.Error("error when trying to save user", err)
		return errors.NewInternalServerError("database error")
	}
	return nil
}

func DeleteNewsPost(news *models.News, id int64) *errors.RestErr {
	err := postgresql_db.Client.Debug().Unscoped().Delete(&news).Where("id = ?", id).Error
	if err != nil {
		logger.Error("error when trying to delete news post", err)
		return errors.NewInternalServerError("database error")
	}
	return nil
}

func GetAllNewsPost() ([]models.News, int64, error) {
	var news []models.News
	var count int64
	err := postgresql_db.Client.Debug().Find(&news).Count(&count).Error
	if err != nil {
		log.Println(err)
		return news, count, err
	}
	//TODO:Add password conform field
	return news, count, nil
}

//TODO:chrch patna image, name, since_date
//TODO:Prayer request prayer date...GetAll Getbyid Update,DEl
//TODO:Testimonial testimony ,user , date, GET testmony Edit Testimony Del testimony
//TODO:Sermon url . title , date_pub, duration
//TODO:Events, title ,sub_title, tag(marriage, wedo) schedule date
//TODO:Events, title ,sub_title, tag)
//TODO:Events,, Voluntia for event
//TODO:Books, genre title, file
//TODO:Get user voluntiad jobs...
//Get all users
