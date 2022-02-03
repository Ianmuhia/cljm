package services

import (
	"fmt"
	"log"
	"maranatha_web/internal/models"
	"maranatha_web/internal/repository"
	"time"
)

type TestimoniesService interface {
	CreateTestimony(testimoniesModel models.Testimonies) (*models.Testimonies, error)
	GetAllTestimonies() ([]*models.Testimonies, int64, error)
	DeleteTestimony(id uint) error
	GetSingleTestimony(id uint) (*models.Testimonies, error)
	UpdateTestimony(id uint, testimoniesModel models.Testimonies) error
	GetAllTestimoniesByAuthor(id uint) ([]*models.Testimonies, int64, error)
}

type testimoniesService struct {
	dao repository.DAO
}

func NewTestimoniesService(dao repository.DAO) TestimoniesService {
	return &testimoniesService{dao: dao}
}

func (ts *testimoniesService) CreateTestimony(testimoniesModel models.Testimonies) (*models.Testimonies, error) {
	if err := ts.dao.NewTestimonyQuery().CreateTestimony(&testimoniesModel); err != nil {
		return nil, err
	}
	return &testimoniesModel, nil
}

func (ts *testimoniesService) GetAllTestimonies() ([]*models.Testimonies, int64, error) {
	data, count, err := ts.dao.NewTestimonyQuery().GetAllTestimonies()
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

func (ts *testimoniesService) DeleteTestimony(id uint) error {
	err := ts.dao.NewTestimonyQuery().DeleteTestimony(id)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (ts *testimoniesService) GetSingleTestimony(id uint) (*models.Testimonies, error) {
	testimonies, err := ts.dao.NewTestimonyQuery().GetSingleTestimony(id)
	if err != nil {
		log.Println(err)
		return testimonies, err
	}
	return testimonies, nil
}

func (ts *testimoniesService) UpdateTestimony(id uint, testimoniesModel models.Testimonies) error {
	_, err := ts.dao.NewTestimonyQuery().UpdateTestimony(id, testimoniesModel)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (ts *testimoniesService) GetAllTestimoniesByAuthor(id uint) ([]*models.Testimonies, int64, error) {
	testimoniesData, count, err := ts.dao.NewTestimonyQuery().GetAllTestimoniesByAuthor(id)
	if err != nil {
		log.Println(err)
		return testimoniesData, count, err
	}
	return testimoniesData, count, nil
}
