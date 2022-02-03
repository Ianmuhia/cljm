package repository

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"log"
	"maranatha_web/internal/models"
)

type TestimonyQuery interface {
	CreateTestimony(testimonies *models.Testimonies) error
	DeleteTestimony(id uint) error
	GetSingleTestimony(id uint) (*models.Testimonies, error)
	UpdateTestimony(id uint, testimoniesModel models.Testimonies) (*models.Testimonies, error)
	GetAllTestimonies() ([]*models.Testimonies, int64, error)
	GetAllTestimoniesByAuthor(id uint) ([]*models.Testimonies, int64, error)
}

type testimonyQuery struct {
	repo postgresDBRepo
}

func (tq *testimonyQuery) CreateTestimony(testimonies *models.Testimonies) error {
	err := tq.repo.DB.Debug().Create(&testimonies).Error
	if err != nil {
		tq.repo.App.ErrorLog.Error("error when trying to save testimony", zap.Any("error", err))
		return err
	}
	return nil
}

func (tq *testimonyQuery) DeleteTestimony(id uint) error {
	var testimonies models.Testimonies
	err := tq.repo.DB.Debug().Where("id = ?", id).Delete(&testimonies).Error
	if err != nil {
		tq.repo.App.ErrorLog.Error("error when trying to delete testimony post", zap.Any("error", err))
		return err
	}
	return nil
}
func (tq *testimonyQuery) GetSingleTestimony(id uint) (*models.Testimonies, error) {
	var testimonies models.Testimonies
	err := tq.repo.DB.Debug().Preload(clause.Associations).Where("id = ?", id).First(&testimonies).Error
	if err != nil || err == gorm.ErrRecordNotFound {
		tq.repo.App.ErrorLog.Error("error when trying to get testimonies post", zap.Any("error", err))
		return &testimonies, err
	}
	return &testimonies, nil
}
func (tq *testimonyQuery) UpdateTestimony(id uint, testimoniesModel models.Testimonies) (*models.Testimonies, error) {
	err := tq.repo.DB.Debug().Where("id = ?", id).Updates(&testimoniesModel).Error
	if err != nil || err == gorm.ErrRecordNotFound {
		tq.repo.App.ErrorLog.Error("error when trying to update testimonies post", zap.Any("error", err))
		return &testimoniesModel, err
	}
	return &testimoniesModel, nil
}
func (tq *testimonyQuery) GetAllTestimonies() ([]*models.Testimonies, int64, error) {
	var testimonies []*models.Testimonies
	var count int64
	val := tq.repo.DB.Debug().Preload(clause.Associations).Order("created_at desc").Find(&testimonies).Error
	count = int64(len(testimonies))
	if val != nil {
		log.Println(val)
		return nil, 0, val
	}
	return testimonies, count, nil
}
func (tq *testimonyQuery) GetAllTestimoniesByAuthor(id uint) ([]*models.Testimonies, int64, error) {
	var testimonies []*models.Testimonies
	var count int64
	val := tq.repo.DB.Debug().Where("author_id = ?", id).Preload("Author").Order("created_at desc").Find(&testimonies).Count(&count).Error
	if val != nil {
		log.Println(val)
		return nil, 0, val
	}
	return testimonies, count, nil
}
