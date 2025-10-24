package repository

import (
	"github.com/jmoiron/sqlx"
	// "github.com/Mamvriyskiy/lab1-template/person/model"
)

type TicketPostgres struct {
	db *sqlx.DB
}

func NewBonusPostgres(db *sqlx.DB) *TicketPostgres {
	return &TicketPostgres{db: db}
}