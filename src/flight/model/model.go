package model

import "time"

type FlightItem struct {
    FlightNumber string `json:"flightNumber"`
    FromAirport  string `json:"fromAirport"`
    ToAirport    string `json:"toAirport"`
    Date         string `json:"date"`
    Price        int    `json:"price"`
}

type FlightResponse struct {
    Page          int          `json:"page"`
    PageSize      int          `json:"pageSize"`
    TotalElements int          `json:"totalElements"`
    Items         []FlightItem `json:"items"`
}

type Flight struct {
	ID            int       `json:"id"`
	FlightNumber  string    `json:"flightNumber"`
	Datetime      time.Time `json:"datetime"`
	FromAirportID int       `json:"fromAirportId"`
	ToAirportID   int       `json:"toAirportId"`
	Price         int       `json:"price"`
}

