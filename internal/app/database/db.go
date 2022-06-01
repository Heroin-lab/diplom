package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type Database struct {
	config *Config
	db     *sql.DB
}

func New(config *Config) *Database {
	return &Database{
		config: config,
	}
}

func (d *Database) Open() error {
	db, err := sql.Open("mysql", "kirill:Ya8eb2a3l@/diplom_db")
	if err != nil {
		return err
	}

	if err := db.Ping(); err != nil {
		return err
	}

	d.db = db

	return nil
}

func (d *Database) Close() {
	d.db.Close()
}
