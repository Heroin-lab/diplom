package repositories

import (
	"fmt"
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

	query := fmt.Sprintf("INSERT INTO %s (login, password_hash) VALUES ($1, $2) RETURNING user_id", usersTable)
	row := r.db.QueryRow(query, user.Login, user.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *AuthPostgres) GetUserByLogin(username string, password string) (User, error) {
	var user User

	query := fmt.Sprintf("SELECT user_id FROM %s WHERE login= $1 AND password_hash= $2", usersTable)
	err := r.db.Get(&user, query, username, password)

	return user, err
}
