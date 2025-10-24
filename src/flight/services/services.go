package services

import (
	"github.com/Mamvriyskiy/lab1-template/person/model"
	"github.com/Mamvriyskiy/lab1-template/person/repository"
)

type Flight interface {
}

type Services struct {
	Flight
}

func NewServices(repo * repository.Repository) *Services {
	return &Services{
		Flight: NewFlightService(repo.RepoFlight),
	}
}

