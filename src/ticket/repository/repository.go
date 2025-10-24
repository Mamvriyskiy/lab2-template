package repository

import (
	"github.com/jmoiron/sqlx"
	//
)

type RepoTicket interface {
}

type Repository struct {
	RepoTicket
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		RepoTicket: NewTicketPostgres(db),
	}
}
