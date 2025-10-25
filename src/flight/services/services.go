package services

import (
	"github.com/Mamvriyskiy/lab2-template/src/flight/model"
	"github.com/Mamvriyskiy/lab2-template/src/flight/repository"
)

type Flight interface {
	GetInfoAboutFlight(page, size int) (model.FlightResponse, error)
}

type Services struct {
	Flight
}

func NewServices(repo *repository.Repository) *Services {
	return &Services{
		Flight: NewFlightService(repo.RepoFlight),
	}
}
