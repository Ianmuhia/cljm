package books

import (
	"log"

	"gorm.io/gorm"        //nolint:goimports
	"gorm.io/gorm/clause" //nolint:goimports
	postgresql_db "maranatha_web/datasources/postgresql"
	"maranatha_web/logger"
	"maranatha_web/models"
	"maranatha_web/utils/errors"
)

func CreateBook(book *models.Books) *errors.RestErr {
	err := postgresql_db.Client.Debug().Create(&book).Error
	if err != nil {
		logger.Error("error when trying to save user", err)
		return errors.NewInternalServerError("database error")
	}
	return nil
}

func DeleteBook(id uint) *errors.RestErr {
	var book models.Books
	err := postgresql_db.Client.Debug().Where("id = ?", id).Delete(&book).Error
	if err != nil {
		logger.Error("error when trying to delete book post", err)
		return errors.NewInternalServerError("database error")
	}
	return nil
}
func GetSingleBook(id uint) (*models.Books, *errors.RestErr) {
	var book models.Books
	err := postgresql_db.Client.Debug().Preload(clause.Associations).Where("id = ?", id).First(&book).Error
	if err != nil || err == gorm.ErrRecordNotFound {
		logger.Error("error when trying to get  book post", err)
		return &book, errors.NewInternalServerError("database error")
	}
	return &book, nil
}
func UpdateBook(id uint, bookModel models.Books) (*models.Books, *errors.RestErr) {
	err := postgresql_db.Client.Debug().Where("id = ?", id).Updates(&bookModel).Error
	if err != nil || err == gorm.ErrRecordNotFound {
		logger.Error("error when trying to update  book post", err)
		return &bookModel, errors.NewInternalServerError("database error")
	}
	return &bookModel, nil
}

func GetAllBook() ([]*models.Books, int64, error) {
	var book []*models.Books
	var count int64
	val := postgresql_db.Client.Debug().Preload(clause.Associations).Order("created_at desc").Find(&book).Error
	if val != nil {
		log.Println(val)
		return nil, 0, val
	}
	return book, count, nil
}

func GetAllBookByAuthor(id uint) ([]*models.Books, int64, error) {
	var book []*models.Books
	var count int64
	val := postgresql_db.Client.Debug().Where("author_id = ?", id).Preload("Author").Order("created_at desc").Find(&book).Count(&count).Error
	if val != nil {
		log.Println(val)
		return nil, 0, val
	}
	return book, count, nil
}
