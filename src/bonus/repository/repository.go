package repository

import (
	"github.com/jmoiron/sqlx"
	model "github.com/Mamvriyskiy/lab2-template/src/bonus/model"
)

type RepoBonus interface {
	GetInfoAboutUserPrivilege(username string) (model.PrivilegeResponse, error)
	UpdateBonus(username, ticketUID string, price int) (model.PrivilegeInfo, error)
	UpdateBonusBonus(username string, price int) error
}

type Repository struct {
	RepoBonus
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		RepoBonus: NewBonusPostgres(db),
	}
}
