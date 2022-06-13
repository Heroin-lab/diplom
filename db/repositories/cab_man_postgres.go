package repositories

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

type CabManPostgres struct {
	db *sqlx.DB
}

type CabMan struct {
	CabManId      int    `json:"cab_man_id" db:"cab_man_id"`
	FirstName     string `json:"first_name" db:"first_name"`
	SecondName    string `json:"second_name" db:"second_name"`
	VehicleNumber string `json:"vehicle_number" db:"vehicle_number"`
	Image         string `json:"image" db:"image"`
}

func NewCabManPostgres(db *sqlx.DB) *CabManPostgres {
	return &CabManPostgres{db: db}
}

func (r *CabManPostgres) GetAll() ([]CabMan, error) {
	var cabList []CabMan

	query := fmt.Sprintf("SELECT * FROM %s", cabMansTable)
	err := r.db.Select(&cabList, query)

	return cabList, err
}

func (r *CabManPostgres) GetOne(rowId int) ([]CabMan, error) {
	var cabMan []CabMan

	query := fmt.Sprintf("SELECT * FROM %s WHERE cab_man_id= $1", cabMansTable)
	err := r.db.Select(&cabMan, query, rowId)

	return cabMan, err
}

func (r *CabManPostgres) CreateNewCabMan(newCabMan CabMan) (int, error) {
	var id int

	query := fmt.Sprintf("INSERT INTO %s(first_name, second_name, vehicle_number, image) VALUES ($1, $2, $3, $4) RETURNING cab_man_id",
		cabMansTable)
	row := r.db.QueryRow(query, newCabMan.FirstName, newCabMan.SecondName, newCabMan.VehicleNumber, newCabMan.Image)
	err := row.Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (r *CabManPostgres) UpdateOldCabMan(rowId int, cabInfo CabMan) (int, error) {
	var updRow int

	query := fmt.Sprintf("UPDATE %s SET first_name= $1, second_name= $2, vehicle_number= $3, image= $4 WHERE cab_man_id= %d RETURNING cab_man_id",
		cabMansTable, rowId)
	row := r.db.QueryRow(query, cabInfo.FirstName, cabInfo.SecondName, cabInfo.VehicleNumber, cabInfo.Image)
	err := row.Scan(&updRow)

	return updRow, err
}

func (r *CabManPostgres) DeleteOldCabMan(rowId int) (int, error) {
	var delRowId int

	query := fmt.Sprintf("DELETE FROM %s WHERE cab_man_id= %d RETURNING cab_man_id",
		cabMansTable, rowId)
	row := r.db.QueryRow(query)
	err := row.Scan(&delRowId)
	if err != nil {
		return 0, err
	}

	return delRowId, nil
}
