package services

import (
	"fmt"
	"log"

	"maranatha_web/internal/config"
	"maranatha_web/internal/models"
	"maranatha_web/internal/repository"
)

type ChurchPartnersService interface {
	CreateChurchPartner(partnersModel models.ChurchPartner) (*models.ChurchPartner, error)
	GetAllChurchPartner() ([]*models.ChurchPartner, int64, error)
	DeleteChurchPartner(id uint) error
	GetSingleChurchPartner(id uint) (*models.ChurchPartner, error)
	UpdateChurchPartner(id uint, newModel models.ChurchPartner) error
}
type churchPartnersService struct {
	dao repository.DAO
	cfg *config.AppConfig
}

func NewChurchPartnersService(dao repository.DAO, cfg *config.AppConfig) ChurchPartnersService {
	return &churchPartnersService{dao: dao, cfg: cfg}
}

func (cps *churchPartnersService) CreateChurchPartner(churchPartnersModel models.ChurchPartner) (*models.ChurchPartner, error) {

	if err := cps.dao.NewPartnersQuery().CreateChurchPartner(&churchPartnersModel); err != nil {
		return nil, err
	}
	return &churchPartnersModel, nil
}

func (cps *churchPartnersService) GetAllChurchPartner() ([]*models.ChurchPartner, int64, error) {
	data, count, err := cps.dao.NewPartnersQuery().GetAllChurchPartner()
	for _, v := range data {
		v.Image = fmt.Sprintf("%s/%s/%s", cps.cfg.StorageURL, cps.cfg.StorageBucket, v.Image)
	}
	if err != nil {
		return data, count, err
	}
	return data, count, nil
}

func (cps *churchPartnersService) DeleteChurchPartner(id uint) error {
	err := cps.dao.NewPartnersQuery().DeleteChurchPartner(id)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (cps *churchPartnersService) GetSingleChurchPartner(id uint) (*models.ChurchPartner, error) {
	data, err := cps.dao.NewPartnersQuery().GetSingleChurchPartner(id)
	storageURL := cps.cfg.StorageURL.String()
	url := fmt.Sprintf("%s/%s", storageURL, data.Image)
	data.Image = url
	if err != nil {
		log.Println(err)
		return data, err
	}
	return data, nil
}

func (cps *churchPartnersService) UpdateChurchPartner(id uint, newModel models.ChurchPartner) error {
	data, err := cps.dao.NewPartnersQuery().UpdateChurchPartner(id, newModel)
	url := fmt.Sprintf("%s/%s/%s", cps.cfg.StorageURL.String(), cps.cfg.StorageBucket, data.Image)
	data.Image = url
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
