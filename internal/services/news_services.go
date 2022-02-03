package services

import (
	"fmt"
	"log"
	"maranatha_web/internal/models"
	"maranatha_web/internal/repository"
	"time"
)

type NewsService interface {
	CreateNewsPost(newsModel models.News) (*models.News, error)
	GetAllNewsPost() ([]*models.News, int64, error)
	DeleteNewsPost(id uint) error
	GetSingleNewsPost(id uint) (*models.News, error)
	UpdateNewsPost(id uint, newModel models.News) error
	GetAllNewsPostByAuthor(id uint) ([]*models.News, int64, error)
}

type newsService struct {
	dao repository.DAO
}

func NewNewsService(dao repository.DAO) NewsService {
	return &newsService{dao: dao}
}

func (ns *newsService) CreateNewsPost(newsModel models.News) (*models.News, error) {
	if err := ns.dao.NewNewsQuery().CreateNewsPost(&newsModel); err != nil {
		return nil, err
	}
	return &newsModel, nil
}

func (ns *newsService) GetAllNewsPost() ([]*models.News, int64, error) {
	data, count, err := ns.dao.NewNewsQuery().GetAllNewsPost()
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
		return data, count, err

	}

	return data, count, nil
}

func (ns *newsService) DeleteNewsPost(id uint) error {
	err := ns.dao.NewNewsQuery().DeleteNewsPost(id)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (ns *newsService) GetSingleNewsPost(id uint) (*models.News, error) {
	news, err := ns.dao.NewNewsQuery().GetSingleNewsPost(id)
	url := fmt.Sprintf("http://localhost:9000/mono/%s", news.CoverImage)
	news.CoverImage = url
	if err != nil {
		log.Println(err)
		return news, err
	}
	return news, nil
}

func (ns *newsService) UpdateNewsPost(id uint, newModel models.News) error {
	news, err := ns.dao.NewNewsQuery().UpdateNewsPost(id, newModel)
	url := fmt.Sprintf("http://localhost:9000/mono/%s", news.CoverImage)
	news.CoverImage = url
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (ns *newsService) GetAllNewsPostByAuthor(id uint) ([]*models.News, int64, error) {
	newsData, count, err := ns.dao.NewNewsQuery().GetAllNewsPostByAuthor(id)
	for _, v := range newsData {
		v.CoverImage = fmt.Sprintf("http://localhost:9000/mono/%s", v.CoverImage)
	}
	if err != nil {
		log.Println(err)
		return newsData, count, err
	}
	return newsData, count, nil
}
