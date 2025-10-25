package repository

import (
	"github.com/jmoiron/sqlx"
	model "github.com/Mamvriyskiy/lab2-template/src/bonus/model"
)

type RepoBonus interface {
	GetInfoAboutUserPrivilege(username string) (model.PrivilegeResponse, error)
	UpdateBonus(username, ticketUid string) error
}

type Repository struct {
	RepoBonus
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		RepoBonus: NewBonusPostgres(db),
	}
}
