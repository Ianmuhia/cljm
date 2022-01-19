package services

import (
	"fmt"
	"log"
	"maranatha_web/domain/genre"
	"maranatha_web/models"
	"maranatha_web/utils/errors"
	"time"
)

var (
	GenreService genreServiceInterface = &genreService{}
)

type genreService struct{}

type genreServiceInterface interface {
	CreateGenrePost(genreModel models.Genre) (*models.Genre, *errors.RestErr)
	GetAllGenrePost() ([]*models.Genre, int64, *errors.RestErr)
	DeleteGenrePost(id uint) *errors.RestErr
	GetSingleGenrePost(id uint) (*models.Genre, *errors.RestErr)
	UpdateGenrePost(id uint, genreModel models.Genre) *errors.RestErr
}

func (s *genreService) CreateGenrePost(genreModel models.Genre) (*models.Genre, *errors.RestErr) {
	if err := genre.CreateGenrePost(&genreModel); err != nil {
		return nil, err
	}
	return &genreModel, nil
}

func (s *genreService) GetAllGenrePost() ([]*models.Genre, int64, *errors.RestErr) {
	data, count, err := genre.GetAllGenrePost()
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
		return data, count, errors.NewBadRequestError("Could not get genres")

	}

	return data, count, nil
}

func (s *genreService) DeleteGenrePost(id uint) *errors.RestErr {
	err := genre.DeleteGenrePost(id)
	if err != nil {
		log.Println(err)
		return errors.NewBadRequestError("Could not delete genre")
	}
	return nil
}

func (s *genreService) GetSingleGenrePost(id uint) (*models.Genre, *errors.RestErr) {
	genre, err := genre.GetSingleGenrePost(id)
	if err != nil {
		log.Println(err)
		return genre, errors.NewBadRequestError("Could not get single genre")
	}
	return genre, nil
}

func (s *genreService) UpdateGenrePost(id uint, genreModel models.Genre) *errors.RestErr {
	_, err := genre.UpdateGenrePost(id, genreModel)
	if err != nil {
		log.Println(err)
		return errors.NewBadRequestError("Could not get single genre")
	}
	return nil
}
