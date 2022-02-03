package repository

import (
	"go.uber.org/zap"
	"log"
	"maranatha_web/internal/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type JobsQuery interface {
	CreateJob(job *models.ChurchJob) error
	DeleteJob(id uint) error
	GetJobByEvent(eventId uint, jobId uint) (*models.ChurchJob, error)
	GetAllEventsJobs(id uint) ([]*models.ChurchJob, int64, error)
	UpdateJob(id uint, job models.ChurchJob) (*models.ChurchJob, error)
	GetSingleJob(id uint) (*models.ChurchJob, error)
}

type jobsQuery struct {
	dbRepo postgresDBRepo
}

func (jq *jobsQuery) CreateJob(job *models.ChurchJob) error {
	err := jq.dbRepo.DB.Debug().Create(&job).Error
	if err != nil {
		jq.dbRepo.App.ErrorLog.Error("error when trying to create church event.", zap.Any("error", err))
		return err
	}
	return nil
}

func (jq *jobsQuery) DeleteJob(id uint) error {
	var job models.ChurchJob
	err := jq.dbRepo.DB.Debug().Where("id = ?", id).Delete(&job).Error
	if err != nil {
		jq.dbRepo.App.ErrorLog.Error("error when trying to delete events post", zap.Any("error", err))
		return err
	}
	return nil
}

func (jq *jobsQuery) GetJobByEvent(eventId uint, jobId uint) (*models.ChurchJob, error) {
	var job models.ChurchJob
	err := jq.dbRepo.DB.Debug().Preload(clause.Associations).Where("id  = ?", jobId).Where("church_event_id = ?", eventId).First(&job).Error
	if err != nil || err == gorm.ErrRecordNotFound {
		jq.dbRepo.App.ErrorLog.Error("error when trying to get  job", zap.Any("error", err))
		return &job, err
	}
	return &job, nil
}

func (jq *jobsQuery) GetAllEventsJobs(id uint) ([]*models.ChurchJob, int64, error) {
	var jobs []*models.ChurchJob
	var count int64

	//val := jq.dbRepo.DB.Raw("SELECT ce.*,eo.*, cj.* FROM church_jobs AS ce, church_jobs AS cj, users AS eo\nWHERE cj.church_event_id = ce.id and eo.id = ce.organizer_id and ce.deleted_at IS NULL and cj.deleted_at IS NULL;").Preload(clause.Associations).Scan(&jobs).Error
	val := jq.dbRepo.DB.Debug().Preload(clause.Associations).Where("church_event_id = ?", id).Find(&jobs).Error
	count = int64(len(jobs))
	if val != nil {
		log.Println(val)
		return jobs, 0, val
	}

	log.Println(&jobs)
	return jobs, count, nil
}

func (jq *jobsQuery) UpdateJob(id uint, job models.ChurchJob) (*models.ChurchJob, error) {
	err := jq.dbRepo.DB.Debug().Where("id = ?", id).Updates(&job).Error
	if err != nil || err == gorm.ErrRecordNotFound {
		jq.dbRepo.App.ErrorLog.Error("error when trying to update  job", zap.Any("error", err))
		return &job, err
	}
	return &job, nil
}
func (jq *jobsQuery) GetSingleJob(id uint) (*models.ChurchJob, error) {
	var job models.ChurchJob
	err := jq.dbRepo.DB.Debug().Where("id = ?", id).First(&job).Error
	if err != nil || err == gorm.ErrRecordNotFound {
		jq.dbRepo.App.ErrorLog.Error("error when trying to get single job ", zap.Any("error", err))
		return &job, err
	}
	return &job, nil
}
