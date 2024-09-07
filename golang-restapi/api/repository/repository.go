package repository

import "database/sql"

type DatabaseRepository interface {
	Connection() *sql.DB
	GetAll() ([]string, error)
}
