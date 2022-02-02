package services

// import (
// 	"fmt"
// 	"log"
// 	"time"
//

// 	"maranatha_web/domain/genre"
// 	"maranatha_web/models"
// 	"maranatha_web/utils/errors"
// )

// var (
// 	GenreService genreServiceInterface = &genreService{}
// )

// type genreService struct{}

// type genreServiceInterface interface {
// 	CreateGenrePost(genreModel models.Genre) (*models.Genre, *errors.RestErr)
// 	GetAllGenres() ([]*models.Genre, int64, *errors.RestErr)
// 	DeleteGenre(id uint) *errors.RestErr
// 	GetSingleGenre(name string) (*models.Genre, *errors.RestErr)
// 	UpdateGenre(id uint, genreModel models.Genre) *errors.RestErr
// }

// func (s *genreService) CreateGenrePost(genreModel models.Genre) (*models.Genre, *errors.RestErr) {
// 	if err := genre.CreateGenrePost(&genreModel); err != nil {
// 		return nil, err
// 	}
// 	return &genreModel, nil
// }

// func (s *genreService) GetAllGenres() ([]*models.Genre, int64, *errors.RestErr) {
// 	data, count, err := genre.GetAllGenres()
// 	for _, v := range data {
// 		d := v.CreatedAt.Format(time.RFC822)

// 		myDate, err := time.Parse(time.RFC822, d)
// 		if err != nil {
// 			panic(err)
// 		}

// 		v.CreatedAt = myDate
// 		fmt.Println(v.CreatedAt.Format(time.RFC1123))
// 	}
// 	if err != nil {
// 		return data, count, errors.NewBadRequestError("Could not get genres")

// 	}

// 	return data, count, nil
// }

// func (s *genreService) DeleteGenre(id uint) *errors.RestErr {
// 	err := genre.DeleteGenre(id)
// 	if err != nil {
// 		log.Println(err)
// 		return errors.NewBadRequestError("Could not delete genre")
// 	}
// 	return nil
// }

// func (s *genreService) GetSingleGenre(name string) (*models.Genre, *errors.RestErr) {
// 	genre, err := genre.GetSingleGenre(name)
// 	if err != nil {
// 		log.Println(err)
// 		return genre, err
// 	}
// 	return genre, nil
// }

// func (s *genreService) UpdateGenre(id uint, genreModel models.Genre) *errors.RestErr {
// 	_, err := genre.UpdateGenre(id, genreModel)
// 	if err != nil {
// 		log.Println(err)
// 		return errors.NewBadRequestError("Could not get single genre")
// 	}
// 	return nil
// }
