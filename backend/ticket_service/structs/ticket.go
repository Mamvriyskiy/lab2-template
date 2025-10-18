package TS_structs

type Ticket struct {
	ID           int64  `json:"id"`
	TicketUid    string `json:"ticketUid"`
	Username     string `json:"username"`
	FlightNumber string `json:"flight_number"`
	Price        int64  `json:"price"`
	Status       string `json:"status"`
}

func (Ticket) TableName() string {
	return "ticket"
}

type Tickets []Ticket

func (Tickets) TableName() string {
	return "ticket"
}
