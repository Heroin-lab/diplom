package repositories

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

type OrderPostgres struct {
	db *sqlx.DB
}

type CurrentDriverOrder struct {
	Id            int    `json:"id" db:"id"`
	FirstName     string `json:"first_name" db:"first_name"`
	SecondName    string `json:"second_name" db:"second_name"`
	VehicleNumber string `json:"vehicle_number" db:"vehicle_number"`
	Image         string `json:"image" db:"image"`
	Status        string `json:"status" db:"status"`
	StartLocation string `json:"start_location" db:"start_location"`
	EndLocation   string `json:"end_location" db:"end_location"`
}

type Order struct {
	Id            int
	CabManId      int
	StatusId      int
	StartLocation string
	EndLocation   string
}

func NewOrderPostgres(db *sqlx.DB) *OrderPostgres {
	return &OrderPostgres{db: db}
}

func (r *OrderPostgres) GetDriversOrders(status int, driverId int, offsetNum int) ([]CurrentDriverOrder, error) {
	var ordersList []CurrentDriverOrder
	var resultErr error
	const limit = "20"
	whereStatement := ""
	searchArgument := 2

	if status != 0 {
		searchArgument = status
		whereStatement = "WHERE o.id_status= $1"
	}

	// Use rebind and sqlx.In

	if driverId != 0 {
		searchArgument = driverId
		whereStatement = "WHERE o.id_cab_man= $1"
	}

	query := fmt.Sprint("SELECT o.id, cm.first_name, cm.second_name, cm.vehicle_number, cm.image, s.name as status, o.start_location, o.end_location " +
		"FROM orders o" +
		" INNER JOIN cab_mans cm on cm.cab_man_id = o.id_cab_man " +
		"INNER JOIN statuses s on o.id_status = s.status_id " + whereStatement +
		" ORDER BY id DESC LIMIT $2 OFFSET $3")

	a := []interface{}{1, 2}

	if status == 0 && driverId == 0 {
		resultErr = r.db.Select(&ordersList, query, a...)
	} else {
		resultErr = r.db.Select(&ordersList, query, searchArgument, limit, offsetNum)
	}

	return ordersList, resultErr
}

func (r *OrderPostgres) CreateNewOrders(startLoc []string, endLoc []string) error {

	query := fmt.Sprint("INSERT INTO orders(id_cab_man, id_status, start_location, end_location, created_at) VALUES" +
		"(null, 1, $1, $4, current_timestamp)," +
		"(null, 1, $2, $5, current_timestamp)," +
		"(null, 1, $3, $6, current_timestamp);")

	_, err := r.db.Exec(query, startLoc[0], startLoc[1], startLoc[2], endLoc[0], endLoc[1], endLoc[2])

	return err
}

func (r *OrderPostgres) CreateNewLocationRow(lat string, lon string, orderId int) error {
	query := fmt.Sprintf("INSERT INTO %s(id_order, latitude, longitude) VALUES ($1, $2, $3)",
		ordersHistoryTable)
	_, err := r.db.Exec(query, orderId, lat, lon)

	return err
}
