package testimonies

import (
	"log"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	postgresql_db "maranatha_web/datasources/postgresql"
	"maranatha_web/logger"
	"maranatha_web/models"
	"maranatha_web/utils/errors"
)

func CreateTestimony(testimonies *models.Testimonies) *errors.RestErr {
	err := postgresql_db.Client.Debug().Create(&testimonies).Error
	if err != nil {
		logger.Error("error when trying to save testimony", err)
		return errors.NewInternalServerError("database error")
	}
	return nil
}

func DeleteTestimony(id uint) *errors.RestErr {
	var testimonies models.Testimonies
	err := postgresql_db.Client.Debug().Where("id = ?", id).Delete(&testimonies).Error
	if err != nil {
		logger.Error("error when trying to delete testimony post", err)
		return errors.NewInternalServerError("database error")
	}
	return nil
}
func GetSingleTestimony(id uint) (*models.Testimonies, *errors.RestErr) {
	var testimonies models.Testimonies
	err := postgresql_db.Client.Debug().Preload(clause.Associations).Where("id = ?", id).First(&testimonies).Error
	if err != nil || err == gorm.ErrRecordNotFound {
		logger.Error("error when trying to get testimonies post", err)
		return &testimonies, errors.NewInternalServerError("database error")
	}
	return &testimonies, nil
}
func UpdateTestimony(id uint, testimoniesModel models.Testimonies) (*models.Testimonies, *errors.RestErr) {
	err := postgresql_db.Client.Debug().Where("id = ?", id).Updates(&testimoniesModel).Error
	if err != nil || err == gorm.ErrRecordNotFound {
		logger.Error("error when trying to update testimonies post", err)
		return &testimoniesModel, errors.NewInternalServerError("database error")
	}
	return &testimoniesModel, nil
}
func GetAllTestimonies() ([]*models.Testimonies, int64, error) {
	var testimonies []*models.Testimonies
	var count int64
	val := postgresql_db.Client.Debug().Preload(clause.Associations).Order("created_at desc").Find(&testimonies).Error
	count = int64(len(testimonies))
	if val != nil {
		log.Println(val)
		return nil, 0, val
	}
	return testimonies, count, nil
}
func GetAllTestimoniesByAuthor(id uint) ([]*models.Testimonies, int64, error) {
	var testimonies []*models.Testimonies
	var count int64
	val := postgresql_db.Client.Debug().Where("author_id = ?", id).Preload("Author").Order("created_at desc").Find(&testimonies).Count(&count).Error
	if val != nil {
		log.Println(val)
		return nil, 0, val
	}
	return testimonies, count, nil
}
