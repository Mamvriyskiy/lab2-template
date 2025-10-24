package services

type BonusService struct {
	repo repository.RepoBonus
}

func NewBonusService(repo repository.RepoBonus) *BonusService {
	return &BonusService{repo: repo}
}
