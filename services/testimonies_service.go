package services

import (
	"fmt"
	"log"
	"maranatha_web/domain/testimonies"
	"maranatha_web/models"
	"maranatha_web/utils/errors"
	"time"
)

var (
	TestimoniesService testimoniesServiceInterface = &testimoniesService{}
)

type testimoniesService struct{}

type testimoniesServiceInterface interface {
	CreateTestimoniesPost(testimoniesModel models.Testimonies) (*models.Testimonies, *errors.RestErr)
	GetAllTestimoniesPost() ([]*models.Testimonies, int64, *errors.RestErr)
	DeleteTestimoniesPost(id uint) *errors.RestErr
	GetSingleTestimoniesPost(id uint) (*models.Testimonies, *errors.RestErr)
	UpdateTestimoniesPost(id uint, testimoniesModel models.Testimonies) *errors.RestErr
	GetAllTestimoniesPostByAuthor(id uint) ([]*models.Testimonies, int64, *errors.RestErr)
}

func (s *testimoniesService) CreateTestimoniesPost(testimoniesModel models.Testimonies) (*models.Testimonies, *errors.RestErr) {
	if err := testimonies.CreateTestimoniesPost(&testimoniesModel); err != nil {
		return nil, err
	}
	return &testimoniesModel, nil
}

func (s *testimoniesService) GetAllTestimoniesPost() ([]*models.Testimonies, int64, *errors.RestErr) {
	data, count, err := testimonies.GetAllTestimoniesPost()
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
		return data, count, errors.NewBadRequestError("Could not get testimonies")

	}

	return data, count, nil
}

func (s *testimoniesService) DeleteTestimoniesPost(id uint) *errors.RestErr {
	err := testimonies.DeleteTestimoniesPost(id)
	if err != nil {
		log.Println(err)
		return errors.NewBadRequestError("Could not delete testimonies")
	}
	return nil
}

func (s *testimoniesService) GetSingleTestimoniesPost(id uint) (*models.Testimonies, *errors.RestErr) {
	testimonies, err := testimonies.GetSingleTestimoniesPost(id)
	if err != nil {
		log.Println(err)
		return testimonies, errors.NewBadRequestError("Could not get single testimonies")
	}
	return testimonies, nil
}

func (s *testimoniesService) UpdateTestimoniesPost(id uint, testimoniesModel models.Testimonies) *errors.RestErr {
	_, err := testimonies.UpdateTestimoniesPost(id, testimoniesModel)
	if err != nil {
		log.Println(err)
		return errors.NewBadRequestError("Could not get single testimonies")
	}
	return nil
}

func (s *testimoniesService) GetAllTestimoniesPostByAuthor(id uint) ([]*models.Testimonies, int64, *errors.RestErr) {
	testimoniesData, count, err := testimonies.GetAllTestimoniesPostByAuthor(id)
	if err != nil {
		log.Println(err)
		return testimoniesData, count, errors.NewBadRequestError("Could not get post by author.")
	}
	return testimoniesData, count, nil
}
