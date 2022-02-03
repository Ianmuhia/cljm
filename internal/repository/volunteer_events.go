package repository

import (
	"maranatha_web/internal/models"
)

type VolunteerQuery interface {
}

type volunteerQuery struct {
	repo postgresDBRepo
}

func (vq *volunteerQuery) CreateSubscribeToChurchJob(volunteerChurchJob *models.VolunteerChurchJob) error {
	err := vq.repo.DB.Debug().Create(&volunteerChurchJob).Error
	if err != nil {
		vq.repo.App.ErrorLog.Error("Error when trying to save volunteerChurchJob")
		return err
	}
	return nil
}
