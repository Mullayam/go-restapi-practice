package domain

import (
	"database/sql"
	"github.com/go-sql-driver/mysql"
	"time"
)

type CustomerRepositoryDB struct {
	db *sql.DB
}

func NewCustomerRepositoryDB() *CustomerRepositoryDB {
	db, err := sql.Open("mysql", "user:password@/dbname")
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(10)
	return &CustomerRepositoryDB{
		db: db,
	}
}
func (d *CustomerRepositoryDB) FindAll() ([]Customer, error) {

	var sqlQuery = `SELECT id, name, city, zipcode, date_of_birth FROM customer`
	rows, err := d.db.Query(sqlQuery)
	if err != nil {
		if mysqlErr, ok := err.(*mysql.MySQLError); ok {
			if mysqlErr.Number == 1044 {
				return []Customer{}, nil
			}
		}
		return nil, err
	}
	defer rows.Close()
	customers := make([]Customer, 0)
	for rows.Next() {
		var c Customer
		if err := rows.Scan(&c.Id, &c.Name, &c.City, &c.Zipcode, &c.DOB); err != nil {
			return nil, err
		}
		customers = append(customers, c)
	}
	return customers, nil
}
