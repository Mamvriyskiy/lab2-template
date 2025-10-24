package services

import (
	"github.com/Mamvriyskiy/lab1-template/person/model"
	"github.com/Mamvriyskiy/lab1-template/person/repository"
)

type FlightService struct {
	repo repository.RepoFlight
}

func NewFlightService(repo repository.RepoFlight) *FlightService {
	return &FlightService{repo: repo}
}

