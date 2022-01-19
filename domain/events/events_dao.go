package events

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"log"
	postgresql_db "maranatha_web/datasources/postgresql"
	"maranatha_web/logger"
	"maranatha_web/models"
	"maranatha_web/utils/errors"
)

func CreateEventsPost(events *models.Event) *errors.RestErr {
	err := postgresql_db.Client.Debug().Create(&events).Error
	if err != nil {
		logger.Error("error when trying to save user", err)
		return errors.NewInternalServerError("database error")
	}
	return nil
}

func DeleteEventsPost(id uint) *errors.RestErr {
	var events models.Event
	err := postgresql_db.Client.Debug().Where("id = ?", id).Delete(&events).Error
	if err != nil {
		logger.Error("error when trying to delete events post", err)
		return errors.NewInternalServerError("database error")
	}
	return nil
}
func GetSingleEventsPost(id uint) (*models.Event, *errors.RestErr) {
	var events models.Event
	err := postgresql_db.Client.Debug().Preload(clause.Associations).Where("id = ?", id).First(&events).Error
	if err != nil || err == gorm.ErrRecordNotFound {
		logger.Error("error when trying to get  events post", err)
		return &events, errors.NewInternalServerError("database error")
	}
	return &events, nil
}
func UpdateEventsPost(id uint, eventsModel models.Event) (*models.Event, *errors.RestErr) {
	err := postgresql_db.Client.Debug().Where("id = ?", id).Updates(&eventsModel).Error
	if err != nil || err == gorm.ErrRecordNotFound {
		logger.Error("error when trying to update  events post", err)
		return &eventsModel, errors.NewInternalServerError("database error")
	}
	return &eventsModel, nil
}

func GetAllEventsPost() ([]*models.Event, int64, error) {
	var events []*models.Event
	var count int64
	val := postgresql_db.Client.Debug().Preload(clause.Associations).Order("created_at desc").Find(&events).Error
	if val != nil {
		log.Println(val)
		return nil, 0, val
	}
	return events, count, nil
}

func GetAllEventsPostByAuthor(id uint) ([]*models.Event, int64, error) {
	var events []*models.Event
	var count int64
	val := postgresql_db.Client.Debug().Where("author_id = ?", id).Preload("Author").Order("created_at desc").Find(&events).Count(&count).Error
	if val != nil {
		log.Println(val)
		return nil, 0, val
	}
	return events, count, nil
}

//Get all users
