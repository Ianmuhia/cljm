package services

import (
	"fmt"
	"log"
	"time"

	"maranatha_web/internal/config"
	"maranatha_web/internal/models"
	"maranatha_web/internal/repository"
)

type NewsService interface {
	CreateNewsPost(newsModel models.News) (*models.News, error)
	GetAllNewsPost() ([]*models.News, int, error)
	DeleteNewsPost(id uint) error
	GetSingleNewsPost(id uint) (*models.News, error)
	UpdateNewsPost(id uint, newModel models.News) error
	GetAllNewsPostByAuthor(id uint) ([]*models.News, int, error)
}

type newsService struct {
	dao repository.DAO
	cfg *config.AppConfig
}

func NewNewsService(dao repository.DAO, cfg *config.AppConfig) NewsService {
	return &newsService{dao: dao, cfg: cfg}
}

func (ns *newsService) CreateNewsPost(newsModel models.News) (*models.News, error) {
	if err := ns.dao.NewNewsQuery().CreateNewsPost(&newsModel); err != nil {
		return nil, err
	}
	return &newsModel, nil
}

func (ns *newsService) GetAllNewsPost() ([]*models.News, int, error) {
	data, count, err := ns.dao.NewNewsQuery().GetAllNewsPost()
	for _, v := range data {
		//TODO:get this from app config

		v.CoverImage = fmt.Sprintf("%s/%s/%s", ns.cfg.StorageURL.String(), ns.cfg.StorageBucket, v.CoverImage)
		j, e := time.Parse(time.RFC3339, v.CreatedAt.Format(time.RFC3339))
		log.Println(j)
		if e != nil {
			log.Println(e)
		}
		v.CreatedAt = j
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
	url := fmt.Sprintf("%s/%s/%s", ns.cfg.StorageURL.String(), ns.cfg.StorageBucket, news.CoverImage)
	url2 := fmt.Sprintf("%s/%s/%s", ns.cfg.StorageURL.String(), ns.cfg.StorageBucket, news.Author.ProfileImage)
	news.CoverImage = url
	news.Author.ProfileImage = url2
	if err != nil {
		log.Println(err)
		return news, err
	}
	return news, nil
}

func (ns *newsService) UpdateNewsPost(id uint, newModel models.News) error {
	news, err := ns.dao.NewNewsQuery().UpdateNewsPost(id, newModel)
	url := fmt.Sprintf("%s/%s/%s", ns.cfg.StorageURL.String(), ns.cfg.StorageBucket, news.CoverImage)
	news.CoverImage = url
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (ns *newsService) GetAllNewsPostByAuthor(id uint) ([]*models.News, int, error) {
	newsData, count, err := ns.dao.NewNewsQuery().GetAllNewsPostByAuthor(id)
	for _, v := range newsData {
		v.CoverImage = fmt.Sprintf("%s/%s/%s", ns.cfg.StorageURL.String(), ns.cfg.StorageBucket, v.CoverImage)
	}
	if err != nil {
		log.Println(err)
		return newsData, count, err
	}
	return newsData, count, nil
}
