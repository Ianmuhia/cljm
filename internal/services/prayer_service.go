package services

import (
	"fmt"
	"log"
	"maranatha_web/internal/models"
	"maranatha_web/internal/repository"
	"time"
)

type PrayerRequestServiceInterface interface {
	CreatePrayerRequest(prayerModel models.Prayer) (*models.Prayer, error)
	GetAllPrayerRequests() ([]*models.Prayer, int64, error)
	DeletePrayerRequest(id uint) error
	GetSinglePrayerRequest(id uint) (*models.Prayer, error)
	UpdatePrayerRequest(id uint, prayerModel models.Prayer) error
	GetAllPrayerRequestsByAuthor(id uint) ([]*models.Prayer, int64, error)
}

type prayerRequestService struct {
	dao repository.DAO
}

func NewPrayerRequestService(dao repository.DAO) PrayerRequestServiceInterface {
	return &prayerRequestService{dao: dao}
}

func (prs *prayerRequestService) CreatePrayerRequest(prayerModel models.Prayer) (*models.Prayer, error) {
	if err := prs.dao.NewPrayerRequestQuery().CreatePrayerRequest(&prayerModel); err != nil {
		return nil, err
	}
	return &prayerModel, nil
}

func (prs *prayerRequestService) GetAllPrayerRequests() ([]*models.Prayer, int64, error) {
	data, count, err := prs.dao.NewPrayerRequestQuery().GetAllPrayerRequests()
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
		return data, count, err

	}

	return data, count, nil
}

func (prs *prayerRequestService) DeletePrayerRequest(id uint) error {
	err := prs.dao.NewPrayerRequestQuery().DeletePrayerRequest(id)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (prs *prayerRequestService) GetSinglePrayerRequest(id uint) (*models.Prayer, error) {
	prayer, err := prs.dao.NewPrayerRequestQuery().GetSinglePrayerRequest(id)
	if err != nil {
		log.Println(err)
		return prayer, err
	}
	return prayer, nil
}

func (prs *prayerRequestService) UpdatePrayerRequest(id uint, prayerModel models.Prayer) error {
	_, err := prs.dao.NewPrayerRequestQuery().UpdatePrayerRequest(id, prayerModel)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (prs *prayerRequestService) GetAllPrayerRequestsByAuthor(id uint) ([]*models.Prayer, int64, error) {
	prayerData, count, err := prs.dao.NewPrayerRequestQuery().GetAllPrayerRequestsByAuthor(id)
	if err != nil {
		log.Println(err)
		return prayerData, count, err
	}
	return prayerData, count, nil
}
