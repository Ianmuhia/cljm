package services

// import (
// 	"log"
//

// 	"maranatha_web/domain/sermons"
// 	"maranatha_web/models"
// 	"maranatha_web/utils/errors"
// )

// var (
// 	SermonService sermonServiceInterface = &sermonService{}
// )

// type sermonService struct{}

// type sermonServiceInterface interface {
// 	CreateSermon(partnersModel models.Sermon) (*models.Sermon, *errors.RestErr)
// 	GetAllSermon() ([]models.Sermon, int64, *errors.RestErr)
// 	DeleteSermon(id uint) *errors.RestErr
// 	GetSingleSermon(id uint) (*models.Sermon, *errors.RestErr)
// 	UpdateSermon(id uint, newModel models.Sermon) *errors.RestErr
// }

// func (s *sermonService) CreateSermon(churchPartnersModel models.Sermon) (*models.Sermon, *errors.RestErr) {
// 	if err := sermons.CreateSermon(&churchPartnersModel); err != nil {
// 		return nil, err
// 	}
// 	return &churchPartnersModel, nil
// }

// func (s *sermonService) GetAllSermon() ([]models.Sermon, int64, *errors.RestErr) {
// 	data, count, err := sermons.GetAllSermon()
// 	if err != nil {
// 		return data, count, errors.NewBadRequestError("Could not get sermon")

// 	}

// 	return data, count, nil
// }

// func (s *sermonService) DeleteSermon(id uint) *errors.RestErr {
// 	err := sermons.DeleteSermon(id)
// 	if err != nil {
// 		log.Println(err)
// 		return errors.NewBadRequestError("Could not delete sermon")
// 	}
// 	return nil
// }

// func (s *sermonService) GetSingleSermon(id uint) (*models.Sermon, *errors.RestErr) {
// 	data, err := sermons.GetSingleSermon(id)
// 	if err != nil {
// 		log.Println(err)
// 		return data, errors.NewBadRequestError("Could not get single sermon")
// 	}
// 	return data, nil
// }

// func (s *sermonService) UpdateSermon(id uint, newModel models.Sermon) *errors.RestErr {
// 	_, err := sermons.UpdateSermon(id, newModel)
// 	if err != nil {
// 		log.Println(err)
// 		return errors.NewBadRequestError("Could not update sermon")
// 	}
// 	return nil
// }
