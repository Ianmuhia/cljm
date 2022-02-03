package repository

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"log"
	"maranatha_web/internal/models"
)

type EventsQuery interface {
	CreateEvent(events *models.ChurchEvent) error
	DeleteEvent(id uint) error
	GetSingleEvent(id uint) (*models.ChurchEvent, error)
	UpdateEventsPost(id uint, eventsModel models.ChurchEvent) (*models.ChurchEvent, error)
	GetAllEvents() ([]*models.ChurchEvent, int64, error)
	GetAllEventsByAuthor(id uint) ([]*models.ChurchEvent, int64, error)
}

type eventsQuery struct {
	dbRepo postgresDBRepo
}

func (eq *eventsQuery) CreateEvent(events *models.ChurchEvent) error {
	err := eq.dbRepo.DB.Debug().Create(&events).Error
	if err != nil {
		eq.dbRepo.App.ErrorLog.Error("error when trying to create church event.", zap.Any("error", err))
		return err
	}
	return nil
}

func (eq *eventsQuery) DeleteEvent(id uint) error {
	var events models.ChurchEvent
	err := eq.dbRepo.DB.Debug().Where("id = ?", id).Delete(&events).Error
	if err != nil {
		eq.dbRepo.App.ErrorLog.Error("error when trying to delete events post", zap.Any("error", err))
		return err
	}
	return nil
}

func (eq *eventsQuery) GetSingleEvent(id uint) (*models.ChurchEvent, error) {
	var events models.ChurchEvent
	err := eq.dbRepo.DB.Debug().Preload(clause.Associations).Where("id = ?", id).First(&events).Error
	if err != nil || err == gorm.ErrRecordNotFound {
		eq.dbRepo.App.ErrorLog.Error("error when trying to get  events post", zap.Any("error", err))
		return &events, err
	}
	return &events, nil
}

func (eq *eventsQuery) UpdateEventsPost(id uint, eventsModel models.ChurchEvent) (*models.ChurchEvent, error) {
	err := eq.dbRepo.DB.Debug().Where("id = ?", id).Updates(&eventsModel).Error
	if err != nil || err == gorm.ErrRecordNotFound {
		eq.dbRepo.App.ErrorLog.Error("error when trying to update  events post", zap.Any("error", err))
		return &eventsModel, err
	}
	return &eventsModel, nil
}

func (eq *eventsQuery) GetAllEvents() ([]*models.ChurchEvent, int64, error) {
	var events []*models.ChurchEvent
	var count int64
	val := eq.dbRepo.DB.Table("church_events").Preload("ChurchJobs").Preload("Organizer").Find(&events).Error
	count = int64(len(events))
	if val != nil {

		return events, 0, val
	}
	return events, count, nil
}

func (eq *eventsQuery) GetAllEventsByAuthor(id uint) ([]*models.ChurchEvent, int64, error) {
	var events []*models.ChurchEvent
	var count int64
	val := eq.dbRepo.DB.Debug().Where("organizer_id = ?", id).Table("church_events").Order("created_at desc").Find(&events).Count(&count).Error
	if val != nil {
		log.Println(val)
		return nil, 0, val
	}
	return events, count, nil
}
