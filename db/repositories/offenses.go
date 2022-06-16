package repositories

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

type OffensePostgres struct {
	db *sqlx.DB
}

func NewOffensePostgres(db *sqlx.DB) *OffensePostgres {
	return &OffensePostgres{db: db}
}

type Offense struct {
	Id             int    `json:"id" db:"id"`
	Latitude       string `json:"latitude" db:"latitude"`
	Longitude      string `json:"longitude" db:"longitude"`
	CrimeCodeId    int    `json:"crime_code_id" db:"crime_code_id"`
	TitleCriminal  string `json:"title_criminal" db:"title_criminal"`
	NumberCriminal int    `json:"number_criminal" db:"number_criminal"`
	UserId         int    `json:"user_id" db:"user_id"`
	FirstName      string `json:"first_name" db:"first_name"`
	SecondName     string `json:"second_name" db:"second_name"`
	Patronymic     string `json:"patronymic" db:"patronymic"`
	PhoneNumber    string `json:"phone_number" db:"phone_number"`
	Description    string `json:"description" db:"description"`
	Time           string `json:"time" db:"time"`
}

func (r *OffensePostgres) GetAllUserOffenses(uId int) ([]Offense, error) {
	var allOffenses []Offense

	query := fmt.Sprintf("SELECT o.id, longitude, latitude, o.time, description, first_name, second_name, patronymic, phone_number, title_criminal, number_criminal, user_id " +
		"FROM offenses o " +
		"INNER JOIN criminals c on c.id = o.crime_code_id " +
		"INNER JOIN users u on u.id = o.user_id " +
		"WHERE user_id= $1")
	err := r.db.Select(&allOffenses, query, uId)

	return allOffenses, err
}

func (r *OffensePostgres) CreateUserOffense(offense Offense) (int, error) {
	var id int

	query := fmt.Sprint("INSERT INTO offenses(user_id, crime_code_id, longitude, latitude, description) VALUES ($1, $2, $3, $4, $5) RETURNING id")
	row := r.db.QueryRow(query, offense.UserId, offense.CrimeCodeId, offense.Longitude, offense.Latitude, offense.Description)
	err := row.Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (r *OffensePostgres) UpdateUserOffense(offense Offense) (int, error) {
	var updRow int

	query := fmt.Sprintf("UPDATE offenses SET user_id= $1, crime_code_id= $2, longitude= $3, latitude= $4, description= $5 WHERE id= $6 RETURNING id")
	row := r.db.QueryRow(query, offense.UserId, offense.CrimeCodeId, offense.Longitude, offense.Latitude, offense.Description, offense.Id)
	err := row.Scan(&updRow)

	return updRow, err
}

func (r *OffensePostgres) DeleteUserOffense(offenseId int) (int, error) {
	var delId int

	query := fmt.Sprint("DELETE FROM offenses WHERE id= $1 RETURNING id")
	row := r.db.QueryRow(query, offenseId)
	err := row.Scan(&delId)
	if err != nil {
		return 0, err
	}

	return delId, nil
}
