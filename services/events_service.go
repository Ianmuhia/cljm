package services

import (
	"fmt"
	"log"
	"time"

	"maranatha_web/domain/events"
	"maranatha_web/models"
	"maranatha_web/utils/errors"
)

var (
	EventsService eventsServiceInterface = &eventsService{}
)

type eventsService struct{}

type eventsServiceInterface interface {
	CreateEventsPost(eventsModel models.ChurchEvent) (*models.ChurchEvent, *errors.RestErr)
	GetAllEventsPost() ([]*models.ChurchEvent, int64, *errors.RestErr)
	DeleteEventsPost(id uint) *errors.RestErr
	GetSingleEventsPost(id uint) (*models.ChurchEvent, *errors.RestErr)
	UpdateEventsPost(id uint, newModel models.ChurchEvent) *errors.RestErr
	GetAllEventsPostByAuthor(id uint) ([]*models.ChurchEvent, int64, *errors.RestErr)
}

func (s *eventsService) CreateEventsPost(eventsModel models.ChurchEvent) (*models.ChurchEvent, *errors.RestErr) {
	if err := events.CreateEventsPost(&eventsModel); err != nil {
		return nil, err
	}
	return &eventsModel, nil
}

func (s *eventsService) GetAllEventsPost() ([]*models.ChurchEvent, int64, *errors.RestErr) {
	data, count, err := events.GetAllEventsPost()
	for _, v := range data {
		v.CoverImage = fmt.Sprintf("http://localhost:9000/mono/%s", v.CoverImage)

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

	return data, count, nil
}

func (s *eventsService) DeleteEventsPost(id uint) *errors.RestErr {
	err := events.DeleteEventsPost(id)
	if err != nil {
		log.Println(err)
		return errors.NewBadRequestError("Could not delete post")
	}
	return nil
}

func (s *eventsService) GetSingleEventsPost(id uint) (*models.ChurchEvent, *errors.RestErr) {
	events, err := events.GetSingleEventsPost(id)
	url := fmt.Sprintf("http://localhost:9000/mono/%s", events.CoverImage)
	events.CoverImage = url
	if err != nil {
		log.Println(err)
		return events, errors.NewBadRequestError("Could not get single post")
	}
	return events, nil
}

func (s *eventsService) UpdateEventsPost(id uint, newModel models.ChurchEvent) *errors.RestErr {
	events, err := events.UpdateEventsPost(id, newModel)
	url := fmt.Sprintf("http://localhost:9000/mono/%s", events.CoverImage)
	events.CoverImage = url
	if err != nil {
		log.Println(err)
		return errors.NewBadRequestError("Could not get single post")
	}
	return nil
}

func (s *eventsService) GetAllEventsPostByAuthor(id uint) ([]*models.ChurchEvent, int64, *errors.RestErr) {
	eventsData, count, err := events.GetAllEventsPostByAuthor(id)
	for _, v := range eventsData {
		v.CoverImage = fmt.Sprintf("http://localhost:9000/mono/%s", v.CoverImage)
	}
	if err != nil {
		log.Println(err)
		return eventsData, count, errors.NewBadRequestError("Could not get post by author.")
	}
	return eventsData, count, nil
}
