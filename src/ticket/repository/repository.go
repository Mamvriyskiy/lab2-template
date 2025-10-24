package repository

import (
	"github.com/jmoiron/sqlx"
	// "github.com/Mamvriyskiy/lab1-template/person/model"
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
