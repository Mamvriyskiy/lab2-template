package services

import (
	"github.com/Mamvriyskiy/lab1-template/person/model"
	"github.com/Mamvriyskiy/lab1-template/person/repository"
)

type BonusService struct {
	repo repository.RepoBonus
}

func NewBonusService(repo repository.RepoBonus) *BonusService {
	return &BonusService{repo: repo}
}

