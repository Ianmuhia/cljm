package services

import (
	"fmt"
	"log"

	"maranatha_web/domain/news"
	"maranatha_web/models"
	"maranatha_web/utils/errors"
)

var (
	NewsService newsServiceInterface = &newsService{}
)

type newsService struct{}

type newsServiceInterface interface {
	CreateNewsPost(newsModel models.News) (*models.News, *errors.RestErr)
	GetAllNewsPost() ([]models.News, int64, *errors.RestErr)
	DeleteNewsPost(id uint) *errors.RestErr
	GetSingleNewsPost(id uint) (*models.News, *errors.RestErr)
}

func (s *newsService) CreateNewsPost(newsModel models.News) (*models.News, *errors.RestErr) {
	if err := news.CreateNewsPost(&newsModel); err != nil {
		return nil, err
	}
	return &newsModel, nil
}

func (s *newsService) GetAllNewsPost() ([]models.News, int64, *errors.RestErr) {
	data, count, err := news.GetAllNewsPost()
	if err != nil {
		return data, count, errors.NewBadRequestError("Could not get post")

	}

	return data, count, nil
}

func (s *newsService) DeleteNewsPost(id uint) *errors.RestErr {
	err := news.DeleteNewsPost(id)
	if err != nil {
		log.Println(err)
		return errors.NewBadRequestError("Could not delete post")
	}
	return nil
}

func (s *newsService) GetSingleNewsPost(id uint) (*models.News, *errors.RestErr) {
	news, err := news.GetSingleNewsPost(id)
	url := fmt.Sprintf("http://localhost:9000/mono/%s", news.CoverImage)
	news.CoverImage = url
	if err != nil {
		log.Println(err)
		return news, errors.NewBadRequestError("Could not get single post")
	}
	return news, nil
}

//func (s *newsService) GetUserByEmail(email string) (*models.User, error) {
//	user, err := users.GetUserByEmail(email)
//	if err != nil {
//		return user, err
//	}
//	return user, err
//}
//
//func (s *newsService) UpdateUserStatus(email string) *errors.RestErr {
//	err := users.UpdateVerifiedUserStatus(email)
//	if err != nil {
//		log.Println(err)
//		return err
//	}
//	return nil
//}
