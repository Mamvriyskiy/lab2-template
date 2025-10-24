package services

import "github.com/!mamvriyskiy/database_course/main/pkg/repository"

type Bonus interface {
}

type Services struct {
	Bonus
}

func NewServices(repo *repository.Repository) *Services {
	return &Services{
		Bonus: NewBonusService(repo.RepoBonus),
	}
}
