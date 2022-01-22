package volunteer_events

import (
	postgresql_db "maranatha_web/datasources/postgresql"
	"maranatha_web/logger"
	"maranatha_web/models"
	"maranatha_web/utils/errors"
)

func CreateSubscribeToChurchJob(volunteerChurchJob *models.VolunteerChurchJob) *errors.RestErr {
	err := postgresql_db.Client.Debug().Create(&volunteerChurchJob).Error
	if err != nil {
		return errors.NewInternalServerError("Error when trying to save volunteerChurchJob")
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
