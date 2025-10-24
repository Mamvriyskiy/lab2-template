package services

import (
	"github.com/Mamvriyskiy/lab1-template/person/model"
	"github.com/Mamvriyskiy/lab1-template/person/repository"
)

type Bonus interface {
}

type Services struct {
	Bonus
}

func NewServices(repo * repository.Repository) *Services {
	return &Services{
		Bonus: NewBonusService(repo.RepoBonus),
	}
}

