package repository

import (
	"time"

	model "github.com/Mamvriyskiy/lab2-template/src/flight/model"
	"github.com/jmoiron/sqlx"
)

type FlightPostgres struct {
	db *sqlx.DB
}

func NewFlightPostgres(db *sqlx.DB) *FlightPostgres {
	return &FlightPostgres{db: db}
}

func (r *FlightPostgres) GetInfoAboutFlightByFlightNumber(flightNumber string) (model.Flight, error) {
	var flight model.Flight

	query := `
		SELECT id, flight_number, datetime, from_airport_id, to_airport_id, price
		FROM flight
		WHERE flight_number = $1
	`

	err := r.db.QueryRow(query, flightNumber).Scan(
		&flight.ID,
		&flight.FlightNumber,
		&flight.Datetime,
		&flight.FromAirportID,
		&flight.ToAirportID,
		&flight.Price,
	)
	if err != nil {
		return model.Flight{}, err
	}

	return flight, nil
}


func (r *FlightPostgres) GetFlights(page, size int) (model.FlightResponse, error) {
	offset := (page - 1) * size

	rows, err := r.db.Query(`
        SELECT
            f.flight_number,
            a_from.name AS from_airport,
            a_to.name AS to_airport,
            f.datetime,
            f.price
        FROM flight f
        JOIN airport a_from ON f.from_airport_id = a_from.id
        JOIN airport a_to ON f.to_airport_id = a_to.id
        ORDER BY f.id
        LIMIT $1 OFFSET $2`, size, offset)
	if err != nil {
		return model.FlightResponse{}, err
	}
	defer rows.Close()

	var items []model.FlightItem
	for rows.Next() {
		var item model.FlightItem
		var dt time.Time
		if err := rows.Scan(&item.FlightNumber, &item.FromAirport, &item.ToAirport, &dt, &item.Price); err != nil {
			return model.FlightResponse{}, err
		}
		item.Date = dt.Format("2006-01-02 15:04")
		items = append(items, item)
	}

	var total int
	err = r.db.QueryRow("SELECT COUNT(*) FROM flight").Scan(&total)
	if err != nil {
		return model.FlightResponse{}, err
	}

	return model.FlightResponse{
		Page:          page,
		PageSize:      size,
		TotalElements: total,
		Items:         items,
	}, nil
}
