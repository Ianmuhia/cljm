package services

import (
	"log"
	"maranatha_web/internal/logger"
	"maranatha_web/internal/models"
	"maranatha_web/internal/repository"
)

type JobsServiceInterface interface {
	GetAllEventJobs(id uint) ([]*models.ChurchJob, int64, error)
	GetJobByEvent(eventId uint, jobId uint) (*models.ChurchJob, error)
	CreateEventJob(jobsModel models.ChurchJob) (*models.ChurchJob, error)
	DeleteJob(jobId uint) error
	UpdateJob(id uint, job models.ChurchJob) error
	GetSingleJob(id uint) (*models.ChurchJob, error)
}

type jobsService struct {
	dao repository.DAO
}

func NewJobsService(dao repository.DAO) JobsServiceInterface {
	return &jobsService{dao: dao}
}

func (js *jobsService) CreateEventJob(jobsModel models.ChurchJob) (*models.ChurchJob, error) {
	if err := js.dao.NewJobsQuery().CreateJob(&jobsModel); err != nil {
		return nil, err
	}
	return &jobsModel, nil
}

func (js *jobsService) GetAllEventJobs(id uint) ([]*models.ChurchJob, int64, error) {
	data, count, err := js.dao.NewJobsQuery().GetAllEventsJobs(id)
	if err != nil {
		return data, count, err
	}
	return data, count, nil
}

func (js *jobsService) GetJobByEvent(eventId uint, jobId uint) (*models.ChurchJob, error) {
	job, err := js.dao.NewJobsQuery().GetJobByEvent(eventId, jobId)
	if err != nil {
		log.Println(err)
		return job, err
	}
	return job, nil
}

func (js *jobsService) DeleteJob(jobId uint) error {
	err := js.dao.NewJobsQuery().DeleteJob(jobId)
	if err != nil {
		logger.GetLogger().Error("Could not delete item")
		return err
	}
	return nil
}

func (js *jobsService) UpdateJob(id uint, job models.ChurchJob) error {
	_, err := js.dao.NewJobsQuery().UpdateJob(id, job)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (js *jobsService) GetSingleJob(id uint) (*models.ChurchJob, error) {
	job, err := js.dao.NewJobsQuery().GetSingleJob(id)
	if err != nil {
		log.Println(err)
		return job, err
	}
	return job, nil
}
