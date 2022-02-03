package repository

import (
	"go.uber.org/zap"
	"maranatha_web/internal/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type GenresQuery interface {
	CreateGenrePost(genre *models.Genre) error
	DeleteGenre(id uint) error
	GetSingleGenre(name string) (*models.Genre, error)
	UpdateGenre(id uint, genreModel models.Genre) (*models.Genre, error)
	GetAllGenres() ([]*models.Genre, int64, error)
	GetAllGenresByAuthor(id uint) ([]*models.Genre, int64, error)
}

type genresQuery struct {
	dbRepo postgresDBRepo
}

func (gq *genresQuery) CreateGenrePost(genre *models.Genre) error {
	err := gq.dbRepo.DB.Debug().Create(&genre).Error
	if err != nil {
		gq.dbRepo.App.ErrorLog.Error("error when trying to save genre", zap.Any("error", err))
		return err
	}
	return nil
}

func (gq *genresQuery) DeleteGenre(id uint) error {
	var genre models.Genre
	err := gq.dbRepo.DB.Debug().Where("id = ?", id).Delete(&genre).Error
	if err != nil {
		gq.dbRepo.App.ErrorLog.Error("error when trying to delete genre post", zap.Any("error", err))
		return err
	}
	return nil
}
func (gq *genresQuery) GetSingleGenre(name string) (*models.Genre, error) {
	var genre models.Genre
	err := gq.dbRepo.DB.Debug().Preload(clause.Associations).Where("name = ?", name).First(&genre).Error
	if err != nil || err == gorm.ErrRecordNotFound {
		gq.dbRepo.App.ErrorLog.Error("error when trying to get genre post", zap.Any("error", err))
		return &genre, err
	}
	return &genre, nil
}
func (gq *genresQuery) UpdateGenre(id uint, genreModel models.Genre) (*models.Genre, error) {
	err := gq.dbRepo.DB.Debug().Where("id = ?", id).Updates(&genreModel).Error
	if err != nil || err == gorm.ErrRecordNotFound {
		gq.dbRepo.App.ErrorLog.Error("error when trying to update genre post", zap.Any("error", err))
		return &genreModel, err
	}
	return &genreModel, nil
}
func (gq *genresQuery) GetAllGenres() ([]*models.Genre, int64, error) {
	var genre []*models.Genre
	var count int64
	val := gq.dbRepo.DB.Debug().Preload(clause.Associations).Order("created_at desc").Find(&genre).Count(&count).Error
	if val != nil {

		return nil, 0, val
	}
	return genre, count, nil
}
func (gq *genresQuery) GetAllGenresByAuthor(id uint) ([]*models.Genre, int64, error) {
	var genre []*models.Genre
	var count int64
	val := gq.dbRepo.DB.Debug().Where("author_id = ?", id).Preload("Author").Order("created_at desc").Find(&genre).Count(&count).Error
	if val != nil {
		return nil, 0, val
	}
	return genre, count, nil
}
