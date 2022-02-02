package repository

// import (
// 	"log"
// 	//nolint:goimports
//

// 	"gorm.io/gorm" //nolint:goimports

// 	PostgreSQL "maranatha_web/datasources/postgresql"
// 	"maranatha_web/logger"
// 	"maranatha_web/models"
// 	"maranatha_web/utils/errors" //nolint:goimports
// )

// func CreateSermon(sermon *models.Sermon) *errors.RestErr {
// 	err := PostgreSQL.Client.Debug().Create(&sermon).Error
// 	if err != nil {
// 		logger.Error("error when trying to save sermon", err)
// 		return errors.NewInternalServerError("database error")
// 	}
// 	return nil
// }

// func DeleteSermon(id uint) *errors.RestErr {
// 	var sermon models.Sermon
// 	err := PostgreSQL.Client.Debug().Where("id = ?", id).Delete(&sermon).Error
// 	if err != nil {
// 		logger.Error("error when trying to delete sermon", err)
// 		return errors.NewInternalServerError("database error")
// 	}
// 	return nil
// }
// func GetSingleSermon(id uint) (*models.Sermon, *errors.RestErr) {
// 	var sermon models.Sermon
// 	err := PostgreSQL.Client.Debug().Where("id = ?", id).First(&sermon).Error
// 	if err != nil || err == gorm.ErrRecordNotFound {
// 		logger.Error("error when trying to get  sermon ", err)
// 		return &sermon, errors.NewInternalServerError("database error")
// 	}
// 	return &sermon, nil
// }
// func UpdateSermon(id uint, sermonModel models.Sermon) (*models.Sermon, *errors.RestErr) {
// 	err := PostgreSQL.Client.Debug().Where("id = ?", id).Updates(&sermonModel).Error
// 	if err != nil || err == gorm.ErrRecordNotFound {
// 		logger.Error("error when trying to update  partner", err)
// 		return &sermonModel, errors.NewInternalServerError("database error")
// 	}
// 	return &sermonModel, nil
// }

// func GetAllSermon() ([]models.Sermon, int64, error) {
// 	var sermons []models.Sermon
// 	var count int64
// 	val := PostgreSQL.Client.Debug().Order("created_at desc").Find(&sermons).Error
// 	if val != nil {
// 		log.Println(val)
// 		return nil, 0, val
// 	}
// 	return sermons, count, nil
// }
