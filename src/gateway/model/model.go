package model

type TicketInfo struct {
    TicketUID    string `json:"ticketUid"`
    FlightNumber string `json:"flightNumber"`
    FromAirport  string `json:"fromAirport"`
    ToAirport    string `json:"toAirport"`
    Date         string `json:"date"`
    Price        int    `json:"price"`
    Status       string `json:"status"`
}
