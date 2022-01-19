package services

import (
	"fmt"
	"log"
	"maranatha_web/domain/prayer_request"
	"maranatha_web/models"
	"maranatha_web/utils/errors"
	"time"
)

var (
	PrayerService prayerServiceInterface = &prayerService{}
)

type prayerService struct{}

type prayerServiceInterface interface {
	CreatePrayerPost(prayerModel models.Prayer) (*models.Prayer, *errors.RestErr)
	GetAllPrayerPost() ([]*models.Prayer, int64, *errors.RestErr)
	DeletePrayerPost(id uint) *errors.RestErr
	GetSinglePrayerPost(id uint) (*models.Prayer, *errors.RestErr)
	UpdatePrayerPost(id uint, prayerModel models.Prayer) *errors.RestErr
	GetAllPrayerPostByAuthor(id uint) ([]*models.Prayer, int64, *errors.RestErr)
}

func (s *prayerService) CreatePrayerPost(prayerModel models.Prayer) (*models.Prayer, *errors.RestErr) {
	if err := prayer_request.CreatePrayerPost(&prayerModel); err != nil {
		return nil, err
	}
	return &prayerModel, nil
}

func (s *prayerService) GetAllPrayerPost() ([]*models.Prayer, int64, *errors.RestErr) {
	data, count, err := prayer_request.GetAllPrayerPost()
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

func (s *prayerService) DeletePrayerPost(id uint) *errors.RestErr {
	err := prayer_request.DeletePrayerPost(id)
	if err != nil {
		log.Println(err)
		return errors.NewBadRequestError("Could not delete prayer")
	}
	return nil
}

func (s *prayerService) GetSinglePrayerPost(id uint) (*models.Prayer, *errors.RestErr) {
	prayer, err := prayer_request.GetSinglePrayerPost(id)
	if err != nil {
		log.Println(err)
		return prayer, errors.NewBadRequestError("Could not get single prayer")
	}
	return prayer, nil
}

func (s *prayerService) UpdatePrayerPost(id uint, prayerModel models.Prayer) *errors.RestErr {
	_, err := prayer_request.UpdatePrayerPost(id, prayerModel)
	if err != nil {
		log.Println(err)
		return errors.NewBadRequestError("Could not get single prayer")
	}
	return nil
}

func (s *prayerService) GetAllPrayerPostByAuthor(id uint) ([]*models.Prayer, int64, *errors.RestErr) {
	prayerData, count, err := prayer_request.GetAllPrayerPostByAuthor(id)
	if err != nil {
		log.Println(err)
		return prayerData, count, errors.NewBadRequestError("Could not get post by author.")
	}
	return prayerData, count, nil
}
