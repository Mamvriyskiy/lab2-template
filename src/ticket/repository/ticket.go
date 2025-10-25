package repository

import (
	"github.com/jmoiron/sqlx"
	model "github.com/Mamvriyskiy/lab2-template/src/ticket/model"
)

type TicketPostgres struct {
	db *sqlx.DB
}

func NewTicketPostgres(db *sqlx.DB) *TicketPostgres {
	return &TicketPostgres{db: db}
}

func (r *TicketPostgres) GetInfoAboutTiket(ticketUID string) (*model.Ticket, error) {
    query := `
        SELECT ticket_uid, username, flight_number, price, status
        FROM ticket
        WHERE ticket_uid = $1;
    `
    var ticket model.Ticket
    err := r.db.QueryRow(query, ticketUID).Scan(
        &ticket.TicketUID,
        &ticket.Username,
        &ticket.FlightNumber,
        &ticket.Price,
        &ticket.Status,
    )

    if err != nil {
        return nil, err
    }

    return &ticket, nil
}

func (r *TicketPostgres) GetInfoAboutTikets(username string) ([]model.Ticket, error) {
    query := `
        SELECT ticket_uid, username, flight_number, price, status
        FROM ticket
        WHERE username = $1;
    `
    rows, err := r.db.Query(query, username)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var tickets []model.Ticket
    for rows.Next() {
        var t model.Ticket
        if err := rows.Scan(&t.TicketUID, &t.Username, &t.FlightNumber, &t.Price, &t.Status); err != nil {
            return nil, err
        }
        tickets = append(tickets, t)
    }

    return tickets, nil
}
