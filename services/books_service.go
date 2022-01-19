package services

import (
	"fmt"
	"log"
	"maranatha_web/domain/books"
	"maranatha_web/models"
	"maranatha_web/utils/errors"
	"time"
)

var (
	BooksService booksServiceInterface = &booksService{}
)

type booksService struct{}

type booksServiceInterface interface {
	CreateBooksPost(booksModel models.Books) (*models.Books, *errors.RestErr)
	GetAllBooks() ([]*models.Books, int64, *errors.RestErr)
	DeleteBooksPost(id uint) *errors.RestErr
	GetSingleBooksPost(id uint) (*models.Books, *errors.RestErr)
	UpdateBooksPost(id uint, newModel models.Books) *errors.RestErr
}

func (b *booksService) CreateBooksPost(booksModel models.Books) (*models.Books, *errors.RestErr) {
	if err := books.CreateBook(&booksModel); err != nil {
		return nil, err
	}
	return &booksModel, nil
}

func (b *booksService) GetAllBooks() ([]*models.Books, int64, *errors.RestErr) {
	data, count, err := books.GetAllBook()
	for _, v := range data {
		v.File = fmt.Sprint("http://localhost:9000/mono/%s", v.File)
		d := v.CreatedAt.Format(time.RFC822)
		myDate, err := time.Parse(time.RFC822, d)
		if err != nil {
			panic(err)
		}

		v.CreatedAt = myDate
		fmt.Println(v.CreatedAt.Format(time.RFC1123))
	}

	if err != nil {
		return data, count, errors.NewBadRequestError("Could not get post")

	}

	return data, count, nil
}

func (b *booksService) DeleteBooksPost(id uint) *errors.RestErr {
	err := books.DeleteBook(id)
	if err != nil {
		log.Println(err)
		return errors.NewBadRequestError("Could not delete item")
	}
	return nil
}

func (b *booksService) GetSingleBooksPost(id uint) (*models.Books, *errors.RestErr) {
	books, err := books.GetSingleBook(id)
	url := fmt.Sprintf("http://localhost:9000/mono/%s", books.File)
	books.File = url
	if err != nil {
		log.Println(err)
		return books, errors.NewBadRequestError("Could not get single post")
	}
	return books, nil
}

func (b *booksService) UpdateBooksPost(id uint, booksModel models.Books) *errors.RestErr {
	books, err := books.UpdateBook(id, booksModel)
	url := fmt.Sprintf("http://localhost:9000/mono/%s", books.File)
	books.File = url
	if err != nil {
		log.Println(err)
		return errors.NewBadRequestError("Could not get single book")
	}
	return nil
}
