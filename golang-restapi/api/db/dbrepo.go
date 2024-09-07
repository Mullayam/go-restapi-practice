package dbrepo

import (
	"context"
	"database/sql"
	"log"
	"time"
)

type PostgresDBRepo struct {
	DB *sql.DB
}

const dbTimeout = time.Second * 3

// GetAll implements repository.Repository.
func (m *PostgresDBRepo) GetAll() ([]string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()
	db := m.DB
	var p = []string{"1", "2", "3"}
	err := db.PingContext(ctx)
	// rows,err:= db.QueryContext(ctx,"SELECT * FROM posts")
	// for  rows.Next() {
	// 	var p Post
	// 	err = rows.Scan(&p.ID,&p.Title,&p.Content,&p.Created,&p.Updated,&p.Expires)
	//  p.append(&p)
	// }

	// rows.Close()
	if err != nil {
		log.Println(err)
	}
	return p, nil
}

// Connection implements repository.DatabaseRepository.
func (m *PostgresDBRepo) Connection() *sql.DB {
	return m.DB
}
