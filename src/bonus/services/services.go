package services

import (
	repository "github.com/Mamvriyskiy/lab2-template/src/bonus/repository"
	model "github.com/Mamvriyskiy/lab2-template/src/bonus/model"
)

type Bonus interface {
	GetInfoAboutUserPrivilege(username string) (model.PrivilegeResponse, error)
	UpdateBonus(username, ticketUid string) error
}

type Services struct {
	Bonus
}

func NewServices(repo *repository.Repository) *Services {
	return &Services{
		Bonus: NewBonusService(repo.RepoBonus),
	}
}
