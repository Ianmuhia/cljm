package services

import (
	"fmt"
	"log"
	"maranatha_web/internal/models"
	"maranatha_web/internal/repository"
	"time"
)

type GenreServiceInterface interface {
	CreateGenrePost(genreModel models.Genre) (*models.Genre, error)
	GetAllGenres() ([]*models.Genre, int64, error)
	DeleteGenre(id uint) error
	GetSingleGenre(name string) (*models.Genre, error)
	UpdateGenre(id uint, genreModel models.Genre) error
}

type genreService struct {
	dao repository.DAO
}

func NewGenreService(dao repository.DAO) GenreServiceInterface {
	return &genreService{dao: dao}
}

func (gs *genreService) CreateGenrePost(genreModel models.Genre) (*models.Genre, error) {
	if err := gs.dao.NewGenresQuery().CreateGenrePost(&genreModel); err != nil {
		return nil, err
	}
	return &genreModel, nil
}

func (gs *genreService) GetAllGenres() ([]*models.Genre, int64, error) {
	data, count, err := gs.dao.NewGenresQuery().GetAllGenres()
	for _, v := range data {
		d := v.CreatedAt.Format(time.RFC822)

		myDate, err := time.Parse(time.RFC822, d)
		if err != nil {
			panic(err)
		}

		v.CreatedAt = myDate
		fmt.Println(v.CreatedAt.Format(time.RFC1123))
	}
	if err != nil {
		return data, count, err

	}

	return data, count, nil
}

func (gs *genreService) DeleteGenre(id uint) error {
	err := gs.dao.NewGenresQuery().DeleteGenre(id)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (gs *genreService) GetSingleGenre(name string) (*models.Genre, error) {
	genre, err := gs.dao.NewGenresQuery().GetSingleGenre(name)
	if err != nil {
		log.Println(err)
		return genre, err
	}
	return genre, nil
}

func (gs *genreService) UpdateGenre(id uint, genreModel models.Genre) error {
	_, err := gs.dao.NewGenresQuery().UpdateGenre(id, genreModel)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
