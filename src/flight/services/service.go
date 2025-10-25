package services

import (
	model "github.com/Mamvriyskiy/lab2-template/src/flight/model"
	"github.com/Mamvriyskiy/lab2-template/src/flight/repository"
)

type FlightService struct {
	repo repository.RepoFlight
}

func NewFlightService(repo repository.RepoFlight) *FlightService {
	return &FlightService{repo: repo}
}

func (s *FlightService) GetInfoAboutFlight(page, size int) (model.FlightResponse, error) {
	return s.repo.GetFlights(page, size)
}
