package models

import (
	"enjoys.in/first-app/database"
	"time"
)

type Event struct {
	ID          int64
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserID      int64
}

var events = []Event{}

func (e Event) Save() error {
	stmt, err := database.DB.Prepare("INSERT INTO events (name, description, location, date_time, user_id) VALUES (?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	result, err := stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserID)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	e.ID = id

	events = append(events, e)
	return err

}
func GetALlEvents() ([]Event, error) {
	query := "SELECT * FROM events"
	rows, err := database.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var events []Event
	for rows.Next() {
		var event Event

		err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)
		if err != nil {
			return nil, err
		}

		events = append(events, event)
	}
	return events, nil
}
func GetEventById(id int64) (Event, error) {
	query := "SELECT * FROM events WHERE id = ?"
	row := database.DB.QueryRow(query, id)
	var event Event
	err := row.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)
	if err != nil {
		return event, err
	}
	return event, nil
}
