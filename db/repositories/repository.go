package repositories

import (
	"github.com/Heroin-lab/taxi_service.git/server/request"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user User) (int, error)
	GetUserByPhoneNumber(phoneNum string, password string) (request.UserSignIn, error)
}

type CabMansRepository interface {
	GetAll() ([]CabMan, error)
	GetOne(rowId int) ([]CabMan, error)
	CreateNewCabMan(newCabMan CabMan) (int, error)
	UpdateOldCabMan(rowId int, cabInfo CabMan) (int, error)
	DeleteOldCabMan(rowId int) (int, error)
}

type OrdersRepository interface {
	GetDriversOrders(status int, driverId int, offsetNum int) ([]CurrentDriverOrder, error)
	CreateNewLocationRow(lat string, lon string, orderId int) error
	CreateNewOrders(startLoc []string, endLoc []string) error
}

type OffenseRepository interface {
	GetAllUserOffenses(uId int) ([]Offense, error)
	CreateUserOffense(offense Offense) (int, error)
	UpdateUserOffense(offense Offense) (int, error)
	DeleteUserOffense(offenseId int) (int, error)
}

type Repositories struct {
	Authorization
	CabMansRepository
	OrdersRepository
	OffenseRepository
}

func NewRepositories(db *sqlx.DB) *Repositories {
	return &Repositories{
		Authorization:     NewAuthPostgres(db),
		CabMansRepository: NewCabManPostgres(db),
		OrdersRepository:  NewOrderPostgres(db),
		OffenseRepository: NewOffensePostgres(db),
	}
}
