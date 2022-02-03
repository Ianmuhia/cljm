package repository

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"log"
	"maranatha_web/internal/models"
)

type PrayerRequestQuery interface {
	CreatePrayerRequest(prayer *models.Prayer) error
	DeletePrayerRequest(id uint) error
	GetSinglePrayerRequest(id uint) (*models.Prayer, error)
	UpdatePrayerRequest(id uint, prayerModel models.Prayer) (*models.Prayer, error)
	GetAllPrayerRequests() ([]*models.Prayer, int64, error)
	GetAllPrayerRequestsByAuthor(id uint) ([]*models.Prayer, int64, error)
}

type prayerRequestQuery struct {
	repo postgresDBRepo
}

func (prq *prayerRequestQuery) CreatePrayerRequest(prayer *models.Prayer) error {
	err := prq.repo.DB.Debug().Create(&prayer).Error
	if err != nil {
		prq.repo.App.ErrorLog.Error("error when trying to save prayer", zap.Any("error", err))
		return err
	}
	return nil
}
func (prq *prayerRequestQuery) DeletePrayerRequest(id uint) error {
	var prayer models.Prayer
	err := prq.repo.DB.Debug().Where("id = ?", id).Delete(&prayer).Error
	if err != nil {
		prq.repo.App.ErrorLog.Error("error when trying to delete prayer post", zap.Any("error", err))
		return err
	}
	return nil
}
func (prq *prayerRequestQuery) GetSinglePrayerRequest(id uint) (*models.Prayer, error) {
	var prayer models.Prayer
	err := prq.repo.DB.Debug().Preload(clause.Associations).Where("id = ?", id).First(&prayer).Error
	if err != nil || err == gorm.ErrRecordNotFound {
		prq.repo.App.ErrorLog.Error("error when trying to get prayer post", zap.Any("error", err))
		return &prayer, err
	}
	return &prayer, nil
}
func (prq *prayerRequestQuery) UpdatePrayerRequest(id uint, prayerModel models.Prayer) (*models.Prayer, error) {
	err := prq.repo.DB.Debug().Where("id = ?", id).Updates(&prayerModel).Error
	if err != nil || err == gorm.ErrRecordNotFound {
		prq.repo.App.ErrorLog.Error("error when trying to update prayer post", zap.Any("error", err))
		return &prayerModel, err
	}
	return &prayerModel, nil
}
func (prq *prayerRequestQuery) GetAllPrayerRequests() ([]*models.Prayer, int64, error) {
	var prayer []*models.Prayer
	var count int64
	val := prq.repo.DB.Debug().Preload(clause.Associations).Order("created_at desc").Find(&prayer).Error
	if val != nil {
		log.Println(val)
		return nil, 0, val
	}
	count = int64(len(prayer))
	return prayer, count, nil
}
func (prq *prayerRequestQuery) GetAllPrayerRequestsByAuthor(id uint) ([]*models.Prayer, int64, error) {
	var prayer []*models.Prayer
	var count int64
	val := prq.repo.DB.Debug().Where("author_id = ?", id).Preload("Author").Order("created_at desc").Find(&prayer).Count(&count).Error
	log.Println(val)
	if val != nil {
		log.Println(val)
		return nil, 0, val
	}
	return prayer, count, nil
}
