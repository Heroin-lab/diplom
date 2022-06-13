package repositories

import (
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user User) (int, error)
	GetUserByLogin(username string, password string) (User, error)
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

type Repositories struct {
	Authorization
	CabMansRepository
	OrdersRepository
}

func NewRepositories(db *sqlx.DB) *Repositories {
	return &Repositories{
		Authorization:     NewAuthPostgres(db),
		CabMansRepository: NewCabManPostgres(db),
		OrdersRepository:  NewOrderPostgres(db),
	}
}
