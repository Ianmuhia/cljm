package news

import (
	"log"

	"gorm.io/gorm/clause" //nolint:goimports
	postgresql_db "maranatha_web/datasources/postgresql"
	"maranatha_web/logger"
	"maranatha_web/models"
	"maranatha_web/utils/errors"
)

func CreateNewsPost(news *models.News) *errors.RestErr {
	err := postgresql_db.Client.Debug().Create(&news).Error
	if err != nil {
		logger.Error("error when trying to save user", err)
		return errors.NewInternalServerError("database error")
	}
	return nil
}

func DeleteNewsPost(id uint) *errors.RestErr {
	var news models.News
	err := postgresql_db.Client.Debug().Where("id = ?", id).Delete(&news).Error
	if err != nil {
		logger.Error("error when trying to delete news post", err)
		return errors.NewInternalServerError("database error")
	}
	return nil
}

func GetAllNewsPost() ([]models.News, int64, error) {
	var news []models.News
	var count int64
	val := postgresql_db.Client.Debug().Preload(clause.Associations).Find(&news).Error
	if val != nil {
		log.Println(val)
		return nil, 0, val
	}

	//err := postgresql_db.Client.Debug().Find(&news).Count(&count).Error
	//if err != nil {
	//	log.Println(err)
	//	return news, count, err
	//}
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
