package services

import (
	"maranatha_web/internal/models"
	"maranatha_web/internal/repository"
)

type volunteerChurchJobService struct {
	dao repository.DAO
}

type VolunteerChurchJobService interface {
	CreateSubscribeToChurchJob(churchPartnersModel models.VolunteerChurchJob) (*models.VolunteerChurchJob, error)
}

func NewVolunteerChurchJobService(dao repository.DAO) VolunteerChurchJobService {
	return &volunteerChurchJobService{dao: dao}
}

func (vs *volunteerChurchJobService) CreateSubscribeToChurchJob(churchPartnersModel models.VolunteerChurchJob) (*models.VolunteerChurchJob, error) {
	if err := vs.dao.NewVolunteerQuery().CreateSubscribeToChurchJob(&churchPartnersModel); err != nil {
		return nil, err
	}
	return &churchPartnersModel, nil
}
