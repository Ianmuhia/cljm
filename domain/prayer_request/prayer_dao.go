package prayer_request

import (
	"log"
	"maranatha_web/logger"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	postgresql_db "maranatha_web/datasources/postgresql"
	"maranatha_web/models"
	"maranatha_web/utils/errors"
)

func CreatePrayerPost(prayer *models.Prayer) *errors.RestErr {
	err := postgresql_db.Client.Debug().Create(&prayer).Error
	if err != nil {
		logger.Error("error when trying to save prayer", err)
		return errors.NewInternalServerError("database error")
	}
	return nil
}

func DeletePrayerPost(id uint) *errors.RestErr {
	var prayer models.Prayer
	err := postgresql_db.Client.Debug().Where("id = ?", id).Delete(&prayer).Error
	if err != nil {
		logger.Error("error when trying to delete prayer post", err)
		return errors.NewInternalServerError("database error")
	}
	return nil
}
func GetSinglePrayerPost(id uint) (*models.Prayer, *errors.RestErr) {
	var prayer models.Prayer
	err := postgresql_db.Client.Debug().Preload(clause.Associations).Where("id = ?", id).First(&prayer).Error
	if err != nil || err == gorm.ErrRecordNotFound {
		logger.Error("error when trying to get prayer post", err)
		return &prayer, errors.NewInternalServerError("database error")
	}
	return &prayer, nil
}
func UpdatePrayerPost(id uint, prayerModel models.Prayer) (*models.Prayer, *errors.RestErr) {
	err := postgresql_db.Client.Debug().Where("id = ?", id).Updates(&prayerModel).Error
	if err != nil || err == gorm.ErrRecordNotFound {
		logger.Error("error when trying to update prayer post", err)
		return &prayerModel, errors.NewInternalServerError("database error")
	}
	return &prayerModel, nil
}
func GetAllPrayerPost() ([]*models.Prayer, int64, error) {
	var prayer []*models.Prayer
	var count int64
	val := postgresql_db.Client.Debug().Preload(clause.Associations).Order("created_at desc").Find(&prayer).Error
	if val != nil {
		log.Println(val)
		return nil, 0, val
	}
	return prayer, count, nil
}
func GetAllPrayerPostByAuthor(id uint) ([]*models.Prayer, int64, error) {
	var prayer []*models.Prayer
	var count int64
	val := postgresql_db.Client.Debug().Where("author_id = ?", id).Preload("Author").Order("created_at desc").Find(&prayer).Count(&count).Error
	if val != nil {
		log.Println(val)
		return nil, 0, val
	}
	return prayer, count, nil
}