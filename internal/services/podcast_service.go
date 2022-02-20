package services

import (
	"fmt"
	"log"
	"maranatha_web/internal/config"
	"maranatha_web/internal/models"
	"maranatha_web/internal/repository"
	"time"
)

type PodcastService interface {
	CreatePodcast(podcastModel models.Podcast) (*models.Podcast, error)
	GetAllPodcast() ([]*models.Podcast, int, error)
	DeletePodcast(id uint) error
	GetSinglePodcast(id uint) (*models.Podcast, error)
	UpdatePodcast(id uint, newModel models.Podcast) error
	GetAllPodcastByAuthor(id uint) ([]*models.Podcast, int, error)
}

type podcastService struct {
	dao repository.DAO
	cfg *config.AppConfig
}

func NewPodcastService(dao repository.DAO, cfg *config.AppConfig) PodcastService {
	return &podcastService{dao: dao, cfg: cfg}
}
func (ps *podcastService) CreatePodcast(podcastModel models.Podcast) (*models.Podcast, error) {
	if err := ps.dao.NewPodcastQuery().CreatePodcast(&podcastModel); err != nil {
		return nil, err
	}
	return &podcastModel, nil
}

func (ps *podcastService) GetAllPodcast() ([]*models.Podcast, int, error) {
	data, count, err := ps.dao.NewPodcastQuery().GetAllPodcast()
	for _, v := range data {

		v.CoverImage = fmt.Sprintf("%s/%s/%s", ps.cfg.StorageURL.String(), ps.cfg.StorageBucket, v.CoverImage)
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

func (ps *podcastService) GetSinglePodcast(id uint) (*models.Podcast, error) {
	podcast, err := ps.dao.NewPodcastQuery().GetSinglePodcast(id)
	url := fmt.Sprintf("%s/%s/%s", ps.cfg.StorageURL.String(), ps.cfg.StorageBucket, podcast.CoverImage)
	url2 := fmt.Sprintf("%s/%s/%s", ps.cfg.StorageURL.String(), ps.cfg.StorageBucket, podcast.Author.ProfileImage)
	podcast.CoverImage = url
	podcast.Author.ProfileImage = url2
	if err != nil {
		log.Println(err)
		return podcast, err
	}
	return podcast, nil
}

func (ps *podcastService) UpdatePodcast(id uint, newModel models.Podcast) error {
	podcast, err := ps.dao.NewPodcastQuery().UpdatePodcast(id, newModel)
	url := fmt.Sprintf("%s/%s/%s", ps.cfg.StorageURL.String(), ps.cfg.StorageBucket, podcast.CoverImage)
	podcast.CoverImage = url
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (ps *podcastService) GetAllPodcastByAuthor(id uint) ([]*models.Podcast, int, error) {
	podcastData, count, err := ps.dao.NewPodcastQuery().GetAllPodcastByAuthor(id)
	for _, v := range podcastData {
		v.CoverImage = fmt.Sprintf("%s/%s/%s", ps.cfg.StorageURL.String(), ps.cfg.StorageBucket, v.CoverImage)
	}
	if err != nil {
		log.Println(err)
		return podcastData, count, err
	}
	return podcastData, count, nil
}

func (ps *podcastService) DeletePodcast(id uint) error {
	err := ps.dao.NewPodcastQuery().DeletePodcast(id)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
