package services

import (
	"fmt"
	"log"
	"time"

	"maranatha_web/domain/prayer_request"
	"maranatha_web/models"
	"maranatha_web/utils/errors"
)

var (
	PrayerService prayerServiceInterface = &prayerService{}
)

type prayerService struct{}

type prayerServiceInterface interface {
	CreatePrayerRequest(prayerModel models.Prayer) (*models.Prayer, *errors.RestErr)
	GetAllPrayerRequests() ([]*models.Prayer, int64, *errors.RestErr)
	DeletePrayerRequest(id uint) *errors.RestErr
	GetSinglePrayerRequest(id uint) (*models.Prayer, *errors.RestErr)
	UpdatePrayerRequest(id uint, prayerModel models.Prayer) *errors.RestErr
	GetAllPrayerRequestsByAuthor(id uint) ([]*models.Prayer, int64, *errors.RestErr)
}

func (s *prayerService) CreatePrayerRequest(prayerModel models.Prayer) (*models.Prayer, *errors.RestErr) {
	if err := prayer_request.CreatePrayerRequest(&prayerModel); err != nil {
		return nil, err
	}
	return &prayerModel, nil
}

func (s *prayerService) GetAllPrayerRequests() ([]*models.Prayer, int64, *errors.RestErr) {
	data, count, err := prayer_request.GetAllPrayerRequests()
	for _, v := range data {
		d := v.CreatedAt.Format(time.RFC822)

		myDate, err := time.Parse(time.RFC822, d)
		if err != nil {
			panic(err)
		}

		v.CreatedAt = myDate
		fmt.Println(v.CreatedAt.Format(time.RFC1123))
	}
	if err != nil {
		return data, count, errors.NewBadRequestError("Could not get prayers")

	}

	return data, count, nil
}

func (s *prayerService) DeletePrayerRequest(id uint) *errors.RestErr {
	err := prayer_request.DeletePrayerRequest(id)
	if err != nil {
		log.Println(err)
		return errors.NewBadRequestError("Could not delete prayer")
	}
	return nil
}

func (s *prayerService) GetSinglePrayerRequest(id uint) (*models.Prayer, *errors.RestErr) {
	prayer, err := prayer_request.GetSinglePrayerRequest(id)
	if err != nil {
		log.Println(err)
		return prayer, errors.NewBadRequestError("Could not get single prayer")
	}
	return prayer, nil
}

func (s *prayerService) UpdatePrayerRequest(id uint, prayerModel models.Prayer) *errors.RestErr {
	_, err := prayer_request.UpdatePrayerRequest(id, prayerModel)
	if err != nil {
		log.Println(err)
		return errors.NewBadRequestError("Could not get single prayer")
	}
	return nil
}

func (s *prayerService) GetAllPrayerRequestsByAuthor(id uint) ([]*models.Prayer, int64, *errors.RestErr) {
	prayerData, count, err := prayer_request.GetAllPrayerRequestsByAuthor(id)
	if err != nil {
		log.Println(err)
		return prayerData, count, errors.NewBadRequestError("Could not get prayer requests by author.")
	}
	return prayerData, count, nil
}
