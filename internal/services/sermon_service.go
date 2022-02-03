package services

import (
	"log"
	"maranatha_web/internal/models"
	"maranatha_web/internal/repository"
)

type SermonService interface {
	CreateSermon(partnersModel models.Sermon) (*models.Sermon, error)
	GetAllSermon() ([]models.Sermon, int64, error)
	DeleteSermon(id uint) error
	GetSingleSermon(id uint) (*models.Sermon, error)
	UpdateSermon(id uint, newModel models.Sermon) error
}

type sermonService struct {
	dao repository.DAO
}

func NewSermonService(dao repository.DAO) SermonService {
	return &sermonService{dao: dao}
}

func (ss *sermonService) CreateSermon(churchPartnersModel models.Sermon) (*models.Sermon, error) {
	if err := ss.dao.NewSermonQuery().CreateSermon(&churchPartnersModel); err != nil {

		return nil, err
	}
	return &churchPartnersModel, nil
}

func (ss *sermonService) GetAllSermon() ([]models.Sermon, int64, error) {
	data, count, err := ss.dao.NewSermonQuery().GetAllSermon()
	if err != nil {
		return data, count, err

	}

	return data, count, nil
}

func (ss *sermonService) DeleteSermon(id uint) error {
	err := ss.dao.NewSermonQuery().DeleteSermon(id)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (ss *sermonService) GetSingleSermon(id uint) (*models.Sermon, error) {
	data, err := ss.dao.NewSermonQuery().GetSingleSermon(id)
	if err != nil {
		log.Println(err)
		return data, err
	}
	return data, nil
}

func (ss *sermonService) UpdateSermon(id uint, newModel models.Sermon) error {
	_, err := ss.dao.NewSermonQuery().UpdateSermon(id, newModel)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
