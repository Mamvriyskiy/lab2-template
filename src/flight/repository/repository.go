package repository

import (
	"github.com/jmoiron/sqlx"
	//
)

type RepoFlight interface {
}

type Repository struct {
	RepoFlight
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		RepoFlight: NewFlightPostgres(db),
	}
}
