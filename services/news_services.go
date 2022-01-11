package services

import (
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
	CreateNewsPost(news_model models.News) (*models.News, *errors.RestErr)
	GetAllNewsPost() ([]models.News, *errors.RestErr)
}

func (s *newsService) CreateNewsPost(news_model models.News) (*models.News, *errors.RestErr) {
	if err := news.CreateNewsPost(&news_model); err != nil {
		return nil, err
	}
	log.Println(news_model)
	return &news_model, nil
}

func (s *newsService) GetAllNewsPost() ([]models.News, *errors.RestErr) {
	data, err := news.GetAllNewsPost()
	if err != nil {
		return data, errors.NewBadRequestError("Could not get post")

	}
	return data, nil
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
