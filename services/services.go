package services

import (
	"database/sql"
	"errors"
	"github.com/Heroin-lab/taxi_service.git/db/repositories"
	"github.com/Heroin-lab/taxi_service.git/server/request"
	"github.com/sirupsen/logrus"
	"strings"
)

type Authorization interface {
	CreateUser(user request.User) (int, error)
	GenerateToken(login string, password string, duration int) (string, error)
	ParseToken(token string) (int, error)
}

type CabMansServices interface {
	GetAllCabMans() ([]repositories.CabMan, error)
	GetOneCabMan(rowId string) ([]repositories.CabMan, error)
	CreateNewCabMan(cm request.CabMan) (int, error)
	UpdateCabMan(rowId string, cabInfo request.CabMan) (int, error)
	DeleteCabMan(rowId string) (int, error)
}

type OrdersServices interface {
	GetOrdersByStatus(status string, offsetNum string) ([]repositories.CurrentDriverOrder, error)
	GetOrdersByDriverId(driverId string, offsetNum string) ([]repositories.CurrentDriverOrder, error)
	CreateOrders(startLoc []string, endLoc []string) error
	CreateNewLocationRow(lat string, lon string, orderId int) error
}

type Services struct {
	Authorization
	CabMansServices
	OrdersServices
}

func NewServices(repos *repositories.Repositories) *Services {
	return &Services{
		Authorization:   NewAuthService(repos.Authorization),
		CabMansServices: NewCabManService(repos.CabMansRepository),
		OrdersServices:  NewOrdersService(repos.OrdersRepository),
	}
}

func CheckError(err error, mode string) error {
	logrus.Error(err)
	if errors.Is(err, sql.ErrNoRows) {
		logrus.Errorf("No rows in result set when user try to %s some row!", mode)
		return errors.New("Cab man not found!")
	}

	if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
		return errors.New("Cab man with this vehicle number already exist!")
	}

	return errors.New("Something went wrong!")
}
