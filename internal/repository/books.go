package repository

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
	"log"
	"maranatha_web/internal/models"

	"gorm.io/gorm/clause" //nolint:goimports
)

type BookQuery interface {
	CreateBook(book *models.Books) error
	DeleteBook(id uint) error
	GetSingleBook(id uint) (*models.Books, error)
	UpdateBook(id uint, bookModel models.Books) (*models.Books, error)
	GetAllBook() ([]*models.Books, int64, error)
	GetAllBookByAuthor(id uint) ([]*models.Books, int64, error)
}

type bookQuery struct {
	dbRepo postgresDBRepo
}

func (bq *bookQuery) CreateBook(book *models.Books) error {
	err := bq.dbRepo.DB.Debug().Create(&book).Error
	if err != nil {
		//bq.dbRepo.App.ErrorLog.Println("error when trying to save user", err)
		return err
	}
	return nil
}

func (bq *bookQuery) DeleteBook(id uint) error {
	var book models.Books
	err := bq.dbRepo.DB.Debug().Where("id = ?", id).Delete(&book).Error
	if err != nil {
		bq.dbRepo.App.ErrorLog.Info("error when trying to delete book post", zap.Any("error", err))
		return err
	}
	return nil
}
func (bq *bookQuery) GetSingleBook(id uint) (*models.Books, error) {
	var book models.Books
	err := bq.dbRepo.DB.Debug().Preload(clause.Associations).Where("id = ?", id).First(&book).Error
	if err != nil || err == gorm.ErrRecordNotFound {
		bq.dbRepo.App.ErrorLog.Info("error when trying to get  book post", zap.Any("error", err))
		return &book, err
	}
	return &book, nil
}
func (bq *bookQuery) UpdateBook(id uint, bookModel models.Books) (*models.Books, error) {
	err := bq.dbRepo.DB.Debug().Where("id = ?", id).Updates(&bookModel).Error
	if err != nil || err == gorm.ErrRecordNotFound {

		bq.dbRepo.App.ErrorLog.Error("error when trying to update  book post %v")
		return &bookModel, err
	}
	return &bookModel, nil
}

func (bq *bookQuery) GetAllBook() ([]*models.Books, int64, error) {
	var book []*models.Books
	var count int64
	val := bq.dbRepo.DB.Debug().Preload(clause.Associations).Order("created_at desc").Find(&book).Error
	count = int64(len(book))

	bq.dbRepo.App.ErrorLog.Info("errr", zap.Any("e", val))
	if val != nil {
		log.Println(val)
		return nil, 0, val
	}
	return book, count, nil
}

func (bq *bookQuery) GetAllBookByAuthor(id uint) ([]*models.Books, int64, error) {
	var book []*models.Books
	var count int64
	val := bq.dbRepo.DB.Debug().Where("author_id = ?", id).Preload("Author").Order("created_at desc").Find(&book).Count(&count).Error
	if val != nil {
		log.Println(val)
		return nil, 0, val
	}
	return book, count, nil
}
