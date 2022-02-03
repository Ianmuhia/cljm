package repository

import (
	"go.uber.org/zap"
	"gorm.io/gorm"        //nolint:goimports
	"gorm.io/gorm/clause" //nolint:goimports
	"log"
	"maranatha_web/internal/models"
)

type NewsQuery interface {
	CreateNewsPost(news *models.News) error
	DeleteNewsPost(id uint) error
	GetSingleNewsPost(id uint) (*models.News, error)
	UpdateNewsPost(id uint, newsModel models.News) (*models.News, error)
	GetAllNewsPost() ([]*models.News, int64, error)
	GetAllNewsPostByAuthor(id uint) ([]*models.News, int64, error)
}

type newsQuery struct {
	dbRepo postgresDBRepo
}

func (nq *newsQuery) CreateNewsPost(news *models.News) error {
	err := nq.dbRepo.DB.Debug().Create(&news).Error
	if err != nil {
		nq.dbRepo.App.ErrorLog.Error("error when trying to save user", zap.Any("error", err))
		return err
	}
	return nil
}

func (nq *newsQuery) DeleteNewsPost(id uint) error {
	var news models.News
	err := nq.dbRepo.DB.Debug().Where("id = ?", id).Delete(&news).Error
	if err != nil {
		nq.dbRepo.App.ErrorLog.Error("error when trying to delete news post", zap.Any("error", err))
		return err
	}
	return nil
}
func (nq *newsQuery) GetSingleNewsPost(id uint) (*models.News, error) {
	var news models.News
	err := nq.dbRepo.DB.Debug().Preload(clause.Associations).Where("id = ?", id).First(&news).Error
	if err != nil || err == gorm.ErrRecordNotFound {
		nq.dbRepo.App.ErrorLog.Error("error when trying to get  news post", zap.Any("error", err))
		return &news, err
	}
	return &news, nil
}
func (nq *newsQuery) UpdateNewsPost(id uint, newsModel models.News) (*models.News, error) {
	err := nq.dbRepo.DB.Debug().Where("id = ?", id).Updates(&newsModel).Error
	if err != nil || err == gorm.ErrRecordNotFound {
		nq.dbRepo.App.ErrorLog.Error("error when trying to update  news post", zap.Any("error", err))
		return &newsModel, err
	}
	return &newsModel, nil
}

func (nq *newsQuery) GetAllNewsPost() ([]*models.News, int64, error) {
	var news []*models.News
	var count int64
	val := nq.dbRepo.DB.Debug().Preload(clause.Associations).Order("created_at desc").Find(&news).Error
	if val != nil {
		log.Println(val)
		return nil, 0, val
	}
	return news, count, nil
}

func (nq *newsQuery) GetAllNewsPostByAuthor(id uint) ([]*models.News, int64, error) {
	var news []*models.News
	var count int64
	val := nq.dbRepo.DB.Debug().Where("author_id = ?", id).Preload("Author").Order("created_at desc").Find(&news).Count(&count).Error
	if val != nil {
		return nil, 0, val
	}
	return news, count, nil
}
