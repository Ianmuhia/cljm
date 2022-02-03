package services

import (
	"fmt"
	"log"
	"maranatha_web/internal/logger"
	"maranatha_web/internal/models"
	"maranatha_web/internal/repository"
	"maranatha_web/internal/utils/errors"
	"time"
)

type EventsServiceInterface interface {
	CreateEvent(eventsModel models.ChurchEvent) (*models.ChurchEvent, error)
	GetAllEvents() ([]*models.ChurchEvent, int64, *errors.RestErr)
	DeleteEvent(id uint) *errors.RestErr
	GetSingleEvent(id uint) (*models.ChurchEvent, *errors.RestErr)
	UpdateEventsPost(id uint, newModel models.ChurchEvent) *errors.RestErr
	GetAllEventsByAuthor(id uint) ([]*models.ChurchEvent, int64, *errors.RestErr)
}

type eventsService struct {
	dao repository.DAO
}

func NewEventsService(dao repository.DAO) EventsServiceInterface {
	return &eventsService{dao: dao}
}

func (es *eventsService) CreateEvent(eventsModel models.ChurchEvent) (*models.ChurchEvent, error) {
	if err := es.dao.NewEventsQuery().CreateEvent(&eventsModel); err != nil {
		return nil, err
	}
	return &eventsModel, nil
}

func (es *eventsService) GetAllEvents() ([]*models.ChurchEvent, int64, *errors.RestErr) {
	data, count, err := es.dao.NewEventsQuery().GetAllEvents()
	for _, v := range data {
		v.CoverImage = fmt.Sprintf("http://192.168.0.101:9000/mono/%s", v.CoverImage)

		d := v.CreatedAt.Format(time.RFC822)

		myDate, err := time.Parse(time.RFC822, d)
		if err != nil {
			panic(err)
		}

		v.CreatedAt = myDate
		fmt.Println(v.CreatedAt.Format(time.RFC1123))

	}
	if err != nil {
		return data, count, errors.NewBadRequestError("Could not get post")

	}
	logger.GetLogger().Info("Events endpoint hit.")
	logger.GetLogger().Error("Events endpoint hit.")
	return data, count, nil
}

func (es *eventsService) DeleteEvent(id uint) *errors.RestErr {
	err := es.dao.NewEventsQuery().DeleteEvent(id)
	if err != nil {
		log.Println(err)
		return errors.NewBadRequestError("Could not delete post")
	}
	return nil
}

func (es *eventsService) GetSingleEvent(id uint) (*models.ChurchEvent, *errors.RestErr) {
	events, err := es.dao.NewEventsQuery().GetSingleEvent(id)
	url := fmt.Sprintf("http://192.168.0.101:9000/mono/%s", events.CoverImage)
	events.CoverImage = url
	if err != nil {
		log.Println(err)
		return events, errors.NewBadRequestError("Could not get single post")
	}
	return events, nil
}

func (es *eventsService) UpdateEventsPost(id uint, newModel models.ChurchEvent) *errors.RestErr {
	events, err := es.dao.NewEventsQuery().UpdateEventsPost(id, newModel)
	url := fmt.Sprintf("http://192.168.0.101:9000/mono/%s", events.CoverImage)
	events.CoverImage = url
	if err != nil {
		log.Println(err)

		return errors.NewBadRequestError("Could not get single post")
	}
	return nil
}

func (es *eventsService) GetAllEventsByAuthor(id uint) ([]*models.ChurchEvent, int64, *errors.RestErr) {
	eventsData, count, err := es.dao.NewEventsQuery().GetAllEventsByAuthor(id)
	for _, v := range eventsData {
		v.CoverImage = fmt.Sprintf("http://192.168.0.101:9000/mono/%s", v.CoverImage)
	}
	if err != nil {
		log.Println(err)
		return eventsData, count, errors.NewBadRequestError("Could not get post by author.")
	}
	return eventsData, count, nil
}
