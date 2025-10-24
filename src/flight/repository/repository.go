package repository

import (
	"github.com/jmoiron/sqlx"
	// "github.com/Mamvriyskiy/lab1-template/person/model"
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
