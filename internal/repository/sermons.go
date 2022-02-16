package repository

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
	"log"
	"maranatha_web/internal/models"
)

type SermonQuery interface {
	CreateSermon(sermon *models.Sermon) error
	DeleteSermon(id uint) error
	GetSingleSermon(id uint) (*models.Sermon, error)
	UpdateSermon(id uint, sermonModel models.Sermon) (*models.Sermon, error)
	GetAllSermon() ([]*models.Sermon, int64, error)
}

type sermonQuery struct {
	repo postgresDBRepo
}

func (sq *sermonQuery) CreateSermon(sermon *models.Sermon) error {
	err := sq.repo.DB.Debug().Create(&sermon).Error
	if err != nil {
		sq.repo.App.ErrorLog.Error("error when trying to save sermon", zap.Any("error", err))
		return err
	}
	return nil
}
func (sq *sermonQuery) DeleteSermon(id uint) error {
	var sermon models.Sermon
	err := sq.repo.DB.Debug().Where("id = ?", id).Delete(&sermon).Error
	if err != nil {
		sq.repo.App.ErrorLog.Error("error when trying to delete sermon", zap.Any("error", err))
		return err
	}
	return nil
}
func (sq *sermonQuery) GetSingleSermon(id uint) (*models.Sermon, error) {
	var sermon models.Sermon
	err := sq.repo.DB.Debug().Where("id = ?", id).First(&sermon).Error
	if err != nil || err == gorm.ErrRecordNotFound {
		sq.repo.App.ErrorLog.Error("error when trying to get  sermon ", zap.Any("error", err))
		return &sermon, err
	}
	return &sermon, nil
}
func (sq *sermonQuery) UpdateSermon(id uint, sermonModel models.Sermon) (*models.Sermon, error) {
	err := sq.repo.DB.Debug().Where("id = ?", id).Updates(&sermonModel).Error
	if err != nil || err == gorm.ErrRecordNotFound {
		sq.repo.App.ErrorLog.Error("error when trying to update  partner", zap.Any("error", err))
		return &sermonModel, err
	}
	return &sermonModel, nil
}
func (sq *sermonQuery) GetAllSermon() ([]*models.Sermon, int64, error) {
	var sermons []*models.Sermon
	var count int64
	val := sq.repo.DB.Debug().Order("created_at desc").Find(&sermons).Error
	if val != nil {
		log.Println(val)
		return nil, 0, val
	}
	return sermons, count, nil
}
