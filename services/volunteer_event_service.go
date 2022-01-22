package services

import (
	"maranatha_web/domain/volunteer_events"
	"maranatha_web/models"
	"maranatha_web/utils/errors"
)

var (
	VolunteerChurchJobService volunteerChurchJobServiceInterface = &volunteerChurchJobService{}
)

type volunteerChurchJobService struct{}

type volunteerChurchJobServiceInterface interface {
	CreateSubscribeToChurchJob(churchPartnersModel models.VolunteerChurchJob) (*models.VolunteerChurchJob, *errors.RestErr)
}

func (s *volunteerChurchJobService) CreateSubscribeToChurchJob(churchPartnersModel models.VolunteerChurchJob) (*models.VolunteerChurchJob, *errors.RestErr) {
	if err := volunteer_events.CreateSubscribeToChurchJob(&churchPartnersModel); err != nil {
		return nil, err
	}
	return &churchPartnersModel, nil
}
