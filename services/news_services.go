package services

import (
	"fmt"
	"log"
	"time"

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
	GetAllNewsPost() ([]*models.News, int64, *errors.RestErr)
	DeleteNewsPost(id uint) *errors.RestErr
	GetSingleNewsPost(id uint) (*models.News, *errors.RestErr)
	UpdateNewsPost(id uint, newModel models.News) *errors.RestErr
	GetAllNewsPostByAuthor(id uint) ([]*models.News, int64, *errors.RestErr)
}

func (s *newsService) CreateNewsPost(newsModel models.News) (*models.News, *errors.RestErr) {
	if err := news.CreateNewsPost(&newsModel); err != nil {
		return nil, err
	}
	return &newsModel, nil
}

func (s *newsService) GetAllNewsPost() ([]*models.News, int64, *errors.RestErr) {
	data, count, err := news.GetAllNewsPost()
	for _, v := range data {
		v.CoverImage = fmt.Sprintf("http://localhost:9000/mono/%s", v.CoverImage)

		d := v.CreatedAt.Format(time.RFC822)

		myDate, err := time.Parse(time.RFC822, d)
		if err != nil {
			panic(err)
		}

		v.CreatedAt = myDate
		fmt.Println(v.CreatedAt.Format(time.RFC1123))
		// d, e := time.Parse("January 02, 2006",string(v.CreatedAt))
		// if e!=nil {
		// 	log.Println(e)
		// }
		// v.CreatedAt = d
	}
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

func (s *newsService) UpdateNewsPost(id uint, newModel models.News) *errors.RestErr {
	news, err := news.UpdateNewsPost(id, newModel)
	url := fmt.Sprintf("http://localhost:9000/mono/%s", news.CoverImage)
	news.CoverImage = url
	if err != nil {
		log.Println(err)
		return errors.NewBadRequestError("Could not get single post")
	}
	return nil
}

func (s *newsService) GetAllNewsPostByAuthor(id uint) ([]*models.News, int64, *errors.RestErr) {
	newsData, count, err := news.GetAllNewsPostByAuthor(id)
	for _, v := range newsData {
		v.CoverImage = fmt.Sprintf("http://localhost:9000/mono/%s", v.CoverImage)
	}
	if err != nil {
		log.Println(err)
		return newsData, count, errors.NewBadRequestError("Could not get post by author.")
	}
	return newsData, count, nil
}
