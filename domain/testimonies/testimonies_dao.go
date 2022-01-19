package testimonies

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"log"
	postgresql_db "maranatha_web/datasources/postgresql"
	"maranatha_web/logger"
	"maranatha_web/models"
	"maranatha_web/utils/errors"
)

func CreateTestimoniesPost(testimonies *models.Testimonies) *errors.RestErr {
	err := postgresql_db.Client.Debug().Create(&testimonies).Error
	if err != nil {
		logger.Error("error when trying to save testimony", err)
		return errors.NewInternalServerError("database error")
	}
	return nil
}

func DeleteTestimoniesPost(id uint) *errors.RestErr {
	var testimonies models.Testimonies
	err := postgresql_db.Client.Debug().Where("id = ?", id).Delete(&testimonies).Error
	if err != nil {
		logger.Error("error when trying to delete testimony post", err)
		return errors.NewInternalServerError("database error")
	}
	return nil
}
func GetSingleTestimoniesPost(id uint) (*models.Testimonies, *errors.RestErr) {
	var testimonies models.Testimonies
	err := postgresql_db.Client.Debug().Preload(clause.Associations).Where("id = ?", id).First(&testimonies).Error
	if err != nil || err == gorm.ErrRecordNotFound {
		logger.Error("error when trying to get testimonies post", err)
		return &testimonies, errors.NewInternalServerError("database error")
	}
	return &testimonies, nil
}
func UpdateTestimoniesPost(id uint, testimoniesModel models.Testimonies) (*models.Testimonies, *errors.RestErr) {
	err := postgresql_db.Client.Debug().Where("id = ?", id).Updates(&testimoniesModel).Error
	if err != nil || err == gorm.ErrRecordNotFound {
		logger.Error("error when trying to update testimonies post", err)
		return &testimoniesModel, errors.NewInternalServerError("database error")
	}
	return &testimoniesModel, nil
}
func GetAllTestimoniesPost() ([]*models.Testimonies, int64, error) {
	var testimonies []*models.Testimonies
	var count int64
	val := postgresql_db.Client.Debug().Preload(clause.Associations).Order("created_at desc").Find(&testimonies).Error
	if val != nil {
		log.Println(val)
		return nil, 0, val
	}
	return testimonies, count, nil
}
func GetAllTestimoniesPostByAuthor(id uint) ([]*models.Testimonies, int64, error) {
	var testimonies []*models.Testimonies
	var count int64
	val := postgresql_db.Client.Debug().Where("author_id = ?", id).Preload("Author").Order("created_at desc").Find(&testimonies).Count(&count).Error
	if val != nil {
		log.Println(val)
		return nil, 0, val
	}
	return testimonies, count, nil
}
