package repository

import (
	"github.com/jmoiron/sqlx"
	model "github.com/Mamvriyskiy/lab2-template/src/ticket/model"
)

type RepoTicket interface {
	GetInfoAboutTiket(ticketUID string) (model.Ticket, error)
	GetInfoAboutTikets(username string) ([]model.Ticket, error)
}

type Repository struct {
	RepoTicket
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		RepoTicket: NewTicketPostgres(db),
	}
}
