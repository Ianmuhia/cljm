package services

import (
	"fmt"
	"log"

	church_patners "maranatha_web/domain/church_patner"
	"maranatha_web/models"
	"maranatha_web/utils/errors"
)

var (
	ChurchPartnersService churchPartnersServiceInterface = &churchPartnersService{}
)

type churchPartnersService struct{}

type churchPartnersServiceInterface interface {
	CreateChurchPartner(partnersModel models.ChurchPartner) (*models.ChurchPartner, *errors.RestErr)
	GetAllChurchPartner() ([]models.ChurchPartner, int64, *errors.RestErr)
	DeleteChurchPartner(id uint) *errors.RestErr
	GetSingleChurchPartner(id uint) (*models.ChurchPartner, *errors.RestErr)
	UpdateChurchPartner(id uint, newModel models.ChurchPartner) *errors.RestErr
}

func (s *churchPartnersService) CreateChurchPartner(churchPartnersModel models.ChurchPartner) (*models.ChurchPartner, *errors.RestErr) {
	if err := church_patners.CreateChurchPartner(&churchPartnersModel); err != nil {
		return nil, err
	}
	return &churchPartnersModel, nil
}

func (s *churchPartnersService) GetAllChurchPartner() ([]models.ChurchPartner, int64, *errors.RestErr) {
	data, count, err := church_patners.GetAllChurchPartner()
	if err != nil {
		return data, count, errors.NewBadRequestError("Could not get partners")

	}

	return data, count, nil
}

func (s *churchPartnersService) DeleteChurchPartner(id uint) *errors.RestErr {
	err := church_patners.DeleteChurchPartner(id)
	if err != nil {
		log.Println(err)
		return errors.NewBadRequestError("Could not delete partner")
	}
	return nil
}

func (s *churchPartnersService) GetSingleChurchPartner(id uint) (*models.ChurchPartner, *errors.RestErr) {
	data, err := church_patners.GetSingleChurchPartner(id)
	url := fmt.Sprintf("http://localhost:9000/mono/%s", data.Image)
	data.Image = url
	if err != nil {
		log.Println(err)
		return data, errors.NewBadRequestError("Could not get single partner")
	}
	return data, nil
}

func (s *churchPartnersService) UpdateChurchPartner(id uint, newModel models.ChurchPartner) *errors.RestErr {
	data, err := church_patners.UpdateChurchPartner(id, newModel)
	url := fmt.Sprintf("http://localhost:9000/mono/%s", data.Image)
	data.Image = url
	if err != nil {
		log.Println(err)
		return errors.NewBadRequestError("Could not update partner")
	}
	return nil
}
