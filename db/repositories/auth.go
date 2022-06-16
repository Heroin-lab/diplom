package repositories

import (
	"fmt"
	"github.com/Heroin-lab/taxi_service.git/server/request"
	"github.com/jmoiron/sqlx"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user User) (int, error) {
	var id int

	query := fmt.Sprint("INSERT INTO users(first_name, second_name, patronymic, phone_number, password_hash) VALUES ($1, $2, $3, $4, $5) RETURNING id")
	row := r.db.QueryRow(query, user.FirstName, user.SecondName, user.Patronymic, user.PhoneNumber, user.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *AuthPostgres) GetUserByPhoneNumber(phoneNum string, password string) (request.UserSignIn, error) {
	var user request.UserSignIn

	query := fmt.Sprintf("SELECT id FROM %s WHERE phone_number= $1 AND password_hash= $2", usersTable)
	err := r.db.Get(&user, query, phoneNum, password)

	return user, err
}
