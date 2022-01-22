package Jobs

import (
	"log"

	"gorm.io/gorm" //nolint:goimports
	"gorm.io/gorm/clause"
	postgresql_db "maranatha_web/datasources/postgresql"
	"maranatha_web/logger"
	"maranatha_web/models"
	"maranatha_web/utils/errors" //nolint:goimports
)

func CreateJob(job *models.ChurchJob) *errors.RestErr {
	err := postgresql_db.Client.Debug().Create(&job).Error
	if err != nil {
		logger.Error("error when trying to create church event.", err)
		return errors.NewInternalServerError("database error")
	}
	return nil
}

func DeleteJob(id uint) *errors.RestErr {
	var job models.ChurchJob
	err := postgresql_db.Client.Debug().Where("id = ?", id).Delete(&job).Error
	if err != nil {
		logger.Error("error when trying to delete events post", err)
		return errors.NewInternalServerError("database error")
	}
	return nil
}

func GetJobByEvent(eventId uint, jobId uint) (*models.ChurchJob, *errors.RestErr) {
	var job models.ChurchJob
	err := postgresql_db.Client.Debug().Preload(clause.Associations).Where("id  = ?", jobId).Where("church_event_id = ?", eventId).First(&job).Error
	if err != nil || err == gorm.ErrRecordNotFound {
		logger.Error("error when trying to get  job", err)
		return &job, errors.NewInternalServerError("database error")
	}
	return &job, nil
}

func GetAllEventsJobs(id uint) ([]*models.ChurchJob, int64, error) {
	var jobs []*models.ChurchJob
	var count int64

	//val := postgresql_db.Client.Raw("SELECT ce.*,eo.*, cj.* FROM church_jobs AS ce, church_jobs AS cj, users AS eo\nWHERE cj.church_event_id = ce.id and eo.id = ce.organizer_id and ce.deleted_at IS NULL and cj.deleted_at IS NULL;").Preload(clause.Associations).Scan(&jobs).Error
	val := postgresql_db.Client.Debug().Preload(clause.Associations).Where("church_event_id = ?", id).Find(&jobs).Error
	count = int64(len(jobs))
	if val != nil {
		log.Println(val)
		return jobs, 0, val
	}

	log.Println(&jobs)
	return jobs, count, nil
}

func UpdateJob(id uint, job models.ChurchJob) (*models.ChurchJob, *errors.RestErr) {
	err := postgresql_db.Client.Debug().Where("id = ?", id).Updates(&job).Error
	if err != nil || err == gorm.ErrRecordNotFound {
		logger.Error("error when trying to update  job", err)
		return &job, errors.NewInternalServerError("database error")
	}
	return &job, nil
}
func GetSingleJob(id uint) (*models.ChurchJob, *errors.RestErr) {
	var job models.ChurchJob
	err := postgresql_db.Client.Debug().Where("id = ?", id).First(&job).Error
	if err != nil || err == gorm.ErrRecordNotFound {
		logger.Error("error when trying to get single job ", err)
		return &job, errors.NewInternalServerError("database error")
	}
	return &job, nil
}
