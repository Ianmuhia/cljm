package services

import (
	"fmt"
	"log"
	"maranatha_web/internal/repository"
	"time"

	"maranatha_web/internal/models"
)

type BooksService interface {
	CreateBooksPost(booksModel models.Books) (*models.Books, error)
	GetAllBooks() ([]*models.Books, int64, error)
	DeleteBook(id uint) error
	GetSingleBooksPost(id uint) (*models.Books, error)
	UpdateBooksPost(id uint, newModel models.Books) error
}

type booksService struct {
	dao repository.DAO
}

func NewBookService(dao repository.DAO) BooksService {
	return &booksService{dao: dao}
}

func (bs *booksService) CreateBooksPost(booksModel models.Books) (*models.Books, error) {
	if err := bs.dao.NewBookQuery().CreateBook(&booksModel); err != nil {
		return nil, err
	}
	return &booksModel, nil
}

func (bs *booksService) GetAllBooks() ([]*models.Books, int64, error) {
	data, count, err := bs.dao.NewBookQuery().GetAllBook()
	for _, v := range data {
		v.File = fmt.Sprintf("http://localhost:9000/mono/%s", v.File)
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

func (bs *booksService) DeleteBook(id uint) error {
	err := bs.dao.NewBookQuery().DeleteBook(id)
	if err != nil {
		log.Println(err)
		return err
		//  errors.NewBadRequestError("Could not delete item")
	}
	return nil
}

func (bs *booksService) GetSingleBooksPost(id uint) (*models.Books, error) {
	books, err := bs.dao.NewBookQuery().GetSingleBook(id)
	url := fmt.Sprintf("http://localhost:9000/mono/%s", books.File)
	books.File = url
	if err != nil {
		log.Println(err)
		return books, err
		// errors.NewBadRequestError("Could not get single post")
	}
	return books, nil
}

func (bs *booksService) UpdateBooksPost(id uint, booksModel models.Books) error {
	books, err := bs.dao.NewBookQuery().UpdateBook(id, booksModel)
	url := fmt.Sprintf("http://localhost:9000/mono/%s", books.File)
	books.File = url
	if err != nil {
		log.Println(err)
		return err
		//  errors.NewBadRequestError("Could not get single book")
	}
	return nil
}
