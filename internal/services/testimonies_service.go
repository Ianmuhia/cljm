package services

// import (
// 	"fmt"
// 	"log"
// 	"time"
//

// 	"maranatha_web/domain/testimonies"
// 	"maranatha_web/models"
// 	"maranatha_web/utils/errors"
// )

// var (
// 	TestimoniesService testimoniesServiceInterface = &testimoniesService{}
// )

// type testimoniesService struct{}

// type testimoniesServiceInterface interface {
// 	CreateTestimony(testimoniesModel models.Testimonies) (*models.Testimonies, *errors.RestErr)
// 	GetAllTestimonies() ([]*models.Testimonies, int64, *errors.RestErr)
// 	DeleteTestimony(id uint) *errors.RestErr
// 	GetSingleTestimony(id uint) (*models.Testimonies, *errors.RestErr)
// 	UpdateTestimony(id uint, testimoniesModel models.Testimonies) *errors.RestErr
// 	GetAllTestimoniesByAuthor(id uint) ([]*models.Testimonies, int64, *errors.RestErr)
// }

// func (s *testimoniesService) CreateTestimony(testimoniesModel models.Testimonies) (*models.Testimonies, *errors.RestErr) {
// 	if err := testimonies.CreateTestimony(&testimoniesModel); err != nil {
// 		return nil, err
// 	}
// 	return &testimoniesModel, nil
// }

// func (s *testimoniesService) GetAllTestimonies() ([]*models.Testimonies, int64, *errors.RestErr) {
// 	data, count, err := testimonies.GetAllTestimonies()
// 	for _, v := range data {
// 		d := v.CreatedAt.Format(time.RFC822)

// 		myDate, err := time.Parse(time.RFC822, d)
// 		if err != nil {
// 			panic(err)
// 		}

// 		v.CreatedAt = myDate
// 		fmt.Println(v.CreatedAt.Format(time.RFC1123))
// 	}
// 	if err != nil {
// 		return data, count, errors.NewBadRequestError("Could not get testimonies")

// 	}

// 	return data, count, nil
// }

// func (s *testimoniesService) DeleteTestimony(id uint) *errors.RestErr {
// 	err := testimonies.DeleteTestimony(id)
// 	if err != nil {
// 		log.Println(err)
// 		return errors.NewBadRequestError("Could not delete testimonies")
// 	}
// 	return nil
// }

// func (s *testimoniesService) GetSingleTestimony(id uint) (*models.Testimonies, *errors.RestErr) {
// 	testimonies, err := testimonies.GetSingleTestimony(id)
// 	if err != nil {
// 		log.Println(err)
// 		return testimonies, errors.NewBadRequestError("Could not get single testimonies")
// 	}
// 	return testimonies, nil
// }

// func (s *testimoniesService) UpdateTestimony(id uint, testimoniesModel models.Testimonies) *errors.RestErr {
// 	_, err := testimonies.UpdateTestimony(id, testimoniesModel)
// 	if err != nil {
// 		log.Println(err)
// 		return errors.NewBadRequestError("Could not get single testimonies")
// 	}
// 	return nil
// }

// func (s *testimoniesService) GetAllTestimoniesByAuthor(id uint) ([]*models.Testimonies, int64, *errors.RestErr) {
// 	testimoniesData, count, err := testimonies.GetAllTestimoniesByAuthor(id)
// 	if err != nil {
// 		log.Println(err)
// 		return testimoniesData, count, errors.NewBadRequestError("Could not get post by author.")
// 	}
// 	return testimoniesData, count, nil
// }
