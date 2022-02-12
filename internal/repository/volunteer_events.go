package repository

import (
	"gorm.io/gorm/clause"
	"maranatha_web/internal/models"
)

type VolunteerQuery interface {
	CreateSubscribeToChurchJob(volunteerChurchJob *models.VolunteerChurchJob) error
	GetUserVolunteeredJobs(id int) (int, []*models.VolunteerChurchJob, error)
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

func (vq *volunteerQuery) GetUserVolunteeredJobs(id int) (int, []*models.VolunteerChurchJob, error) {
	var vj []*models.VolunteerChurchJob
	var total int
	err := vq.repo.DB.Debug().Preload(clause.Associations).Where("volunteer_id = ?", id).Find(&vj).Error
	if err != nil {
		vq.repo.App.ErrorLog.Error("Error when trying to save volunteerChurchJob")
		return total, vj, err
	}
	total = len(vj)
	return total, vj, nil
}
