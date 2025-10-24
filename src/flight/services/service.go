package services

import "github.com/!mamvriyskiy/lab2-template/src/flight/repository"

type FlightService struct {
	repo repository.RepoFlight
}

func NewFlightService(repo repository.RepoFlight) *FlightService {
	return &FlightService{repo: repo}
}
