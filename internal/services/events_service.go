package services

import (
	"fmt"
	"log"
	"time"

	"maranatha_web/internal/logger"
	"maranatha_web/internal/models"
	"maranatha_web/internal/repository"
)

type EventsService interface {
	CreateEvent(eventsModel models.ChurchEvent) (*models.ChurchEvent, error)
	GetAllEvents() ([]*models.ChurchEvent, int64, error)
	DeleteEvent(id uint) error
	GetSingleEvent(id uint) (*models.ChurchEvent, error)
	UpdateEventsPost(id uint, newModel models.ChurchEvent) error
	GetAllEventsByAuthor(id uint) ([]*models.ChurchEvent, int64, error)
}

type eventsService struct {
	dao repository.DAO
}

func NewEventsService(dao repository.DAO) EventsService {
	return &eventsService{dao: dao}
}

func (es *eventsService) CreateEvent(eventsModel models.ChurchEvent) (*models.ChurchEvent, error) {
	if err := es.dao.NewEventsQuery().CreateEvent(&eventsModel); err != nil {
		return nil, err
	}
	return &eventsModel, nil
}

func (es *eventsService) GetAllEvents() ([]*models.ChurchEvent, int64, error) {
	data, count, err := es.dao.NewEventsQuery().GetAllEvents()
	for _, v := range data {
		v.CoverImage = fmt.Sprintf("http://0.0.0.0:9000/clj/%s", v.CoverImage)

		d := v.CreatedAt.Format(time.RFC822)

		myDate, err := time.Parse(time.RFC822, d)
		if err != nil {
			panic(err)
		}

		v.CreatedAt = myDate
		fmt.Println(v.CreatedAt.Format(time.RFC1123))

	}
	if err != nil {
		return data, count, err

	}
	logger.GetLogger().Info("Events endpoint hit.")
	logger.GetLogger().Error("Events endpoint hit.")
	return data, count, nil
}

func (es *eventsService) DeleteEvent(id uint) error {
	err := es.dao.NewEventsQuery().DeleteEvent(id)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (es *eventsService) GetSingleEvent(id uint) (*models.ChurchEvent, error) {
	events, err := es.dao.NewEventsQuery().GetSingleEvent(id)
	url := fmt.Sprintf("http://0.0.0.0:9000/clj/%s", events.CoverImage)
	events.CoverImage = url
	if err != nil {
		log.Println(err)
		return events, err
	}
	return events, nil
}

func (es *eventsService) UpdateEventsPost(id uint, newModel models.ChurchEvent) error {
	events, err := es.dao.NewEventsQuery().UpdateEventsPost(id, newModel)
	url := fmt.Sprintf("http://0.0.0.0:9000/clj/%s", events.CoverImage)
	events.CoverImage = url
	if err != nil {
		log.Println(err)

		return err
	}
	return nil
}

func (es *eventsService) GetAllEventsByAuthor(id uint) ([]*models.ChurchEvent, int64, error) {
	eventsData, count, err := es.dao.NewEventsQuery().GetAllEventsByAuthor(id)
	for _, v := range eventsData {
		v.CoverImage = fmt.Sprintf("http://0.0.0.0:9000/clj/%s", v.CoverImage)
	}
	if err != nil {
		log.Println(err)
		return eventsData, count, err
	}
	return eventsData, count, nil
}
