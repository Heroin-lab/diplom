package services

import (
	"github.com/Heroin-lab/taxi_service.git/db/repositories"
	"github.com/sirupsen/logrus"
	"strconv"
)

type OffensesService struct {
	repo repositories.OffenseRepository
}

func (s *OffensesService) CreateUserOffense(offense repositories.Offense) (int, error) {
	return s.repo.CreateUserOffense(offense)
}

func (s *OffensesService) UpdateUserOffense(offense repositories.Offense) (int, error) {
	return s.repo.UpdateUserOffense(offense)
}

func (s *OffensesService) DeleteUserOffense(offenseId string) (int, error) {
	offenseIdInt, err := strconv.Atoi(offenseId)
	if err != nil {
		logrus.Info("USUCk Service")
		return 0, err
	}

	return s.repo.DeleteUserOffense(offenseIdInt)
}

func NewOffensesService(repo repositories.OffenseRepository) *OffensesService {
	return &OffensesService{repo: repo}
}

func (s *OffensesService) GetAllUserOffenses(uId string) ([]repositories.Offense, error) {
	uIdInt, err := strconv.Atoi(uId)
	if err != nil {
		logrus.Error("USuck")
		return []repositories.Offense{}, err
	}

	return s.repo.GetAllUserOffenses(uIdInt)
}
