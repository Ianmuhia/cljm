package services

import (
	"log"

	"maranatha_web/domain/Jobs"
	"maranatha_web/logger"
	"maranatha_web/models"
	"maranatha_web/utils/errors"
)

var (
	JobsService jobsServiceInterface = &jobsService{}
)

type jobsService struct{}

type jobsServiceInterface interface {
	GetAllEventJobs(id uint) ([]*models.ChurchJob, int64, *errors.RestErr)
	GetJobByEvent(eventId uint, jobId uint) (*models.ChurchJob, *errors.RestErr)
	CreateEventJob(jobsModel models.ChurchJob) (*models.ChurchJob, *errors.RestErr)
	DeleteJob(jobId uint) *errors.RestErr
	UpdateJob(id uint, job models.ChurchJob) *errors.RestErr
	GetSingleJob(id uint) (*models.ChurchJob, *errors.RestErr)
}

func (b *jobsService) CreateEventJob(jobsModel models.ChurchJob) (*models.ChurchJob, *errors.RestErr) {
	if err := Jobs.CreateJob(&jobsModel); err != nil {
		return nil, err
	}
	return &jobsModel, nil
}

func (b *jobsService) GetAllEventJobs(id uint) ([]*models.ChurchJob, int64, *errors.RestErr) {
	data, count, err := Jobs.GetAllEventsJobs(id)
	if err != nil {
		return data, count, errors.NewBadRequestError("Could not get post")
	}
	return data, count, nil
}

func (b *jobsService) GetJobByEvent(eventId uint, jobId uint) (*models.ChurchJob, *errors.RestErr) {
	job, err := Jobs.GetJobByEvent(eventId, jobId)
	if err != nil {
		log.Println(err)
		return job, errors.NewBadRequestError("Could not delete item")
	}
	return job, nil
}

func (b *jobsService) DeleteJob(jobId uint) *errors.RestErr {
	err := Jobs.DeleteJob(jobId)
	if err != nil {
		logger.GetLogger().Error("Could not delete item")
		return errors.NewBadRequestError("Could not delete item")
	}
	return nil
}

func (b *jobsService) UpdateJob(id uint, job models.ChurchJob) *errors.RestErr {
	_, err := Jobs.UpdateJob(id, job)
	if err != nil {
		log.Println(err)
		return errors.NewBadRequestError("Could not update job")
	}
	return nil
}

func (b *jobsService) GetSingleJob(id uint) (*models.ChurchJob, *errors.RestErr) {
	job, err := Jobs.GetSingleJob(id)
	if err != nil {
		log.Println(err)
		return job, errors.NewBadRequestError("Could not get single job")
	}
	return job, nil
}
