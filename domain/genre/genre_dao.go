package genre

import (
	"log"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	postgresql_db "maranatha_web/datasources/postgresql"
	"maranatha_web/logger"
	"maranatha_web/models"
	"maranatha_web/utils/errors"
)

func CreateGenrePost(genre *models.Genre) *errors.RestErr {
	err := postgresql_db.Client.Debug().Create(&genre).Error
	if err != nil {
		logger.Error("error when trying to save genre", err)
		return errors.NewInternalServerError("database error")
	}
	return nil
}

func DeleteGenre(id uint) *errors.RestErr {
	var genre models.Genre
	err := postgresql_db.Client.Debug().Where("id = ?", id).Delete(&genre).Error
	if err != nil {
		logger.Error("error when trying to delete genre post", err)
		return errors.NewInternalServerError("database error")
	}
	return nil
}
func GetSingleGenre(name string) (*models.Genre, *errors.RestErr) {
	var genre models.Genre
	err := postgresql_db.Client.Debug().Preload(clause.Associations).Where("name = ?", name).First(&genre).Error
	if err != nil || err == gorm.ErrRecordNotFound {
		logger.Error("error when trying to get genre post", err)
		return &genre, errors.NewNotFoundError("The provided genre does not exist.")
	}
	return &genre, nil
}
func UpdateGenre(id uint, genreModel models.Genre) (*models.Genre, *errors.RestErr) {
	err := postgresql_db.Client.Debug().Where("id = ?", id).Updates(&genreModel).Error
	if err != nil || err == gorm.ErrRecordNotFound {
		logger.Error("error when trying to update genre post", err)
		return &genreModel, errors.NewNotFoundError("The requested genre does not exist.")
	}
	return &genreModel, nil
}
func GetAllGenres() ([]*models.Genre, int64, error) {
	var genre []*models.Genre
	var count int64
	val := postgresql_db.Client.Debug().Preload(clause.Associations).Order("created_at desc").Find(&genre).Count(&count).Error
	if val != nil {
		log.Println(val)
		return nil, 0, val
	}
	return genre, count, nil
}
func GetAllGenresByAuthor(id uint) ([]*models.Genre, int64, error) {
	var genre []*models.Genre
	var count int64
	val := postgresql_db.Client.Debug().Where("author_id = ?", id).Preload("Author").Order("created_at desc").Find(&genre).Count(&count).Error
	if val != nil {
		log.Println(val)
		return nil, 0, val
	}
	return genre, count, nil
}
