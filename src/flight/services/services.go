package services

import "github.com/!mamvriyskiy/lab2-template/src/flight/repository"

type Flight interface {
}

type Services struct {
	Flight
}

func NewServices(repo *repository.Repository) *Services {
	return &Services{
		Flight: NewFlightService(repo.RepoFlight),
	}
}
