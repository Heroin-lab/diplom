package services

import (
	"errors"
	"github.com/Heroin-lab/taxi_service.git/db/repositories"
	"github.com/sirupsen/logrus"
	"strconv"
)

type OrdersService struct {
	repo repositories.OrdersRepository
}

func NewOrdersService(repo repositories.OrdersRepository) *OrdersService {
	return &OrdersService{repo: repo}
}

func (s *OrdersService) GetOrdersByStatus(status string, offsetNum string) ([]repositories.CurrentDriverOrder, error) {
	statusInt, err := strconv.Atoi(status)
	if err != nil {
		logrus.Error(err)
		return []repositories.CurrentDriverOrder{}, errors.New("Field has wrong type!")
	}

	offsetInt, err := strconv.Atoi(offsetNum)
	if err != nil {
		logrus.Error(err)
		return []repositories.CurrentDriverOrder{}, errors.New("Field has wrong type!")
	}

	ordersList, err := s.repo.GetDriversOrders(statusInt, 0, offsetInt)
	if err != nil {
		return []repositories.CurrentDriverOrder{}, CheckError(err, "create")
	}

	return ordersList, err
}

func (s *OrdersService) GetOrdersByDriverId(driverId string, offsetNum string) ([]repositories.CurrentDriverOrder, error) {
	driverIdInt, err := strconv.Atoi(driverId)
	if err != nil {
		logrus.Error(err)
		return []repositories.CurrentDriverOrder{}, errors.New("Unprocessable id entity!")
	}

	offsetInt, err := strconv.Atoi(offsetNum)
	if err != nil {
		logrus.Error(err)
		return []repositories.CurrentDriverOrder{}, errors.New("Field has wrong type!")
	}

	orderRow, err := s.repo.GetDriversOrders(0, driverIdInt, offsetInt)
	if err != nil {
		return []repositories.CurrentDriverOrder{}, CheckError(err, "create")
	}

	return orderRow, err
}

func (s *OrdersService) CreateOrders(startLoc []string, endLoc []string) error {
	availableOrders, err := s.repo.GetDriversOrders(1, 0, 0)
	if err != nil {
		logrus.Errorf("Error when trying to create new orders: %s", err)
	}

	if len(availableOrders) < 5 {
		err := s.repo.CreateNewOrders(startLoc, endLoc)
		return err
	}

	return err
}

func (s *OrdersService) CreateNewLocationRow(lat string, lon string, orderId int) error {
	err := s.repo.CreateNewLocationRow(lat, lon, orderId)
	if err != nil {
		logrus.Errorf("Insert location row error: %s", err)
		return errors.New("Problem with insert location!")
	}

	return err
}
