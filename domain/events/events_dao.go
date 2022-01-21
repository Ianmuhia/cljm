package events

import "C"
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
		logger.Error("error when trying to create church event.", err)
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

	//val := postgresql_db.Client.Raw("SELECT ce.*,eo.*, cj.* FROM church_events AS ce, church_jobs AS cj, users AS eo\nWHERE cj.church_event_id = ce.id and eo.id = ce.organizer_id and ce.deleted_at IS NULL and cj.deleted_at IS NULL;").Preload(clause.Associations).Scan(&events).Error
	val := postgresql_db.Client.Table("church_events").Preload("ChurchJobs").Preload("Organizer").Find(&events).Error
	//val := postgresql_db.Client.Model(events).
	//	Scan(&events).Error
	count = int64(len(events))
	if val != nil {
		log.Println(val)
		return events, 0, val
	}

	log.Println(&events)
	return events, count, nil
}

func GetAllEventsByAuthor(id uint) ([]*models.ChurchEvent, int64, error) {
	var events []*models.ChurchEvent
	var count int64
	val := postgresql_db.Client.Debug().Where("organizer_id = ?", id).Table("church_events").Order("created_at desc").Find(&events).Count(&count).Error
	if val != nil {
		log.Println(val)
		return nil, 0, val
	}
	return events, count, nil
}

//Get all users
