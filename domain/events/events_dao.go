package events

import (
	"log"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	postgresql_db "maranatha_web/datasources/postgresql"
	"maranatha_web/logger"
	"maranatha_web/models"
	"maranatha_web/utils/errors"
)

func CreateEvent(events *models.ChurchEvent) *errors.RestErr {
	err := postgresql_db.Client.Debug().Create(&events).Error
	if err != nil {
		logger.Error("error when trying to save user", err)
		return errors.NewInternalServerError("database error")
	}
	return nil
}

func DeleteEvent(id uint) *errors.RestErr {
	var events models.ChurchEvent
	err := postgresql_db.Client.Debug().Where("id = ?", id).Delete(&events).Error
	if err != nil {
		logger.Error("error when trying to delete events post", err)
		return errors.NewInternalServerError("database error")
	}
	return nil
}
func GetSingleEvent(id uint) (*models.ChurchEvent, *errors.RestErr) {
	var events models.ChurchEvent
	err := postgresql_db.Client.Debug().Preload(clause.Associations).Where("id = ?", id).First(&events).Error
	if err != nil || err == gorm.ErrRecordNotFound {
		logger.Error("error when trying to get  events post", err)
		return &events, errors.NewInternalServerError("database error")
	}
	return &events, nil
}
func UpdateEventsPost(id uint, eventsModel models.ChurchEvent) (*models.ChurchEvent, *errors.RestErr) {
	err := postgresql_db.Client.Debug().Where("id = ?", id).Updates(&eventsModel).Error
	if err != nil || err == gorm.ErrRecordNotFound {
		logger.Error("error when trying to update  events post", err)
		return &eventsModel, errors.NewInternalServerError("database error")
	}
	return &eventsModel, nil
}

func GetAllEvents() ([]*models.ChurchEvent, int64, error) {
	var events []*models.ChurchEvent
	var count int64
	val := postgresql_db.Client.Debug().Preload(clause.Associations).Order("created_at desc").Find(&events).Error
	if val != nil {
		log.Println(val)
		return nil, 0, val
	}
	return events, count, nil
}

func GetAllEventsByAuthor(id uint) ([]*models.ChurchEvent, int64, error) {
	var events []*models.ChurchEvent
	var count int64
	val := postgresql_db.Client.Debug().Where("author_id = ?", id).Preload("Author").Order("created_at desc").Find(&events).Count(&count).Error
	if val != nil {
		log.Println(val)
		return nil, 0, val
	}
	return events, count, nil
}

//Get all users
