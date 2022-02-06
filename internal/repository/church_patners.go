package repository

import (
	"go.uber.org/zap"
	"gorm.io/gorm"

	"maranatha_web/internal/models"
)

type PartnersQuery interface {
	CreateChurchPartner(partner *models.ChurchPartner) error
	DeleteChurchPartner(id uint) error
	GetSingleChurchPartner(id uint) (*models.ChurchPartner, error)
	UpdateChurchPartner(id uint, partnerModel models.ChurchPartner) (*models.ChurchPartner, error)
	GetAllChurchPartner() ([]*models.ChurchPartner, int64, error)
}

type partnerQuery struct {
	dbRepo postgresDBRepo
}

func (pq *partnerQuery) CreateChurchPartner(partner *models.ChurchPartner) error {
	err := pq.dbRepo.DB.Debug().Create(&partner).Error
	if err != nil {
		pq.dbRepo.App.ErrorLog.Error("error when trying to save church partner", zap.Any("error", err))
		return err
	}
	return nil
}

func (pq *partnerQuery) DeleteChurchPartner(id uint) error {
	var partner models.ChurchPartner
	err := pq.dbRepo.DB.Debug().Where("id = ?", id).Delete(&partner).Error
	if err != nil {
		pq.dbRepo.App.ErrorLog.Error("error when trying to delete church partner", zap.Any("error", err))
		return err
	}
	return nil
}
func (pq *partnerQuery) GetSingleChurchPartner(id uint) (*models.ChurchPartner, error) {
	var partner models.ChurchPartner
	err := pq.dbRepo.DB.Debug().Where("id = ?", id).First(&partner).Error
	if err != nil || err == gorm.ErrRecordNotFound {
		pq.dbRepo.App.ErrorLog.Error("error when trying to get  partner post", zap.Any("error", err))
		return &partner, err
	}
	return &partner, nil
}
func (pq *partnerQuery) UpdateChurchPartner(id uint, partnerModel models.ChurchPartner) (*models.ChurchPartner, error) {
	err := pq.dbRepo.DB.Debug().Where("id = ?", id).Updates(&partnerModel).Error
	if err != nil || err == gorm.ErrRecordNotFound {
		pq.dbRepo.App.ErrorLog.Error("error when trying to update  partner", zap.Any("error", err))
		return &partnerModel, err
	}
	return &partnerModel, nil
}

func (pq *partnerQuery) GetAllChurchPartner() ([]*models.ChurchPartner, int64, error) {
	var churchPartners []*models.ChurchPartner
	var count int
	val := pq.dbRepo.DB.Debug().Order("created_at desc").Find(&churchPartners).Error
	if val != nil {
		pq.dbRepo.App.ErrorLog.Error("error when trying to get all  partner", zap.Any("error", val))
		return nil, 0, val
	}
	count = len(churchPartners)

	return churchPartners, int64(count), nil
}
