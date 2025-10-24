package repository

import (
	"github.com/jmoiron/sqlx"
	// "github.com/Mamvriyskiy/lab1-template/person/model"
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
