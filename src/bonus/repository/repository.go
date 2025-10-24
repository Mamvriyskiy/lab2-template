package repository

import (
	"github.com/jmoiron/sqlx"
	//
)

type RepoBonus interface {
}

type Repository struct {
	RepoBonus
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		RepoBonus: NewBonusPostgres(db),
	}
}
