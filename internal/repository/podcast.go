package repository

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"log"
	"maranatha_web/internal/models"
)

type PodcastQuery interface {
	CreatePodcast(podcast *models.Podcast) error
	DeletePodcast(id uint) error
	GetSinglePodcast(id uint) (*models.Podcast, error)
	UpdatePodcast(id uint, podcastModel models.Podcast) (*models.Podcast, error)
	GetAllPodcast() ([]*models.Podcast, int, error)
	GetAllPodcastByAuthor(id uint) ([]*models.Podcast, int, error)
}

type podcastQuery struct {
	dbRepo postgresDBRepo
}

func (poq *podcastQuery) CreatePodcast(podcast *models.Podcast) error {
	err := poq.dbRepo.DB.Debug().Create(&podcast).Error
	if err != nil {
		poq.dbRepo.App.ErrorLog.Error("error when trying to save user", zap.Any("error", err))
		return err
	}
	return nil
}

func (poq *podcastQuery) DeletePodcast(id uint) error {
	var podcast models.Podcast
	err := poq.dbRepo.DB.Debug().Where("id = ?", id).Delete(&podcast).Error
	if err != nil {
		poq.dbRepo.App.ErrorLog.Error("error when trying to delete podcast post", zap.Any("error", err))
		return err
	}
	return nil
}
func (poq *podcastQuery) GetSinglePodcast(id uint) (*models.Podcast, error) {
	var podcast models.Podcast
	err := poq.dbRepo.DB.Debug().Preload(clause.Associations).Where("id = ?", id).First(&podcast).Error
	if err != nil || err == gorm.ErrRecordNotFound {
		poq.dbRepo.App.ErrorLog.Error("error when trying to get  podcast post", zap.Any("error", err))
		return &podcast, err
	}
	return &podcast, nil
}
func (poq *podcastQuery) UpdatePodcast(id uint, podcastModel models.Podcast) (*models.Podcast, error) {
	err := poq.dbRepo.DB.Debug().Where("id = ?", id).Updates(&podcastModel).Error
	if err != nil || err == gorm.ErrRecordNotFound {
		poq.dbRepo.App.ErrorLog.Error("error when trying to update  podcast post", zap.Any("error", err))
		return &podcastModel, err
	}
	return &podcastModel, nil
}

func (poq *podcastQuery) GetAllPodcast() ([]*models.Podcast, int, error) {
	var podcast []*models.Podcast
	var count int
	err := poq.dbRepo.DB.Debug().Preload(clause.Associations).Order("created_at desc").Find(&podcast).Error
	if err != nil {
		log.Println(err)
		return nil, 0, err
	}
	count = len(podcast)
	return podcast, count, nil
}

func (poq *podcastQuery) GetAllPodcastByAuthor(id uint) ([]*models.Podcast, int, error) {
	var podcast []*models.Podcast
	var count int
	val := poq.dbRepo.DB.Debug().Where("author_id = ?", id).Preload("Author").Order("created_at desc").Find(&podcast).Error
	if val != nil {
		return nil, 0, val
	}
	count = len(podcast)
	return podcast, count, nil
}
