package services

import (
	repository "github.com/Mamvriyskiy/lab2-template/src/bonus/repository"
	model "github.com/Mamvriyskiy/lab2-template/src/bonus/model"
)

type BonusService struct {
	repo repository.RepoBonus
}

func NewBonusService(repo repository.RepoBonus) *BonusService {
	return &BonusService{repo: repo}
}

func (s *BonusService) GetInfoAboutUserPrivilege(username string) (model.PrivilegeResponse, error) {
	return s.repo.GetInfoAboutUserPrivilege(username)
}

func (s *BonusService) UpdateBonus(username, ticketUid string) error {
	return s.repo.UpdateBonus(username, ticketUid)
}
