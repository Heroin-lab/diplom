package services

import (
	"errors"
	"github.com/Heroin-lab/taxi_service.git/db/repositories"
	"github.com/Heroin-lab/taxi_service.git/server/request"
	"github.com/sirupsen/logrus"
	"strconv"
)

type CabManService struct {
	repo repositories.CabMansRepository
}

func NewCabManService(repo repositories.CabMansRepository) *CabManService {
	return &CabManService{repo: repo}
}

func (s *CabManService) GetAllCabMans() ([]repositories.CabMan, error) {
	cabMansList, err := s.repo.GetAll()
	if err != nil {
		return []repositories.CabMan{}, CheckError(err, "create")
	}

	return cabMansList, err
}

func (s *CabManService) GetOneCabMan(rowId string) ([]repositories.CabMan, error) {
	intRowId, err := strconv.Atoi(rowId)
	if err != nil {
		logrus.Error(err)
		return []repositories.CabMan{}, errors.New("Unprocessable id entity!")
	}

	return s.repo.GetOne(intRowId)
}

func (s *CabManService) CreateNewCabMan(newCab request.CabMan) (int, error) {
	cabMan := repositories.CabMan{
		FirstName:     newCab.FirstName,
		SecondName:    newCab.SecondName,
		VehicleNumber: newCab.VehicleNumber,
		Image:         newCab.Image,
	}

	newCabMan, err := s.repo.CreateNewCabMan(cabMan)
	if err != nil {
		return 0, CheckError(err, "create")
	}

	return newCabMan, err
}

func (s *CabManService) UpdateCabMan(rowId string, cabInfo request.CabMan) (int, error) {
	intRowId, err := strconv.Atoi(rowId)
	if err != nil {
		logrus.Error(err)
		return 0, errors.New("Unprocessable id entity!")
	}

	cabManInfo := repositories.CabMan{
		CabManId:      cabInfo.CabManId,
		FirstName:     cabInfo.FirstName,
		SecondName:    cabInfo.SecondName,
		VehicleNumber: cabInfo.VehicleNumber,
		Image:         cabInfo.Image,
	}

	updatedRowId, err := s.repo.UpdateOldCabMan(intRowId, cabManInfo)
	if err != nil {
		return 0, CheckError(err, "update")
	}

	return updatedRowId, err
}

func (s *CabManService) DeleteCabMan(rowId string) (int, error) {
	intRowId, err := strconv.Atoi(rowId)
	if err != nil {
		logrus.Error(err)
		return 0, errors.New("Unprocessable id entity!")
	}

	deletedRowId, err := s.repo.DeleteOldCabMan(intRowId)
	if err != nil {
		return 0, CheckError(err, "delete")
	}

	return deletedRowId, err
}
