package FS_DA

import (
	"fmt"
	"log"
	"time"
	"os"

	// FS_structs "github.com/lapayka/rsoi-2/flight_service/Structs"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DB struct {
	db *gorm.DB
}

func GetConnectionString() (string) {
	// Берём обязательные части из переменных окружения
	host, ok := os.LookupEnv("DB_HOST")
	if !ok {
		return ""
	}

	port, ok := os.LookupEnv("DB_PORT")
	if !ok {
		return ""
	}

	user, ok := os.LookupEnv("DB_USER")
	if !ok {
		return ""
	}

	password, ok := os.LookupEnv("DB_PASSWORD")
	if !ok {
		return ""
	}

	dbname, ok := os.LookupEnv("DB_NAME")
	if !ok {
		return ""
	}

	// Формируем строку подключения
	connStr := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s",
		user,
		password,
		host,
		port,
		dbname,
	)

	return connStr
}

func New(host, user, db_name, password string) (*DB, error) {
	dsn := GetConnectionString()
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("unable to connect database", err)
	}

	return &DB{db: db}, nil
}

type joinRes struct {
	ID           int64
	FlightNumber string
	Date         time.Time

	FromAirportID      int64
	FromAirportName    string
	FromAirportCity    string
	FromAirportCountry string

	ToAirportID      int64
	ToAirportName    string
	ToAirportCity    string
	ToAirportCountry string
}

type FlightShort struct {
	FlightNumber string `json:"flightNumber"`
	FromAirport  string `json:"fromAirport"`
	ToAirport    string `json:"toAirport"`
	Date         string `json:"date"`
	Price        int    `json:"price"`
}

func (d *DB) GetFlights() ([]FlightShort, error) {
	results := []joinRes{}
	if err := d.db.Table("flight").
		Select("flight.id, flight.flight_number, flight.datetime as date, fa.name as from_airport_name, fa.city as from_airport_city, ta.name as to_airport_name, ta.city as to_airport_city").
		Joins("JOIN airport fa on flight.from_airport_id = fa.id").
		Joins("JOIN airport ta on flight.to_airport_id = ta.id").
		Scan(&results).Error; err != nil {
		return nil, err
	}

	loc, _ := time.LoadLocation("Europe/Moscow")

	flights := make([]FlightShort, 0, len(results))
	for _, r := range results {
		flights = append(flights, FlightShort{
			FlightNumber: r.FlightNumber,
			FromAirport:  fmt.Sprintf("%s %s", r.FromAirportCity, r.FromAirportName),
			ToAirport:    fmt.Sprintf("%s %s", r.ToAirportCity, r.ToAirportName),
			Date:         r.Date.In(loc).Format("2006-01-02 15:04"),
			Price:        1500,
		})
	}

	return flights, nil
}

