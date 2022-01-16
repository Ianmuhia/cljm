package church_patners

import (
	"log"

	"gorm.io/gorm" //nolint:goimports
	//nolint:goimports
	PostgreSQL "maranatha_web/datasources/postgresql"
	"maranatha_web/logger"
	"maranatha_web/models"
	"maranatha_web/utils/errors" //nolint:goimports
)

func CreateChurchPartner(paetner *models.ChurchPartner) *errors.RestErr {
	err := PostgreSQL.Client.Debug().Create(&paetner).Error
	if err != nil {
		logger.Error("error when trying to save church partner", err)
		return errors.NewInternalServerError("database error")
	}
	return nil
}

func DeleteChurchPartner(id uint) *errors.RestErr {
	var partner models.ChurchPartner
	err := PostgreSQL.Client.Debug().Where("id = ?", id).Delete(&partner).Error
	if err != nil {
		logger.Error("error when trying to delete church partner", err)
		return errors.NewInternalServerError("database error")
	}
	return nil
}
func GetSingleChurchPartner(id uint) (*models.ChurchPartner, *errors.RestErr) {
	var partner models.ChurchPartner
	err := PostgreSQL.Client.Debug().Where("id = ?", id).First(&partner).Error
	if err != nil || err == gorm.ErrRecordNotFound {
		logger.Error("error when trying to get  partner post", err)
		return &partner, errors.NewInternalServerError("database error")
	}
	return &partner, nil
}
func UpdateChurchPartner(id uint, partnerModel models.ChurchPartner) (*models.ChurchPartner, *errors.RestErr) {
	err := PostgreSQL.Client.Debug().Where("id = ?", id).Updates(&partnerModel).Error
	if err != nil || err == gorm.ErrRecordNotFound {
		logger.Error("error when trying to update  partner", err)
		return &partnerModel, errors.NewInternalServerError("database error")
	}
	return &partnerModel, nil
}

func GetAllChurchPartner() ([]models.ChurchPartner, int64, error) {
	var churchPartners []models.ChurchPartner
	var count int64
	val := PostgreSQL.Client.Debug().Order("created_at desc").Find(&churchPartners).Error
	if val != nil {
		log.Println(val)
		return nil, 0, val
	}
	return churchPartners, count, nil
}
